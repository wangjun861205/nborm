package nborm

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"unsafe"
)

type ormType int

const (
	TypeStringField ormType = iota
	TypeIntField
	TypeFloatField
	TypeBoolField
	TypeBinaryField
	TypeDateField
	TypeDatetimeField
)

var ormTypeMap = map[string]ormType{
	"*nborm.StringField":   TypeStringField,
	"*nborm.IntField":      TypeIntField,
	"*nborm.FloatField":    TypeFloatField,
	"*nborm.BoolField":     TypeBoolField,
	"*nborm.BinaryField":   TypeBinaryField,
	"*nborm.DateField":     TypeDateField,
	"*nborm.DatetimeField": TypeDatetimeField,
}

type columnInfo struct {
	ormType  ormType
	colName  string
	nullable bool
	isInc    bool
	isPk     bool
	isUni    bool
	defVal   interface{}
	offset   uintptr
}

type oneToOneInfo struct {
	srcCol string
	dstDB  string
	dstTab string
	dstCol string
	offset uintptr
}

type foreignKeyInfo struct {
	srcCol string
	dstDB  string
	dstTab string
	dstCol string
	offset uintptr
}

type reverseForeignKeyInfo struct {
	srcCol string
	dstDB  string
	dstTab string
	dstCol string
	offset uintptr
}

type manyToManyInfo struct {
	srcCol      string
	midDB       string
	midTab      string
	midLeftCol  string
	midRightCol string
	dstDB       string
	dstTab      string
	dstCol      string
	offset      uintptr
}

type tableInfo struct {
	db                 string
	tab                string
	modelType          reflect.Type
	columns            []*columnInfo
	columnMap          map[string]*columnInfo
	oneToOnes          []*oneToOneInfo
	foreignKeys        []*foreignKeyInfo
	reverseForeignKeys []*reverseForeignKeyInfo
	manyToManys        []*manyToManyInfo
	inc                *columnInfo
	pks                []*columnInfo
	unis               []*columnInfo
	syncFlag           uintptr
}

type databaseInfo struct {
	tables   []*tableInfo
	tableMap map[string]*tableInfo
}

type schemaInfo struct {
	databases   []*databaseInfo
	databaseMap map[string]*databaseInfo
}

var schemaCache = schemaInfo{make([]*databaseInfo, 0, 16), make(map[string]*databaseInfo)}
var schemaLock sync.RWMutex

var sqlConstRe = regexp.MustCompile(`^@'(.+)'$`)
var sqlStringRe = regexp.MustCompile(`^s'(.*)'$`)
var sqlIntRe = regexp.MustCompile(`^d'(\d+)'$`)
var sqlFloatRe = regexp.MustCompile(`^f'(.*)'$`)
var sqlBoolRe = regexp.MustCompile(`^b'(true|false)'$`)
var sqlBinaryRe = regexp.MustCompile(`^x'.*'$`)

func parseColumn(field reflect.StructField) *columnInfo {
	c := &columnInfo{}
	c.offset = field.Offset
	var ok bool
	c.ormType, ok = ormTypeMap[field.Type.String()]
	if !ok {
		panic(fmt.Errorf("nborm.parseColumn() error: unsupported field type (%s)", field.Type.Name()))
	}
	if c.colName, ok = field.Tag.Lookup("column"); !ok {
		c.colName = toSnakeCase(field.Name)
	}
	if nullable, ok := field.Tag.Lookup("nullable"); ok && nullable == "true" {
		c.nullable = true
	}
	if isInc, ok := field.Tag.Lookup("auto_increment"); ok && c.ormType == TypeIntField && isInc == "true" {
		c.isInc = true
	}
	if isPk, ok := field.Tag.Lookup("primary_key"); ok && isPk == "true" {
		c.isPk = true
	}
	if isUni, ok := field.Tag.Lookup("unique"); ok && isUni == "true" {
		c.isUni = true
	}
	if defVal, ok := field.Tag.Lookup("default_value"); ok {
		if group := sqlConstRe.FindStringSubmatch(defVal); len(group) > 1 {
			c.defVal = group[1]
		} else {
			var err error
			errInvalidDefaultValue := fmt.Errorf("nborm.parseColumn() error: invalid default value (%s) for field (%s)", defVal, field.Name)
			switch c.ormType {
			case TypeStringField, TypeDateField, TypeDatetimeField:
				if group := sqlStringRe.FindStringSubmatch(defVal); len(group) > 1 {
					c.defVal = fmt.Sprintf("%q", group[1])
				} else {
					panic(errInvalidDefaultValue)
				}
			case TypeIntField:
				if group := sqlIntRe.FindStringSubmatch(defVal); len(group) > 1 {
					if c.defVal, err = strconv.ParseInt(defVal, 10, 64); err != nil {
						panic(err)
					}
				} else {
					panic(errInvalidDefaultValue)
				}
			case TypeFloatField:
				if group := sqlFloatRe.FindStringSubmatch(defVal); len(group) > 1 {
					if c.defVal, err = strconv.ParseFloat(defVal, 64); err != nil {
						panic(err)
					}
				} else {
					panic(errInvalidDefaultValue)
				}
			case TypeBoolField:
				if group := sqlBoolRe.FindStringSubmatch(defVal); len(group) > 1 {
					if c.defVal, err = strconv.ParseBool(defVal); err != nil {
						panic(err)
					}
				} else {
					panic(errInvalidDefaultValue)
				}
			case TypeBinaryField:
				if sqlBinaryRe.MatchString(defVal) {
					c.defVal = defVal
				} else {
					panic(errInvalidDefaultValue)
				}
			}
		}
	}
	return c
}

