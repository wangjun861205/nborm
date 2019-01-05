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
	"nborm.StringField":   TypeStringField,
	"nborm.IntField":      TypeIntField,
	"nborm.FloatField":    TypeFloatField,
	"nborm.BoolField":     TypeBoolField,
	"nborm.BinaryField":   TypeBinaryField,
	"nborm.DateField":     TypeDateField,
	"nborm.DatetimeField": TypeDatetimeField,
}

type columnInfo struct {
	ormType   ormType
	colName   string
	fieldName string
	nullable  bool
	isInc     bool
	isPk      bool
	isUni     bool
	defVal    interface{}
	offset    uintptr
	sqlType   string
	charset   string
	collate   string
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

type primaryKey []*columnInfo

func (pk primaryKey) match(fields ...Field) bool {
	for _, pkCol := range pk {
		get := false
		for _, field := range fields {
			if field.columnName() == pkCol.colName {
				get = true
				break
			}
		}
		if !get {
			return false
		}
	}
	return true
}

func (pk primaryKey) getFields(addr uintptr) []Field {
	fields := make([]Field, len(pk))
	for i, pkCol := range pk {
		fields[i] = getFieldByColumnInfo(addr, pkCol)
	}
	return fields
}

func (pk primaryKey) genCreateClause() string {
	l := make([]string, len(pk))
	for i, col := range pk {
		l[i] = wrap(col.colName)
	}
	return fmt.Sprintf(", PRIMARY KEY (%s)", strings.Join(l, ", "))
}

func (pk primaryKey) methodString() string {
	l := make([]string, len(pk))
	for i, col := range pk {
		l[i] = fmt.Sprintf("%q", col.colName)
	}
	return fmt.Sprintf("return []string{ %s }", strings.Join(l, ", "))
}

type uniqueKeys [][]*columnInfo

func (uks uniqueKeys) match(fields ...Field) (int, bool) {
	for i, uk := range uks {
		fit := true
		for _, col := range uk {
			get := false
			for _, field := range fields {
				if col.colName == field.columnName() {
					get = true
					break
				}
			}
			if !get {
				fit = false
				break
			}
		}
		if fit {
			return i, true
		}
	}
	return -1, false
}

func (uks uniqueKeys) getFields(addr uintptr, ukIndex int) []Field {
	uk := uks[ukIndex]
	fields := make([]Field, len(uk))
	for i, ukCol := range uk {
		fields[i] = getFieldByColumnInfo(addr, ukCol)
	}
	return fields
}

func (uks uniqueKeys) genCreateClause() string {
	keyList := make([]string, len(uks))
	for i, key := range uks {
		colList := make([]string, len(key))
		for j, col := range key {
			colList[j] = wrap(col.colName)
		}
		keyList[i] = fmt.Sprintf(", UNIQUE KEY (%s)", strings.Join(colList, ", "))
	}
	return strings.Join(keyList, ", ")
}

func (uks uniqueKeys) methodString() string {
	keyList := make([]string, len(uks))
	for i, key := range uks {
		colList := make([]string, len(key))
		for j, col := range key {
			colList[j] = fmt.Sprintf("%q", col.colName)
		}
		keyList[i] = fmt.Sprintf("[]string{ %s },\n", strings.Join(colList, ", "))
	}
	return fmt.Sprintf(`[][]string{ 
		%s 
	}`, strings.Join(keyList, ", "))
}

type keys [][]*columnInfo

func (ks keys) genCreateClause() string {
	keyList := make([]string, len(ks))
	for i, key := range ks {
		colList := make([]string, len(key))
		for j, col := range key {
			colList[j] = wrap(col.colName)
		}
		keyList[i] = fmt.Sprintf(", KEY (%s)", strings.Join(colList, ", "))
	}
	return strings.Join(keyList, ", ")
}

func (ks keys) methodString() string {
	keyList := make([]string, len(ks))
	for i, key := range ks {
		colList := make([]string, len(key))
		for j, col := range key {
			colList[j] = fmt.Sprintf("%q", col.colName)
		}
		keyList[i] = fmt.Sprintf("[]string{ %s },\n", strings.Join(colList, ", "))
	}
	return fmt.Sprintf(`[][]string{ 
		%s 
	}`, strings.Join(keyList, ", "))
}

type tableInfo struct {
	db                 string
	tab                string
	modelType          reflect.Type
	modelName          string
	columns            []*columnInfo
	columnMap          map[string]*columnInfo
	oneToOnes          []*oneToOneInfo
	foreignKeys        []*foreignKeyInfo
	reverseForeignKeys []*reverseForeignKeyInfo
	manyToManys        []*manyToManyInfo
	inc                *columnInfo
	pk                 primaryKey
	pkNames            []string
	unis               uniqueKeys
	uniNames           [][]string
	keys               keys
	keyNames           [][]string
	modelStatus        uintptr
	isComplete         bool
	isNewMiddleTable   bool
}

func (ti *tableInfo) lookupColInfoByFieldName(fieldName string) *columnInfo {
	for _, colInfo := range ti.columns {
		if colInfo.fieldName == fieldName {
			return colInfo
		}
	}
	return nil
}

type databaseInfo struct {
	tables   []*tableInfo
	tableMap map[string]*tableInfo
}

func (si *databaseInfo) getOrCreate(tabName string) *tableInfo {
	if info, ok := si.tableMap[tabName]; ok {
		return info
	} else {
		info = &tableInfo{
			columns:            make([]*columnInfo, 0, 32),
			columnMap:          make(map[string]*columnInfo),
			oneToOnes:          make([]*oneToOneInfo, 0, 8),
			foreignKeys:        make([]*foreignKeyInfo, 0, 8),
			reverseForeignKeys: make([]*reverseForeignKeyInfo, 0, 8),
			manyToManys:        make([]*manyToManyInfo, 0, 8),
			pk:                 make(primaryKey, 0, 8),
			unis:               make(uniqueKeys, 0, 8),
			keys:               make(keys, 0, 8),
		}
		si.tables = append(si.tables, info)
		si.tableMap[tabName] = info
		return info
	}
}

type schemaInfo struct {
	databases   []*databaseInfo
	databaseMap map[string]*databaseInfo
}

func (si *schemaInfo) getOrCreate(dbName string) *databaseInfo {
	if db, ok := si.databaseMap[dbName]; ok {
		return db
	} else {
		dbInfo := &databaseInfo{make([]*tableInfo, 0, 16), make(map[string]*tableInfo)}
		si.databases = append(si.databases, dbInfo)
		si.databaseMap[dbName] = dbInfo
		return dbInfo
	}
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
	if colName, ok := field.Tag.Lookup("column"); !ok {
		c.colName = wrap(toSnakeCase(field.Name))
	} else {
		c.colName = wrap(colName)
	}
	// if c.colName, ok = field.Tag.Lookup("column"); !ok {
	// 	c.colName = toSnakeCase(field.Name)
	// }
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
					if c.defVal, err = strconv.ParseInt(group[1], 10, 64); err != nil {
						panic(err)
					}
				} else {
					panic(errInvalidDefaultValue)
				}
			case TypeFloatField:
				if group := sqlFloatRe.FindStringSubmatch(defVal); len(group) > 1 {
					if c.defVal, err = strconv.ParseFloat(group[1], 64); err != nil {
						panic(err)
					}
				} else {
					panic(errInvalidDefaultValue)
				}
			case TypeBoolField:
				if group := sqlBoolRe.FindStringSubmatch(defVal); len(group) > 1 {
					if c.defVal, err = strconv.ParseBool(group[1]); err != nil {
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

// func getSrcCol(field reflect.StructField) string {
// 	if srcCol, ok := field.Tag.Lookup("source_column"); !ok {
// 		panic(fmt.Errorf("nborm.getSrcCol() error: no source column tag for field (%s)", field.Name))
// 	} else {
// 		return srcCol
// 	}
// }

func getSrcCol(field reflect.StructField) string {
	if srcCol, ok := field.Tag.Lookup("src_col"); !ok {
		panic(fmt.Errorf("nborm.getSrcCol() error: no source column tag for field (%s)", field.Name))
	} else {
		return wrap(srcCol)
	}
}

// func getDstCol(field reflect.StructField) (dstDB, dstTab, dstCol string) {
// 	if dstCol, ok := field.Tag.Lookup("destination_column"); !ok {
// 		panic(fmt.Errorf("nborm.getDstCol() error: no destination column tag for field (%s)", field.Name))
// 	} else {
// 		l := strings.Split(dstCol, ".")
// 		if len(l) != 3 {
// 			panic(fmt.Errorf("nborm.getDstCol() error: invalid destination column tag (%s) for field (%s)", dstCol, field.Name))
// 		}
// 		return l[0], l[1], l[2]
// 	}
// }

func getDstCol(field reflect.StructField) (dstDB, dstTab, dstCol string) {
	var ok bool
	if dstDB, ok = field.Tag.Lookup("dst_db"); !ok {
		panic(fmt.Errorf("nborm.getDstCol() error: no destination database tag for field (%s)", field.Name))
	}
	if dstTab, ok = field.Tag.Lookup("dst_tab"); !ok {
		panic(fmt.Errorf("nborm.getDstCol() error: no destination table tag for field (%s)", field.Name))
	}
	if dstCol, ok = field.Tag.Lookup("dst_col"); !ok {
		panic(fmt.Errorf("nborm.getDstCol() error: no destination column tag for field (%s)", field.Name))
	}
	dstDB, dstTab, dstCol = wrap(dstDB), wrap(dstTab), wrap(dstCol)
	return
}

// func getMidCol(field reflect.StructField) (midDB, midTab, midLeftCol, midRightCol string) {
// 	var ok bool
// 	if midDB, ok = field.Tag.Lookup("middle_database"); !ok {
// 		panic(fmt.Errorf("nborm.getMidCol() error: no middle database tag for field (%s)", field.Name))
// 	}
// 	if midTab, ok = field.Tag.Lookup("middle_table"); !ok {
// 		panic(fmt.Errorf("nborm.getMidCol() error: no middle table tag for field (%s)", field.Name))
// 	}
// 	if midLeftCol, ok = field.Tag.Lookup("middle_left_column"); !ok {
// 		panic(fmt.Errorf("nborm.getMidCol() error: no middle left column tag for field (%s)", field.Name))
// 	}
// 	if midRightCol, ok = field.Tag.Lookup("middle_right_column"); !ok {
// 		panic(fmt.Errorf("nborm.getMidCol() error: no middle right column tag for field (%s)", field.Name))
// 	}
// 	return
// }

func getMidCol(field reflect.StructField) (midDB, midTab, midLeftCol, midRightCol string) {
	var ok bool
	if midDB, ok = field.Tag.Lookup("mid_db"); !ok {
		panic(fmt.Errorf("nborm.getMidCol() error: no middle database tag for field (%s)", field.Name))
	}
	if midTab, ok = field.Tag.Lookup("mid_tab"); !ok {
		panic(fmt.Errorf("nborm.getMidCol() error: no middle table tag for field (%s)", field.Name))
	}
	if midLeftCol, ok = field.Tag.Lookup("mid_left_col"); !ok {
		panic(fmt.Errorf("nborm.getMidCol() error: no middle left column tag for field (%s)", field.Name))
	}
	if midRightCol, ok = field.Tag.Lookup("mid_right_col"); !ok {
		panic(fmt.Errorf("nborm.getMidCol() error: no middle right column tag for field (%s)", field.Name))
	}
	midDB, midTab, midLeftCol, midRightCol = wrap(midDB), wrap(midTab), wrap(midLeftCol), wrap(midRightCol)
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
	dbName, tabName := wrap(table.DB()), wrap(table.Tab())
	tabInfo := &tableInfo{db: dbName, tab: tabName, columnMap: make(map[string]*columnInfo)}
	var haveModelStatus bool
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
		case "nborm.StringField", "nborm.IntField", "nborm.FloatField", "nborm.BoolField", "nborm.BinaryField", "nborm.DateField",
			"nborm.DatetimeField":
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
				tabInfo.pk = append(tabInfo.pk, colInfo)
			}
			// if colInfo.isUni {
			// 	tabInfo.unis = append(tabInfo.unis, colInfo)
			// }
		case "nborm.OneToOne":
			otoInfo := parseOneToOne(field)
			tabInfo.oneToOnes = append(tabInfo.oneToOnes, otoInfo)
		case "nborm.ForeignKey":
			fkInfo := parseForeignKey(field)
			tabInfo.foreignKeys = append(tabInfo.foreignKeys, fkInfo)
		case "nborm.ReverseForeignKey":
			rfkInfo := parseReverseForeignKey(field)
			tabInfo.reverseForeignKeys = append(tabInfo.reverseForeignKeys, rfkInfo)
		case "nborm.ManyToMany":
			mtmInfo := parseManyToMany(field)
			tabInfo.manyToManys = append(tabInfo.manyToManys, mtmInfo)
		case "nborm.ModelStatus":
			if haveModelStatus {
				panic(fmt.Errorf("nborm.parseTable() error: multiple ModelStatus field (%s)", typ.Name()))
			}
			tabInfo.modelStatus = field.Offset
			haveModelStatus = true
		}
	}
	if len(tabInfo.pk) == 0 {
		panic(fmt.Errorf("nborm.parseTable() error: no primary key in %s.%s", dbName, tabName))
	}
	if !haveModelStatus {
		panic(fmt.Errorf("nborm.parseTable() error: no ModelStatus field in %s.%s", dbName, tabName))
	}
	return tabInfo
}