func getSrcCol(field reflect.StructField) string {
	if srcCol, ok := field.Tag.Lookup("source_column"); !ok {
		panic(fmt.Errorf("nborm.getSrcCol() error: no source column tag for field (%s)", field.Name))
	} else {
		return srcCol
	}
}

func getDstCol(field reflect.StructField) (dstDB, dstTab, dstCol string) {
	if dstCol, ok := field.Tag.Lookup("destination_column"); !ok {
		panic(fmt.Errorf("nborm.getDstCol() error: no destination column tag for field (%s)", field.Name))
	} else {
		l := strings.Split(dstCol, ".")
		if len(l) != 3 {
			panic(fmt.Errorf("nborm.getDstCol() error: invalid destination column tag (%s) for field (%s)", dstCol, field.Name))
		}
		return l[0], l[1], l[2]
	}
}

func getMidCol(field reflect.StructField) (midDB, midTab, midLeftCol, midRightCol string) {
	var ok bool
	if midDB, ok = field.Tag.Lookup("middle_database"); !ok {
		panic(fmt.Errorf("nborm.getMidCol() error: no middle database tag for field (%s)", field.Name))
	}
	if midTab, ok = field.Tag.Lookup("middle_table"); !ok {
		panic(fmt.Errorf("nborm.getMidCol() error: no middle table tag for field (%s)", field.Name))
	}
	if midLeftCol, ok = field.Tag.Lookup("middle_left_column"); !ok {
		panic(fmt.Errorf("nborm.getMidCol() error: no middle left column tag for field (%s)", field.Name))
	}
	if midRightCol, ok = field.Tag.Lookup("middle_right_column"); !ok {
		panic(fmt.Errorf("nborm.getMidCol() error: no middle right column tag for field (%s)", field.Name))
	}
	return
}

func parseOneToOne(field reflect.StructField) *oneToOneInfo {
	o := &oneToOneInfo{}
	o.srcCol = getSrcCol(field)
	o.dstDB, o.dstTab, o.dstCol = getDstCol(field)
	o.offset = field.Offset
	return o
}

func parseForeignKey(field reflect.StructField) *foreignKeyInfo {
	f := &foreignKeyInfo{}
	f.srcCol = getSrcCol(field)
	f.dstDB, f.dstTab, f.dstCol = getDstCol(field)
	f.offset = field.Offset
	return f
}

func parseReverseForeignKey(field reflect.StructField) *reverseForeignKeyInfo {
	r := &reverseForeignKeyInfo{}
	r.srcCol = getSrcCol(field)
	r.dstDB, r.dstTab, r.dstCol = getDstCol(field)
	r.offset = field.Offset
	return r
}

func parseManyToMany(field reflect.StructField) *manyToManyInfo {
	m := &manyToManyInfo{}
	m.srcCol = getSrcCol(field)
	m.midDB, m.midTab, m.midLeftCol, m.midRightCol = getMidCol(field)
	m.dstDB, m.dstTab, m.dstCol = getDstCol(field)
	m.offset = field.Offset
	return m
}

func parseTable(table table) *tableInfo {
	dbName, tabName := table.DB(), table.Tab()
	tabInfo := &tableInfo{db: dbName, tab: tabName, columnMap: make(map[string]*columnInfo)}
	typ := reflect.TypeOf(table)
	if typ.Kind() != reflect.Ptr {
		panic(fmt.Errorf("nborm.parseTable() error: require a pointer but supported (%s)", typ.Name()))
	}
	elem := typ.Elem()
	kind := elem.Kind()
	var stct reflect.Type
	switch kind {
	case reflect.Struct:
		stct = elem
	case reflect.Array, reflect.Slice:
		sElem := elem.Elem()
		if sElem.Kind() != reflect.Ptr {
			panic(fmt.Errorf("nborm.parseTable() error: require a array of struct pointer (%s)", typ.Name()))
		}
		stct = sElem.Elem()
		if stct.Kind() != reflect.Struct {
			panic(fmt.Errorf("nborm.parseTable() error: the type of element of slice(array) is not a struct pointer (%s)", typ.Name()))
		}
	default:
		panic(fmt.Errorf("nborm.parseTable() error: object must be a struct or array of struct pointer (%s)", typ.Name()))
	}
	tabInfo.modelType = stct
	for i := 0; i < stct.NumField(); i++ {
		field := stct.Field(i)
		switch field.Type.String() {
		case "*nborm.StringField", "*nborm.IntField", "*nborm.FloatField", "*nborm.BoolField", "*nborm.BinaryField", "*nborm.DateField",
			"*nborm.DatetimeField":
			colInfo := parseColumn(field)
			tabInfo.columnMap[colInfo.colName] = colInfo
			tabInfo.columns = append(tabInfo.columns, colInfo)
			if colInfo.isInc {
				if tabInfo.inc != nil {
					panic(fmt.Errorf("nborm.parseTable() error: multiple auto increment columns (%s)", typ.Name()))
				}
				tabInfo.inc = colInfo
			}
			if colInfo.isPk {
				tabInfo.pks = append(tabInfo.pks, colInfo)
			}
			if colInfo.isUni {
				tabInfo.unis = append(tabInfo.unis, colInfo)
			}
		case "*nborm.OneToOne":
			otoInfo := parseOneToOne(field)
			tabInfo.oneToOnes = append(tabInfo.oneToOnes, otoInfo)
		case "*nborm.ForeignKey":
			fkInfo := parseForeignKey(field)
			tabInfo.foreignKeys = append(tabInfo.foreignKeys, fkInfo)
		case "*nborm.ReverseForeignKey":
			rfkInfo := parseReverseForeignKey(field)
			tabInfo.reverseForeignKeys = append(tabInfo.reverseForeignKeys, rfkInfo)
		case "*nborm.ManyToMany":
			mtmInfo := parseManyToMany(field)
			tabInfo.manyToManys = append(tabInfo.manyToManys, mtmInfo)
		case "nborm.SyncFlag":
			if tabInfo.syncFlag != 0 {
				panic(fmt.Errorf("nborm.parseTable() error: multiple SyncFlag (%s)", typ.Name()))
			}
			tabInfo.syncFlag = field.Offset
		}
	}
	if tabInfo.pks == nil {
		panic(fmt.Errorf("nborm.parseTable() error: no primary key in %s.%s", dbName, tabName))
	}
	if tabInfo.syncFlag == 0 {
		panic(fmt.Errorf("nborm.parseTable() error: no SyncFlag in %s.%s", dbName, tabName))
	}
	return tabInfo
}

func getTabInfo(table table) *tableInfo {
	db, tab := table.DB(), table.Tab()
	schemaLock.RLock()
	dbInfo, ok := schemaCache.databaseMap[db]
	if !ok {
		schemaLock.RUnlock()
		tInfo := parseTable(table)
		dInfo := &databaseInfo{[]*tableInfo{tInfo}, map[string]*tableInfo{tab: tInfo}}
		schemaLock.Lock()
		schemaCache.databases = append(schemaCache.databases, dInfo)
		schemaCache.databaseMap[db] = dInfo
		schemaLock.Unlock()
		return tInfo
	}
	tabInfo, ok := dbInfo.tableMap[tab]
	if !ok {
		schemaLock.RUnlock()
		tInfo := parseTable(table)
		schemaLock.Lock()
		dbInfo.tables = append(dbInfo.tables, tInfo)
		dbInfo.tableMap[tab] = tInfo
		schemaLock.Unlock()
		return tInfo
	}
	schemaLock.RUnlock()
	return tabInfo
}