func getTabInfo(table table) *tableInfo {
	db, tab := wrap(table.DB()), wrap(table.Tab())
	schemaLock.RLock()
	dbInfo, ok := schemaCache.databaseMap[db]
	if !ok {
		schemaLock.RUnlock()
		tInfo := parseTable(table)
		dInfo := &databaseInfo{tables: []*tableInfo{tInfo}, tableMap: map[string]*tableInfo{tab: tInfo}}
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
	if dbInfo, ok := schemaCache.databaseMap[wrap(db)]; !ok {
		panic(fmt.Errorf("nborm.getTabInfoByName() error: database not exists (%s)", db))
	} else {
		if tabInfo, ok := dbInfo.tableMap[wrap(tab)]; !ok {
			panic(fmt.Errorf("nborm.getTabInfoByName() error: table not exists (%s.%s)", db, tab))
		} else {
			return tabInfo
		}
	}
}

func initModelWithTableInfo(model table, tabInfo *tableInfo) {
	db, tab := wrap(model.DB()), wrap(model.Tab())
	baseAddr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&model)) + uintptr(8)))
	for _, col := range tabInfo.columns {
		switch col.ormType {
		case TypeStringField:
			field := (*StringField)(unsafe.Pointer(baseAddr + col.offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset = db, tab, col.colName,
				col.nullable, col.isPk, col.isUni, col.defVal, col.offset
		case TypeIntField:
			field := (*IntField)(unsafe.Pointer(baseAddr + col.offset))
			field.db, field.tab, field.column, field.nullable, field.inc, field.pk, field.uni, field.defVal, field.offset = db, tab, col.colName,
				col.nullable, col.isInc, col.isPk, col.isUni, col.defVal, col.offset
		case TypeFloatField:
			field := (*FloatField)(unsafe.Pointer(baseAddr + col.offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset = db, tab, col.colName,
				col.nullable, col.isPk, col.isUni, col.defVal, col.offset
		case TypeBoolField:
			field := (*BoolField)(unsafe.Pointer(baseAddr + col.offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset = db, tab, col.colName,
				col.nullable, col.isPk, col.isUni, col.defVal, col.offset
		case TypeBinaryField:
			field := (*BinaryField)(unsafe.Pointer(baseAddr + col.offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset = db, tab, col.colName,
				col.nullable, col.isPk, col.isUni, col.defVal, col.offset
		case TypeDateField:
			field := (*DateField)(unsafe.Pointer(baseAddr + col.offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset = db, tab, col.colName,
				col.nullable, col.isPk, col.isUni, col.defVal, col.offset
		case TypeDatetimeField:
			field := (*DatetimeField)(unsafe.Pointer(baseAddr + col.offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset = db, tab, col.colName,
				col.nullable, col.isPk, col.isUni, col.defVal, col.offset

		}
	}
	for _, oto := range tabInfo.oneToOnes {
		srcField := getFieldByName(baseAddr, oto.srcCol, tabInfo)
		relField := (*OneToOne)(unsafe.Pointer(baseAddr + oto.offset))
		relField.srcDB, relField.srcTab, relField.srcCol, relField.dstDB, relField.dstTab, relField.dstCol, relField.srcValF = db, tab,
			oto.srcCol, oto.dstDB, oto.dstTab, oto.dstCol, srcField.value
	}
	for _, fk := range tabInfo.foreignKeys {
		srcField := getFieldByName(baseAddr, fk.srcCol, tabInfo)
		relField := (*ForeignKey)(unsafe.Pointer(baseAddr + fk.offset))
		relField.srcDB, relField.srcTab, relField.srcCol, relField.dstDB, relField.dstTab, relField.dstCol, relField.srcValF = db, tab,
			fk.srcCol, fk.dstDB, fk.dstTab, fk.dstCol, srcField.value
	}
	for _, rfk := range tabInfo.reverseForeignKeys {
		srcField := getFieldByName(baseAddr, rfk.srcCol, tabInfo)
		relField := (*ReverseForeignKey)(unsafe.Pointer(baseAddr + rfk.offset))
		relField.srcDB, relField.srcTab, relField.srcCol, relField.dstDB, relField.dstTab, relField.dstCol, relField.srcValF = db, tab,
			rfk.srcCol, rfk.dstDB, rfk.dstTab, rfk.dstCol, srcField.value
	}
	for _, mtm := range tabInfo.manyToManys {
		srcField := getFieldByName(baseAddr, mtm.srcCol, tabInfo)
		relField := (*ManyToMany)(unsafe.Pointer(baseAddr + mtm.offset))
		relField.srcDB, relField.srcTab, relField.srcCol, relField.midDB, relField.midTab, relField.midLeftCol, relField.midRightCol,
			relField.dstDB, relField.dstTab, relField.dstCol, relField.srcValF = db, tab, mtm.srcCol, mtm.midDB, mtm.midTab, mtm.midLeftCol,
			mtm.midRightCol, mtm.dstDB, mtm.dstTab, mtm.dstCol, srcField.value
	}
	setInit(baseAddr, tabInfo)
}

func newModelAddr(tabInfo *tableInfo) uintptr {
	model := reflect.New(tabInfo.modelType)
	initModelWithTableInfo(model.Interface().(table), tabInfo)
	return model.Pointer()
}

func newModel(tabInfo *tableInfo) table {
	model := reflect.New(tabInfo.modelType).Interface().(table)
	initModelWithTableInfo(model, tabInfo)
	return model
}

func InitModel(model table) {
	tabInfo := getTabInfo(model)
	initModelWithTableInfo(model, tabInfo)
}

func InitSlice(slice table) {
	tabInfo := getTabInfo(slice)
	sliceAddr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice)) + uintptr(8)))
	*(*int)(unsafe.Pointer(sliceAddr + uintptr(8))) = 1
	underArrayAddr := *(*uintptr)(unsafe.Pointer(sliceAddr))
	*(*uintptr)(unsafe.Pointer(underArrayAddr)) = newModelAddr(tabInfo)
}

func getFieldByName(addr uintptr, colName string, tabInfo *tableInfo) Field {
	colInfo, ok := tabInfo.columnMap[colName]
	if !ok {
		panic(fmt.Errorf("nborm.getFieldByName() error: %s.%s.%s column not exist", tabInfo.db, tabInfo.tab, colName))
	}
	switch colInfo.ormType {
	case TypeStringField:
		return (*StringField)(unsafe.Pointer(addr + colInfo.offset))
	case TypeIntField:
		return (*IntField)(unsafe.Pointer(addr + colInfo.offset))
	case TypeFloatField:
		return (*FloatField)(unsafe.Pointer(addr + colInfo.offset))
	case TypeBoolField:
		return (*BoolField)(unsafe.Pointer(addr + colInfo.offset))
	case TypeBinaryField:
		return (*BinaryField)(unsafe.Pointer(addr + colInfo.offset))
	case TypeDateField:
		return (*DateField)(unsafe.Pointer(addr + colInfo.offset))
	case TypeDatetimeField:
		return (*DatetimeField)(unsafe.Pointer(addr + colInfo.offset))
	default:
		panic(fmt.Errorf("nborm.getFieldByName() error: unknown field type (%d)", colInfo.ormType))
	}
}

func getFieldByColumnInfo(addr uintptr, colInfo *columnInfo) Field {
	fieldAddr := unsafe.Pointer(addr + colInfo.offset)
	switch colInfo.ormType {
	case TypeStringField:
		return (*StringField)(fieldAddr)
	case TypeIntField:
		return (*IntField)(fieldAddr)
	case TypeFloatField:
		return (*FloatField)(fieldAddr)
	case TypeBoolField:
		return (*BoolField)(fieldAddr)
	case TypeBinaryField:
		return (*BinaryField)(fieldAddr)
	case TypeDateField:
		return (*DateField)(fieldAddr)
	case TypeDatetimeField:
		return (*DatetimeField)(fieldAddr)
	default:
		panic(fmt.Errorf("nborm.getFieldByColumnInfo() error: unknown field type (%d)", colInfo.ormType))
	}

}

// func getPksWithTableInfo(addr uintptr, info *tableInfo) []Field {
// 	l := make([]Field, len(info.pks))
// 	for i, pkCol := range info.pks {
// 		l[i] = getFieldByColumnInfo(addr, pkCol)
// 	}
// 	return l
// }

func getPrimaryKeyFieldsWithTableInfo(addr uintptr, tabInfo *tableInfo) []Field {
	validFields := getValidFieldsWithTableInfo(addr, tabInfo)
	if tabInfo.pk.match(validFields...) {
		return tabInfo.pk.getFields(addr)
	}
	return nil
}

func getIncWithTableInfo(addr uintptr, info *tableInfo) Field {
	if info.inc == nil {
		return nil
	}
	return (*IntField)(unsafe.Pointer(addr + info.inc.offset))
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

func getUniqueFieldsWithTableInfo(addr uintptr, tabInfo *tableInfo) []Field {
	if len(tabInfo.unis) == 0 {
		return nil
	}
	validFields := getValidFieldsWithTableInfo(addr, tabInfo)
	if idx, match := tabInfo.unis.match(validFields...); match {
		return tabInfo.unis.getFields(addr, idx)
	}
	return nil
}

// func getUinsWithTableInfo(addr uintptr, tabInfo *tableInfo) []Field {
// 	if tabInfo.unis == nil {
// 		return nil
// 	}

// 	l := make([][]Field, len(tabInfo.unis))
// 	for i, key := range tabInfo.unis {
// 		subL := make([]Field, len(key))
// 		for j, col := range key {
// 			subL[j] = getFieldByColumnInfo(addr, col)
// 		}
// 		// l[i] = getFieldByColumnInfo(addr, uniCol)
// 		l[i] = subL
// 	}
// 	return l
// }

// func getAllFieldsWithTableInfo(addr uintptr, tabInfo *tableInfo) []interface{} {
// 	l := make([]interface{}, len(tabInfo.columns))
// 	for i, colInfo := range tabInfo.columns {
// 		l[i] = getFieldByColumnInfo(addr, colInfo)
// 	}
// 	return l
// }

func getAllFieldsWithTableInfo(addr uintptr, tabInfo *tableInfo) []Field {
	l := make([]Field, len(tabInfo.columns))
	for i, colInfo := range tabInfo.columns {
		l[i] = getFieldByColumnInfo(addr, colInfo)
	}
	return l
}

func getValidFieldsWithTableInfo(addr uintptr, tabInfo *tableInfo) []Field {
	l := make([]Field, 0, len(tabInfo.columns))
	allFields := getAllFieldsWithTableInfo(addr, tabInfo)
	for _, field := range allFields {
		if field.IsValid() {
			l = append(l, field)
		}
	}
	return l
}

func scanRow(addr uintptr, tabInfo *tableInfo, row *sql.Row) error {
	fields := getAllFieldsWithTableInfo(addr, tabInfo)
	addrList := make([]interface{}, len(fields))
	for i, field := range fields {
		addrList[i] = field.(interface{})
	}
	if err := row.Scan(addrList); err != nil {
		return err
	}
	setSync(addr, tabInfo)
	return nil
}

func unionScanRow(addrs []uintptr, tabInfos []*tableInfo, row *sql.Row) error {
	fields := make([]interface{}, 0, 64)
	for i := 0; i < len(addrs); i++ {
		addrList := make([]interface{}, len(fields))
		for i, field := range getAllFieldsWithTableInfo(addrs[i], tabInfos[i]) {
			addrList[i] = field.(interface{})
		}
		fields = append(fields, addrList...)
	}
	if err := row.Scan(fields...); err != nil {
		return err
	}
	for i := 0; i < len(addrs); i++ {
		setSync(addrs[i], tabInfos[i])
	}
	return nil
}

func scanRows(addr uintptr, tabInfo *tableInfo, rows *sql.Rows) error {
	defer rows.Close()
	lAddr := (*[]uintptr)(unsafe.Pointer(addr))
	for rows.Next() {
		modelAddr := newModelAddr(tabInfo)
		fields := getAllFieldsWithTableInfo(modelAddr, tabInfo)
		fieldAddrList := make([]interface{}, len(fields))
		for i, field := range fields {
			fieldAddrList[i] = field.(interface{})
		}
		if err := rows.Scan(fieldAddrList...); err != nil {
			return err
		}
		setSync(modelAddr, tabInfo)
		*lAddr = append(*lAddr, *(*uintptr)(unsafe.Pointer(&modelAddr)))
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func unionScanRows(addrs []uintptr, tabInfos []*tableInfo, rows *sql.Rows) error {
	defer rows.Close()
	lAddrs := make([](*[]uintptr), len(addrs))
	for i, addr := range addrs {
		lAddrs[i] = (*[]uintptr)(unsafe.Pointer(addr))
	}
	for rows.Next() {
		modelList := make([]uintptr, len(tabInfos))
		fields := make([]interface{}, 0, 64)
		for i, tabInfo := range tabInfos {
			modelAddr := newModelAddr(tabInfo)
			fieldAddrList := make([]interface{}, len(fields))
			for i, field := range getAllFieldsWithTableInfo(modelAddr, tabInfo) {
				fieldAddrList[i] = field.(interface{})
			}
			fields = append(fields, fieldAddrList...)
			setSync(modelAddr, tabInfo)
			modelList[i] = modelAddr
			*lAddrs[i] = append(*lAddrs[i], *(*uintptr)(unsafe.Pointer(&modelList[i])))
		}
		if err := rows.Scan(fields...); err != nil {
			return err
		}
		for i, modAddr := range modelList {
			setSync(modAddr, tabInfos[i])
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