func getTabInfoByName(db, tab string) *tableInfo {
	schemaLock.RLock()
	defer schemaLock.RUnlock()
	if dbInfo, ok := schemaCache.databaseMap[db]; !ok {
		panic(fmt.Errorf("nborm.getTabInfoByName() error: database not exists (%s)", db))
	} else {
		if tabInfo, ok := dbInfo.tableMap[tab]; !ok {
			panic(fmt.Errorf("nborm.getTabInfoByName() error: table not exists (%s.%s)", db, tab))
		} else {
			return tabInfo
		}
	}
}

func initModelWithTableInfo(model table, tabInfo *tableInfo) {
	db, tab := model.DB(), model.Tab()
	baseAddr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&model)) + uintptr(8)))
	for _, col := range tabInfo.columns {
		switch col.ormType {
		case TypeStringField:
			*(**StringField)(unsafe.Pointer(baseAddr + col.offset)) = NewStringField(db, tab, col.colName, col.nullable, col.isPk, col.isUni,
				col.defVal, col.offset)
		case TypeIntField:
			*(**IntField)(unsafe.Pointer(baseAddr + col.offset)) = NewIntField(db, tab, col.colName, col.nullable, col.isPk, col.isInc, col.isUni,
				col.defVal, col.offset)
		case TypeFloatField:
			*(**FloatField)(unsafe.Pointer(baseAddr + col.offset)) = NewFloatField(db, tab, col.colName, col.nullable, col.isPk, col.isUni,
				col.defVal, col.offset)
		case TypeBoolField:
			*(**BoolField)(unsafe.Pointer(baseAddr + col.offset)) = NewBoolField(db, tab, col.colName, col.nullable, col.isPk, col.isUni,
				col.defVal, col.offset)
		case TypeBinaryField:
			*(**BinaryField)(unsafe.Pointer(baseAddr + col.offset)) = NewBinaryField(db, tab, col.colName, col.nullable, col.isPk, col.isUni,
				col.defVal, col.offset)
		case TypeDateField:
			*(**DateField)(unsafe.Pointer(baseAddr + col.offset)) = NewDateField(db, tab, col.colName, col.nullable, col.isPk, col.isUni,
				col.defVal, col.offset)
		case TypeDatetimeField:
			*(**DatetimeField)(unsafe.Pointer(baseAddr + col.offset)) = NewDatetimeField(db, tab, col.colName, col.nullable, col.isPk, col.isUni,
				col.defVal, col.offset)
		}
	}
	for _, oto := range tabInfo.oneToOnes {
		srcField := getFieldByName(baseAddr, oto.srcCol, tabInfo)
		*(**OneToOne)(unsafe.Pointer(baseAddr + oto.offset)) = NewOneToOne(db, tab, oto.srcCol, oto.dstDB, oto.dstTab, oto.dstCol, srcField.value)
	}
	for _, fk := range tabInfo.foreignKeys {
		srcField := getFieldByName(baseAddr, fk.srcCol, tabInfo)
		*(**ForeignKey)(unsafe.Pointer(baseAddr + fk.offset)) = NewForeignKey(db, tab, fk.srcCol, fk.dstDB, fk.dstTab, fk.dstCol, srcField.value)
	}
	for _, rfk := range tabInfo.reverseForeignKeys {
		srcField := getFieldByName(baseAddr, rfk.srcCol, tabInfo)
		*(**ReverseForeignKey)(unsafe.Pointer(baseAddr + rfk.offset)) = NewReverseForeignKey(db, tab, rfk.srcCol, rfk.dstDB, rfk.dstTab, rfk.dstCol,
			srcField.value)
	}
	for _, mtm := range tabInfo.manyToManys {
		srcField := getFieldByName(baseAddr, mtm.srcCol, tabInfo)
		*(**ManyToMany)(unsafe.Pointer(baseAddr + mtm.offset)) = NewManyToMany(db, tab, mtm.srcCol, mtm.midDB, mtm.midTab, mtm.midLeftCol,
			mtm.midRightCol, mtm.dstDB, mtm.dstTab, mtm.dstCol, srcField.value)
	}
}

func newModelAddr(tabInfo *tableInfo) uintptr {
	model := reflect.New(tabInfo.modelType)
	initModelWithTableInfo(model.Interface().(Model), tabInfo)
	return model.Pointer()
}

func newModel(tabInfo *tableInfo) Model {
	model := reflect.New(tabInfo.modelType).Interface().(Model)
	initModelWithTableInfo(model, tabInfo)
	return model
}

func InitModel(model Model) {
	tabInfo := getTabInfo(model)
	initModelWithTableInfo(model, tabInfo)
}

func InitSlice(table table) {
	tabInfo := getTabInfo(table)
	sliceAddr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&table)) + uintptr(8)))
	*(*int)(unsafe.Pointer(sliceAddr + uintptr(8))) = 1
	underArrayAddr := *(*uintptr)(unsafe.Pointer(sliceAddr))
	*(*uintptr)(unsafe.Pointer(underArrayAddr)) = newModelAddr(tabInfo)
}

// func getFieldByName(model table, colName string, tabInfo *tableInfo) Field {
// 	colInfo, ok := tabInfo.columnMap[colName]
// 	if !ok {
// 		panic(fmt.Errorf("nborm.getFieldByName() error: %T.%s column not exist", model, colName))
// 	}
// 	baseAddr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&model)) + uintptr(8)))
// 	switch colInfo.ormType {
// 	case TypeStringField:
// 		return *(**StringField)(unsafe.Pointer(baseAddr + colInfo.offset))
// 	case TypeIntField:
// 		return *(**IntField)(unsafe.Pointer(baseAddr + colInfo.offset))
// 	case TypeFloatField:
// 		return *(**FloatField)(unsafe.Pointer(baseAddr + colInfo.offset))
// 	case TypeBoolField:
// 		return *(**BoolField)(unsafe.Pointer(baseAddr + colInfo.offset))
// 	case TypeBinaryField:
// 		return *(**BinaryField)(unsafe.Pointer(baseAddr + colInfo.offset))
// 	case TypeDateField:
// 		return *(**DateField)(unsafe.Pointer(baseAddr + colInfo.offset))
// 	case TypeDatetimeField:
// 		return *(**DatetimeField)(unsafe.Pointer(baseAddr + colInfo.offset))
// 	default:
// 		panic(fmt.Errorf("nborm.getFieldByName() error: unknown field type (%d)", colInfo.ormType))
// 	}
// }

func getFieldByName(addr uintptr, colName string, tabInfo *tableInfo) Field {
	colInfo, ok := tabInfo.columnMap[colName]
	if !ok {
		panic(fmt.Errorf("nborm.getFieldByName() error: %s.%s.%s column not exist", tabInfo.db, tabInfo.tab, colName))
	}
	switch colInfo.ormType {
	case TypeStringField:
		return *(**StringField)(unsafe.Pointer(addr + colInfo.offset))
	case TypeIntField:
		return *(**IntField)(unsafe.Pointer(addr + colInfo.offset))
	case TypeFloatField:
		return *(**FloatField)(unsafe.Pointer(addr + colInfo.offset))
	case TypeBoolField:
		return *(**BoolField)(unsafe.Pointer(addr + colInfo.offset))
	case TypeBinaryField:
		return *(**BinaryField)(unsafe.Pointer(addr + colInfo.offset))
	case TypeDateField:
		return *(**DateField)(unsafe.Pointer(addr + colInfo.offset))
	case TypeDatetimeField:
		return *(**DatetimeField)(unsafe.Pointer(addr + colInfo.offset))
	default:
		panic(fmt.Errorf("nborm.getFieldByName() error: unknown field type (%d)", colInfo.ormType))
	}
}

func getFieldByColumnInfo(addr uintptr, colInfo *columnInfo) Field {
	fieldAddr := unsafe.Pointer(addr + colInfo.offset)
	switch colInfo.ormType {
	case TypeStringField:
		return *(**StringField)(fieldAddr)
	case TypeIntField:
		return *(**IntField)(fieldAddr)
	case TypeFloatField:
		return *(**FloatField)(fieldAddr)
	case TypeBoolField:
		return *(**BoolField)(fieldAddr)
	case TypeBinaryField:
		return *(**BinaryField)(fieldAddr)
	case TypeDateField:
		return *(**DateField)(fieldAddr)
	case TypeDatetimeField:
		return *(**DatetimeField)(fieldAddr)
	default:
		panic(fmt.Errorf("nborm.getFieldByColumnInfo() error: unknown field type (%d)", colInfo.ormType))
	}

}

// func getPks(model Model) []Field {
// 	tabInfo := getTabInfo(model)
// 	l := make([]Field, len(tabInfo.pks))
// 	for i, pkCol := range tabInfo.pks {
// 		l[i] = getFieldByName(model, pkCol.colName, tabInfo)
// 	}
// 	return l
// }

func getPksWithTableInfo(addr uintptr, info *tableInfo) []Field {
	l := make([]Field, len(info.pks))
	for i, pkCol := range info.pks {
		l[i] = getFieldByColumnInfo(addr, pkCol)
	}
	return l
}

func getPksAndOthers(model Model) (pks []Field, others []Field) {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	for _, col := range tabInfo.columns {
		if col.isPk {
			pks = append(pks, getFieldByName(modAddr, col.colName, tabInfo))
		} else {
			others = append(others, getFieldByName(modAddr, col.colName, tabInfo))
		}

	}
	return
}

func getInc(model Model) Field {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	if tabInfo.inc == nil {
		return nil
	}
	return getFieldByName(modAddr, tabInfo.inc.colName, tabInfo)
}

func getIncWithTableInfo(addr uintptr, info *tableInfo) Field {
	if info.inc == nil {
		return nil
	}
	return *(**IntField)(unsafe.Pointer(addr + info.inc.offset))
}

func getIncAndOthers(addr uintptr, tabInfo *tableInfo) (inc Field, others []Field) {
	for _, col := range tabInfo.columns {
		if col.isInc {
			inc = getFieldByColumnInfo(addr, col)
		} else {
			others = append(others, getFieldByColumnInfo(addr, col))
		}
	}
	return
}

func getUnis(model Model) []Field {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	if tabInfo.unis == nil {
		return nil
	}
	l := make([]Field, len(tabInfo.unis))
	for i, uniCol := range tabInfo.unis {
		l[i] = getFieldByName(modAddr, uniCol.colName, tabInfo)
	}
	return l
}

func getUinsWithTableInfo(addr uintptr, tabInfo *tableInfo) []Field {
	if tabInfo.unis == nil {
		return nil
	}
	l := make([]Field, len(tabInfo.unis))
	for i, uniCol := range tabInfo.unis {
		l[i] = getFieldByColumnInfo(addr, uniCol)
	}
	return l
}

func getUnisAndOthers(model Model) (unis []Field, others []Field) {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	for _, col := range tabInfo.columns {
		if col.isUni {
			unis = append(unis, getFieldByName(modAddr, col.colName, tabInfo))
		} else {
			others = append(others, getFieldByName(modAddr, col.colName, tabInfo))
		}

	}
	return
}

func getAllFields(model Model) []interface{} {
	tabInfo := getTabInfo(model)
	baseAddr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&model)) + uintptr(8)))
	l := make([]interface{}, len(tabInfo.columns))
	for i, colInfo := range tabInfo.columns {
		fieldPointer := unsafe.Pointer(baseAddr + colInfo.offset)
		switch colInfo.ormType {
		case TypeStringField:
			l[i] = *(**StringField)(fieldPointer)
		case TypeIntField:
			l[i] = *(**IntField)(fieldPointer)
		case TypeFloatField:
			l[i] = *(**FloatField)(fieldPointer)
		case TypeBoolField:
			l[i] = *(**FloatField)(fieldPointer)
		case TypeBinaryField:
			l[i] = *(**BinaryField)(fieldPointer)
		case TypeDateField:
			l[i] = *(**DateField)(fieldPointer)
		case TypeDatetimeField:
			l[i] = *(**DatetimeField)(fieldPointer)
		}
	}
	return l
}

func getAllFieldsWithTableInfo(addr uintptr, tabInfo *tableInfo) []interface{} {
	l := make([]interface{}, len(tabInfo.columns))
	for i, colInfo := range tabInfo.columns {
		l[i] = getFieldByColumnInfo(addr, colInfo)
	}
	return l
}

func scanRow(addr uintptr, tabInfo *tableInfo, row *sql.Row) error {
	fields := getAllFieldsWithTableInfo(addr, tabInfo)
	if err := row.Scan(fields...); err != nil {
		return err
	}
	setSync(addr, tabInfo)
	return nil
}

func scanRows(addr uintptr, tabInfo *tableInfo, rows *sql.Rows) error {
	defer rows.Close()
	lAddr := (*[]uintptr)(unsafe.Pointer(addr))
	for rows.Next() {
		modelAddr := newModelAddr(tabInfo)
		if err := rows.Scan(getAllFieldsWithTableInfo(modelAddr, tabInfo)...); err != nil {
			return err
		}
		setSync(addr, tabInfo)
		*lAddr = append(*lAddr, *(*uintptr)(unsafe.Pointer(&modelAddr)))
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
