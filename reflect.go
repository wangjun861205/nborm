package nborm

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"unsafe"
)

type OrmType int

const (
	TypeStringField OrmType = iota
	TypeIntField
	TypeFloatField
	TypeBoolField
	TypeBinaryField
	TypeDateField
	TypeDatetimeField
)

var OrmTypeMap = map[string]OrmType{
	"StringField":   TypeStringField,
	"IntField":      TypeIntField,
	"FloatField":    TypeFloatField,
	"BoolField":     TypeBoolField,
	"BinaryField":   TypeBinaryField,
	"DateField":     TypeDateField,
	"DatetimeField": TypeDatetimeField,
}

type ColumnInfo struct {
	OrmType   OrmType
	ColName   string
	FieldName string
	Nullable  bool
	IsInc     bool
	IsPk      bool
	IsUni     bool
	DefVal    interface{}
	Offset    uintptr
	SqlType   string
	Charset   string
	Collate   string
}

func (ci *ColumnInfo) colName() string {
	return wrap(ci.ColName)
}

type OneToOneInfo struct {
	SrcCol    string
	DstDB     string
	DstTab    string
	DstCol    string
	Offset    uintptr
	FieldName string
}

type ForeignKeyInfo struct {
	SrcCol    string
	DstDB     string
	DstTab    string
	DstCol    string
	Offset    uintptr
	FieldName string
}

type ReverseForeignKeyInfo struct {
	SrcCol    string
	DstDB     string
	DstTab    string
	DstCol    string
	Offset    uintptr
	FieldName string
}

type ManyToManyInfo struct {
	SrcCol      string
	MidDB       string
	MidTab      string
	MidLeftCol  string
	MidRightCol string
	DstDB       string
	DstTab      string
	DstCol      string
	Offset      uintptr
	FieldName   string
}

type PrimaryKey []*ColumnInfo

func (Pk PrimaryKey) match(fields ...Field) bool {
	for _, PkCol := range Pk {
		get := false
		for _, field := range fields {
			if field.columnName() == PkCol.ColName {
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

func (Pk PrimaryKey) getFields(addr uintptr) []Field {
	fields := make([]Field, len(Pk))
	for i, PkCol := range Pk {
		fields[i] = getFieldByColumnInfo(addr, PkCol)
	}
	return fields
}

func (Pk PrimaryKey) genCreateClause() string {
	l := make([]string, len(Pk))
	for i, col := range Pk {
		l[i] = col.colName()
	}
	return fmt.Sprintf(", PRIMARY KEY (%s)", strings.Join(l, ", "))
}

func (Pk PrimaryKey) methodString() string {
	l := make([]string, len(Pk))
	for i, col := range Pk {
		l[i] = fmt.Sprintf("%q", col.ColName)
	}
	return fmt.Sprintf("return []string{ %s }", strings.Join(l, ", "))
}

type UniqueKeys [][]*ColumnInfo

func (uks UniqueKeys) match(fields ...Field) (int, bool) {
	for i, uk := range uks {
		fit := true
		for _, col := range uk {
			get := false
			for _, field := range fields {
				if col.ColName == field.columnName() {
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

func (uks UniqueKeys) getFields(addr uintptr, ukIndex int) []Field {
	uk := uks[ukIndex]
	fields := make([]Field, len(uk))
	for i, ukCol := range uk {
		fields[i] = getFieldByColumnInfo(addr, ukCol)
	}
	return fields
}

func (uks UniqueKeys) genCreateClause() string {
	keyList := make([]string, len(uks))
	for i, key := range uks {
		colList := make([]string, len(key))
		for j, col := range key {
			colList[j] = col.colName()
		}
		keyList[i] = fmt.Sprintf(", UNIQUE KEY (%s)", strings.Join(colList, ", "))
	}
	return strings.Join(keyList, "")
}

func (uks UniqueKeys) methodString() string {
	keyList := make([]string, len(uks))
	for i, key := range uks {
		colList := make([]string, len(key))
		for j, col := range key {
			colList[j] = fmt.Sprintf("%q", col.ColName)
		}
		keyList[i] = fmt.Sprintf("[]string{ %s },\n", strings.Join(colList, ", "))
	}
	return fmt.Sprintf(`return [][]string{ 
		%s 
	}`, strings.Join(keyList, ""))
}

type Keys [][]*ColumnInfo

func (ks Keys) genCreateClause() string {
	keyList := make([]string, len(ks))
	for i, key := range ks {
		colList := make([]string, len(key))
		for j, col := range key {
			colList[j] = col.colName()
		}
		keyList[i] = fmt.Sprintf(", KEY (%s)", strings.Join(colList, ", "))
	}
	return strings.Join(keyList, "")
}

func (ks Keys) methodString() string {
	keyList := make([]string, len(ks))
	for i, key := range ks {
		colList := make([]string, len(key))
		for j, col := range key {
			colList[j] = fmt.Sprintf("%q", col.ColName)
		}
		keyList[i] = fmt.Sprintf("[]string{ %s },\n", strings.Join(colList, ", "))
	}
	return fmt.Sprintf(`return [][]string{ 
		%s 
	}`, strings.Join(keyList, ""))
}

type TableInfo struct {
	DBName             string
	TabName            string
	ModelType          reflect.Type
	ModelName          string
	Columns            []*ColumnInfo
	ColumnMap          map[string]*ColumnInfo
	OneToOnes          []*OneToOneInfo
	ForeignKeys        []*ForeignKeyInfo
	ReverseForeignKeys []*ReverseForeignKeyInfo
	ManyToManys        []*ManyToManyInfo
	Inc                *ColumnInfo
	Pk                 PrimaryKey
	PkNames            []string
	Unis               UniqueKeys
	UniNames           [][]string
	Keys               Keys
	KeyNames           [][]string
	ModelStatus        uintptr
	IsComplete         bool
	IsNewMiddleTable   bool
}

func (ti *TableInfo) lookupColInfoByFieldName(FieldName string) *ColumnInfo {
	for _, colInfo := range ti.Columns {
		if colInfo.FieldName == FieldName {
			return colInfo
		}
	}
	return nil
}

func (ti *TableInfo) complement(tab table) {
	modelType := reflect.TypeOf(tab).Elem()
	if modelType.Kind() == reflect.Slice || modelType.Kind() == reflect.Array {
		modelType = modelType.Elem().Elem()
	}
	ti.ModelType = modelType
	for _, col := range ti.Columns {
		field, ok := modelType.FieldByName(col.FieldName)
		if !ok {
			panic(fmt.Errorf("nborm.TableInfo.complement() error: cannot find field (%s.%s)", modelType.String(), col.FieldName))
		}
		col.Offset = field.Offset
	}
	for _, col := range ti.ColumnMap {
		field, ok := modelType.FieldByName(col.FieldName)
		if !ok {
			panic(fmt.Errorf("nborm.TableInfo.complement() error: cannot find field (%s.%s)", modelType.String(), col.FieldName))
		}
		col.Offset = field.Offset
	}
	for _, oto := range ti.OneToOnes {
		field, ok := modelType.FieldByName(oto.FieldName)
		if !ok {
			panic(fmt.Errorf("nborm.TableInfo.complement() error: cannot find one to one field (%s.%s)", modelType.String(), oto.FieldName))
		}
		oto.Offset = field.Offset
	}
	for _, fk := range ti.ForeignKeys {
		field, ok := modelType.FieldByName(fk.FieldName)
		if !ok {
			panic(fmt.Errorf("nborm.TableInfo.complement() error: cannot find foreign key field (%s.%s)", modelType.String(), fk.FieldName))
		}
		fk.Offset = field.Offset
	}
	for _, rfk := range ti.ReverseForeignKeys {
		field, ok := modelType.FieldByName(rfk.FieldName)
		if !ok {
			panic(fmt.Errorf("nborm.TableInfo.complement() error: cannot find reverse foreign key field (%s.%s)", modelType.String(), rfk.FieldName))
		}
		rfk.Offset = field.Offset

	}
	for _, mtm := range ti.ManyToManys {
		field, ok := modelType.FieldByName(mtm.FieldName)
		if !ok {
			panic(fmt.Errorf("nborm.TableInfo.complement() error: cannot find many to many field (%s.%s)", modelType.String(), mtm.FieldName))
		}
		mtm.Offset = field.Offset
	}
	for _, col := range ti.Pk {
		field, ok := modelType.FieldByName(col.FieldName)
		if !ok {
			panic(fmt.Errorf("nborm.TableInfo.complement() error: cannot find primary key field (%s.%s)", modelType.String(), col.FieldName))
		}
		col.Offset = field.Offset
	}
	for _, uniqueKey := range ti.Unis {
		for _, col := range uniqueKey {
			field, ok := modelType.FieldByName(col.FieldName)
			if !ok {
				panic(fmt.Errorf("nborm.TableInfo.complement() error: cannot find unique key field (%s.%s)", modelType.String(), col.FieldName))
			}
			col.Offset = field.Offset
		}
	}
	for _, key := range ti.Keys {
		for _, col := range key {
			field, ok := modelType.FieldByName(col.FieldName)
			if !ok {
				panic(fmt.Errorf("nborm.TableInfo.complement() error: cannot find key field (%s.%s)", modelType.String(), col.FieldName))
			}
			col.Offset = field.Offset
		}
	}
	if ti.Inc != nil {
		field, ok := modelType.FieldByName(ti.Inc.FieldName)
		if !ok {
			panic(fmt.Errorf("nborm.TableInfo.complement() error: cannot find auto increment field (%s.%s)", modelType.String(), ti.Inc.FieldName))
		}
		ti.Inc.Offset = field.Offset
	}
	var gotModelStatus bool
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		if field.Type.String() == "nborm.ModelStatus" {
			ti.ModelStatus = field.Offset
			gotModelStatus = true
			break
		}
	}
	if !gotModelStatus {
		panic(fmt.Errorf("nborm.TabInfo.complement() error: ModelStatus field not exists (%s)", modelType.String()))
	}
	ti.IsComplete = true
}

func (ti *TableInfo) dbName() string {
	return wrap(ti.DBName)
}

func (ti *TableInfo) tabName() string {
	return wrap(ti.TabName)
}

type DatabaseInfo struct {
	Tables   []*TableInfo
	TableMap map[string]*TableInfo
}

func (si *DatabaseInfo) getOrCreate(tabName string) *TableInfo {
	if info, ok := si.TableMap[tabName]; ok {
		return info
	} else {
		info = &TableInfo{
			Columns:            make([]*ColumnInfo, 0, 32),
			ColumnMap:          make(map[string]*ColumnInfo),
			OneToOnes:          make([]*OneToOneInfo, 0, 8),
			ForeignKeys:        make([]*ForeignKeyInfo, 0, 8),
			ReverseForeignKeys: make([]*ReverseForeignKeyInfo, 0, 8),
			ManyToManys:        make([]*ManyToManyInfo, 0, 8),
			Pk:                 make(PrimaryKey, 0, 8),
			Unis:               make(UniqueKeys, 0, 8),
			Keys:               make(Keys, 0, 8),
		}
		si.Tables = append(si.Tables, info)
		si.TableMap[tabName] = info
		return info
	}
}

type SchemaInfo struct {
	Databases   []*DatabaseInfo
	DatabaseMap map[string]*DatabaseInfo
	mux         sync.RWMutex
}

func (si *SchemaInfo) getOrCreate(dbName string) *DatabaseInfo {
	if db, ok := si.DatabaseMap[dbName]; ok {
		return db
	} else {
		dbInfo := &DatabaseInfo{make([]*TableInfo, 0, 16), make(map[string]*TableInfo)}
		si.Databases = append(si.Databases, dbInfo)
		si.DatabaseMap[dbName] = dbInfo
		return dbInfo
	}
}

func (si *SchemaInfo) getTabInfo(tab table) *TableInfo {
	si.mux.RLock()
	defer si.mux.RUnlock()
	dbName, tabName := tab.DB(), tab.Tab()
	dbInfo, ok := si.DatabaseMap[dbName]
	if !ok {
		panic(fmt.Errorf("nborm.SchemaInfo.getTabInfo() error: database not exist (%s)", dbName))
	}
	tabInfo, ok := dbInfo.TableMap[tabName]
	if !ok {
		panic(fmt.Errorf("nborm.SchemaInfo.getTabInfo() error: table not exists (%s.%s)", dbName, tabName))
	}
	return tabInfo
}

var SchemaCache = SchemaInfo{
	make([]*DatabaseInfo, 0, 16),
	make(map[string]*DatabaseInfo),
	sync.RWMutex{},
}

func CleanSchemaCache() {
	SchemaCache = SchemaInfo{
		make([]*DatabaseInfo, 0, 16),
		make(map[string]*DatabaseInfo),
		sync.RWMutex{},
	}
}

var sqlConstRe = regexp.MustCompile(`^@'(.+)'$`)
var sqlStringRe = regexp.MustCompile(`^s'(.*)'$`)
var sqlIntRe = regexp.MustCompile(`^d'(\d+)'$`)
var sqlFloatRe = regexp.MustCompile(`^f'(.*)'$`)
var sqlBoolRe = regexp.MustCompile(`^b'(true|false)'$`)
var sqlBinaryRe = regexp.MustCompile(`^x'.*'$`)

// func parseColumn(field reflect.StructField) *ColumnInfo {
// 	c := &ColumnInfo{}
// 	c.Offset = field.Offset
// 	var ok bool
// 	c.OrmType, ok = OrmTypeMap[field.Type.String()]
// 	if !ok {
// 		panic(fmt.Errorf("nborm.parseColumn() error: unsupported field type (%s)", field.Type.Name()))
// 	}
// 	if ColName, ok := field.Tag.Lookup("column"); !ok {
// 		c.ColName = wrap(toSnakeCase(field.Name))
// 	} else {
// 		c.ColName = wrap(ColName)
// 	}
// 	if Nullable, ok := field.Tag.Lookup("Nullable"); ok && Nullable == "true" {
// 		c.Nullable = true
// 	}
// 	if IsInc, ok := field.Tag.Lookup("auto_Increment"); ok && c.OrmType == TypeIntField && IsInc == "true" {
// 		c.IsInc = true
// 	}
// 	if IsPk, ok := field.Tag.Lookup("primary_key"); ok && IsPk == "true" {
// 		c.IsPk = true
// 	}
// 	if IsUni, ok := field.Tag.Lookup("unique"); ok && IsUni == "true" {
// 		c.IsUni = true
// 	}
// 	if DefVal, ok := field.Tag.Lookup("default_value"); ok {
// 		if group := sqlConstRe.FindStringSubmatch(DefVal); len(group) > 1 {
// 			c.DefVal = group[1]
// 		} else {
// 			var err error
// 			errInvalidDefaultValue := fmt.Errorf("nborm.parseColumn() error: invalid default value (%s) for field (%s)", DefVal, field.Name)
// 			switch c.OrmType {
// 			case TypeStringField, TypeDateField, TypeDatetimeField:
// 				if group := sqlStringRe.FindStringSubmatch(DefVal); len(group) > 1 {
// 					c.DefVal = fmt.Sprintf("%q", group[1])
// 				} else {
// 					panic(errInvalidDefaultValue)
// 				}
// 			case TypeIntField:
// 				if group := sqlIntRe.FindStringSubmatch(DefVal); len(group) > 1 {
// 					if c.DefVal, err = strconv.ParseInt(group[1], 10, 64); err != nil {
// 						panic(err)
// 					}
// 				} else {
// 					panic(errInvalidDefaultValue)
// 				}
// 			case TypeFloatField:
// 				if group := sqlFloatRe.FindStringSubmatch(DefVal); len(group) > 1 {
// 					if c.DefVal, err = strconv.ParseFloat(group[1], 64); err != nil {
// 						panic(err)
// 					}
// 				} else {
// 					panic(errInvalidDefaultValue)
// 				}
// 			case TypeBoolField:
// 				if group := sqlBoolRe.FindStringSubmatch(DefVal); len(group) > 1 {
// 					if c.DefVal, err = strconv.ParseBool(group[1]); err != nil {
// 						panic(err)
// 					}
// 				} else {
// 					panic(errInvalidDefaultValue)
// 				}
// 			case TypeBinaryField:
// 				if sqlBinaryRe.MatchString(DefVal) {
// 					c.DefVal = DefVal
// 				} else {
// 					panic(errInvalidDefaultValue)
// 				}
// 			}
// 		}
// 	}
// 	return c
// }

// func getSrcCol(field reflect.StructField) string {
// 	if SrcCol, ok := field.Tag.Lookup("src_col"); !ok {
// 		panic(fmt.Errorf("nborm.getSrcCol() error: no source column tag for field (%s)", field.Name))
// 	} else {
// 		return wrap(SrcCol)
// 	}
// }

// func getDstCol(field reflect.StructField) (DstDB, DstTab, DstCol string) {
// 	var ok bool
// 	if DstDB, ok = field.Tag.Lookup("dst_db"); !ok {
// 		panic(fmt.Errorf("nborm.getDstCol() error: no destination database tag for field (%s)", field.Name))
// 	}
// 	if DstTab, ok = field.Tag.Lookup("dst_tab"); !ok {
// 		panic(fmt.Errorf("nborm.getDstCol() error: no destination table tag for field (%s)", field.Name))
// 	}
// 	if DstCol, ok = field.Tag.Lookup("dst_col"); !ok {
// 		panic(fmt.Errorf("nborm.getDstCol() error: no destination column tag for field (%s)", field.Name))
// 	}
// 	DstDB, DstTab, DstCol = wrap(DstDB), wrap(DstTab), wrap(DstCol)
// 	return
// }

// func getMidCol(field reflect.StructField) (MidDB, MidTab, MidLeftCol, MidRightCol string) {
// 	var ok bool
// 	if MidDB, ok = field.Tag.Lookup("mid_db"); !ok {
// 		panic(fmt.Errorf("nborm.getMidCol() error: no middle database tag for field (%s)", field.Name))
// 	}
// 	if MidTab, ok = field.Tag.Lookup("mid_tab"); !ok {
// 		panic(fmt.Errorf("nborm.getMidCol() error: no middle table tag for field (%s)", field.Name))
// 	}
// 	if MidLeftCol, ok = field.Tag.Lookup("mid_left_col"); !ok {
// 		panic(fmt.Errorf("nborm.getMidCol() error: no middle left column tag for field (%s)", field.Name))
// 	}
// 	if MidRightCol, ok = field.Tag.Lookup("mid_right_col"); !ok {
// 		panic(fmt.Errorf("nborm.getMidCol() error: no middle right column tag for field (%s)", field.Name))
// 	}
// 	MidDB, MidTab, MidLeftCol, MidRightCol = wrap(MidDB), wrap(MidTab), wrap(MidLeftCol), wrap(MidRightCol)
// 	return
// }

// func parseOneToOne(field reflect.StructField) *OneToOneInfo {
// 	o := &OneToOneInfo{}
// 	o.SrcCol = getSrcCol(field)
// 	o.DstDB, o.DstTab, o.DstCol = getDstCol(field)
// 	o.Offset = field.Offset
// 	return o
// }

// func parseForeignKey(field reflect.StructField) *ForeignKeyInfo {
// 	f := &ForeignKeyInfo{}
// 	f.SrcCol = getSrcCol(field)
// 	f.DstDB, f.DstTab, f.DstCol = getDstCol(field)
// 	f.Offset = field.Offset
// 	return f
// }

// func parseReverseForeignKey(field reflect.StructField) *ReverseForeignKeyInfo {
// 	r := &ReverseForeignKeyInfo{}
// 	r.SrcCol = getSrcCol(field)
// 	r.DstDB, r.DstTab, r.DstCol = getDstCol(field)
// 	r.Offset = field.Offset
// 	return r
// }

// func parseManyToMany(field reflect.StructField) *ManyToManyInfo {
// 	m := &ManyToManyInfo{}
// 	m.SrcCol = getSrcCol(field)
// 	m.MidDB, m.MidTab, m.MidLeftCol, m.MidRightCol = getMidCol(field)
// 	m.DstDB, m.DstTab, m.DstCol = getDstCol(field)
// 	m.Offset = field.Offset
// 	return m
// }

// func parseTable(table table) *TableInfo {
// 	dbName, tabName := wrap(table.DB()), wrap(table.Tab())
// 	tabInfo := &TableInfo{DB: dbName, Tab: tabName, ColumnMap: make(map[string]*ColumnInfo)}
// 	var haveModelStatus bool
// 	typ := reflect.TypeOf(table)
// 	if typ.Kind() != reflect.Ptr {
// 		panic(fmt.Errorf("nborm.parseTable() error: require a pointer but supported (%s)", typ.Name()))
// 	}
// 	elem := typ.Elem()
// 	kind := elem.Kind()
// 	var stct reflect.Type
// 	switch kind {
// 	case reflect.Struct:
// 		stct = elem
// 	case reflect.Array, reflect.Slice:
// 		sElem := elem.Elem()
// 		if sElem.Kind() != reflect.Ptr {
// 			panic(fmt.Errorf("nborm.parseTable() error: require a array of struct pointer (%s)", typ.Name()))
// 		}
// 		stct = sElem.Elem()
// 		if stct.Kind() != reflect.Struct {
// 			panic(fmt.Errorf("nborm.parseTable() error: the type of element of slice(array) is not a struct pointer (%s)", typ.Name()))
// 		}
// 	default:
// 		panic(fmt.Errorf("nborm.parseTable() error: object must be a struct or array of struct pointer (%s)", typ.Name()))
// 	}
// 	tabInfo.ModelType = stct
// 	for i := 0; i < stct.NumField(); i++ {
// 		field := stct.Field(i)
// 		switch field.Type.String() {
// 		case "nborm.StringField", "nborm.IntField", "nborm.FloatField", "nborm.BoolField", "nborm.BinaryField", "nborm.DateField",
// 			"nborm.DatetimeField":
// 			colInfo := parseColumn(field)
// 			tabInfo.ColumnMap[colInfo.ColName] = colInfo
// 			tabInfo.Columns = append(tabInfo.Columns, colInfo)
// 			if colInfo.IsInc {
// 				if tabInfo.Inc != nil {
// 					panic(fmt.Errorf("nborm.parseTable() error: multiple auto Increment Columns (%s)", typ.Name()))
// 				}
// 				tabInfo.Inc = colInfo
// 			}
// 			if colInfo.IsPk {
// 				tabInfo.Pk = append(tabInfo.Pk, colInfo)
// 			}
// 		case "nborm.OneToOne":
// 			otoInfo := parseOneToOne(field)
// 			tabInfo.OneToOnes = append(tabInfo.OneToOnes, otoInfo)
// 		case "nborm.ForeignKey":
// 			fkInfo := parseForeignKey(field)
// 			tabInfo.ForeignKeys = append(tabInfo.ForeignKeys, fkInfo)
// 		case "nborm.ReverseForeignKey":
// 			rfkInfo := parseReverseForeignKey(field)
// 			tabInfo.ReverseForeignKeys = append(tabInfo.ReverseForeignKeys, rfkInfo)
// 		case "nborm.ManyToMany":
// 			mtmInfo := parseManyToMany(field)
// 			tabInfo.ManyToManys = append(tabInfo.ManyToManys, mtmInfo)
// 		case "nborm.ModelStatus":
// 			if haveModelStatus {
// 				panic(fmt.Errorf("nborm.parseTable() error: multiple ModelStatus field (%s)", typ.Name()))
// 			}
// 			tabInfo.ModelStatus = field.Offset
// 			haveModelStatus = true
// 		}
// 	}
// 	if len(tabInfo.Pk) == 0 {
// 		panic(fmt.Errorf("nborm.parseTable() error: no primary key in %s.%s", dbName, tabName))
// 	}
// 	if !haveModelStatus {
// 		panic(fmt.Errorf("nborm.parseTable() error: no ModelStatus field in %s.%s", dbName, tabName))
// 	}
// 	return tabInfo
// }

// func getTabInfo(table table) *TableInfo {
// 	db, tab := wrap(table.DB()), wrap(table.Tab())
// 	schemaLock.RLock()
// 	dbInfo, ok := SchemaCache.DatabaseMap[db]
// 	if !ok {
// 		schemaLock.RUnlock()
// 		tInfo := parseTable(table)
// 		dInfo := &DatabaseInfo{Tables: []*TableInfo{tInfo}, TableMap: map[string]*TableInfo{tab: tInfo}}
// 		schemaLock.Lock()
// 		SchemaCache.Databases = append(SchemaCache.Databases, dInfo)
// 		SchemaCache.DatabaseMap[db] = dInfo
// 		schemaLock.Unlock()
// 		return tInfo
// 	}
// 	tabInfo, ok := dbInfo.TableMap[tab]
// 	if !ok {
// 		schemaLock.RUnlock()
// 		tInfo := parseTable(table)
// 		schemaLock.Lock()
// 		dbInfo.Tables = append(dbInfo.Tables, tInfo)
// 		dbInfo.TableMap[tab] = tInfo
// 		schemaLock.Unlock()
// 		return tInfo
// 	}
// 	schemaLock.RUnlock()
// 	return tabInfo
// }

// func getTabInfoByName(db, tab string) *TableInfo {
// 	SchemaCache.mux.RLock()
// 	if dbInfo, ok := SchemaCache.DatabaseMap[db]; !ok {
// 		SchemaCache.mux.RUnlock()
// 		panic(fmt.Errorf("nborm.getTabInfoByName() error: database not exists (%s)", db))
// 	} else {
// 		if tabInfo, ok := dbInfo.TableMap[tab]; !ok {
// 			SchemaCache.mux.RUnlock()
// 			panic(fmt.Errorf("nborm.getTabInfoByName() error: table not exists (%s.%s)", db, tab))
// 		} else {
// 			if !tabInfo.IsComplete {
// 				SchemaCache.mux.RUnlock()
// 				SchemaCache.mux.Lock()
// 				info.
// 			}
// 		}
// 	}
// }

func getTabInfo(tab table) *TableInfo {
	info := SchemaCache.getTabInfo(tab)
	if !info.IsComplete {
		SchemaCache.mux.Lock()
		info.complement(tab)
		SchemaCache.mux.Unlock()
	}
	return info
}

func initModelWithTableInfo(model table, tabInfo *TableInfo) {
	db, tab := model.DB(), model.Tab()
	baseAddr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&model)) + uintptr(8)))
	for _, col := range tabInfo.Columns {
		switch col.OrmType {
		case TypeStringField:
			field := (*StringField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset = db, tab, col.ColName,
				col.Nullable, col.IsPk, col.IsUni, col.DefVal, col.Offset
		case TypeIntField:
			field := (*IntField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.inc, field.pk, field.uni, field.defVal, field.offset = db, tab, col.ColName,
				col.Nullable, col.IsInc, col.IsPk, col.IsUni, col.DefVal, col.Offset
		case TypeFloatField:
			field := (*FloatField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset = db, tab, col.ColName,
				col.Nullable, col.IsPk, col.IsUni, col.DefVal, col.Offset
		case TypeBoolField:
			field := (*BoolField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset = db, tab, col.ColName,
				col.Nullable, col.IsPk, col.IsUni, col.DefVal, col.Offset
		case TypeBinaryField:
			field := (*BinaryField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset = db, tab, col.ColName,
				col.Nullable, col.IsPk, col.IsUni, col.DefVal, col.Offset
		case TypeDateField:
			field := (*DateField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset = db, tab, col.ColName,
				col.Nullable, col.IsPk, col.IsUni, col.DefVal, col.Offset
		case TypeDatetimeField:
			field := (*DatetimeField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset = db, tab, col.ColName,
				col.Nullable, col.IsPk, col.IsUni, col.DefVal, col.Offset

		}
	}
	for _, oto := range tabInfo.OneToOnes {
		srcField := getFieldByName(baseAddr, oto.SrcCol, tabInfo)
		relField := (*OneToOne)(unsafe.Pointer(baseAddr + oto.Offset))
		relField.srcDB, relField.srcTab, relField.srcCol, relField.dstDB, relField.dstTab, relField.dstCol, relField.srcValF = db, tab,
			oto.SrcCol, oto.DstDB, oto.DstTab, oto.DstCol, srcField.value
	}
	for _, fk := range tabInfo.ForeignKeys {
		srcField := getFieldByName(baseAddr, fk.SrcCol, tabInfo)
		relField := (*ForeignKey)(unsafe.Pointer(baseAddr + fk.Offset))
		relField.srcDB, relField.srcTab, relField.srcCol, relField.dstDB, relField.dstTab, relField.dstCol, relField.srcValF = db, tab,
			fk.SrcCol, fk.DstDB, fk.DstTab, fk.DstCol, srcField.value
	}
	for _, rfk := range tabInfo.ReverseForeignKeys {
		srcField := getFieldByName(baseAddr, rfk.SrcCol, tabInfo)
		relField := (*ReverseForeignKey)(unsafe.Pointer(baseAddr + rfk.Offset))
		relField.srcDB, relField.srcTab, relField.srcCol, relField.dstDB, relField.dstTab, relField.dstCol, relField.srcValF = db, tab,
			rfk.SrcCol, rfk.DstDB, rfk.DstTab, rfk.DstCol, srcField.value
	}
	for _, mtm := range tabInfo.ManyToManys {
		srcField := getFieldByName(baseAddr, mtm.SrcCol, tabInfo)
		relField := (*ManyToMany)(unsafe.Pointer(baseAddr + mtm.Offset))
		relField.srcDB, relField.srcTab, relField.srcCol, relField.midDB, relField.midTab, relField.midLeftCol, relField.midRightCol,
			relField.dstDB, relField.dstTab, relField.dstCol, relField.srcValF = db, tab, mtm.SrcCol, mtm.MidDB, mtm.MidTab, mtm.MidLeftCol,
			mtm.MidRightCol, mtm.DstDB, mtm.DstTab, mtm.DstCol, srcField.value
	}
	setInit(baseAddr, tabInfo)
}

func newModelAddr(tabInfo *TableInfo) uintptr {
	model := reflect.New(tabInfo.ModelType)
	initModelWithTableInfo(model.Interface().(table), tabInfo)
	return model.Pointer()
}

func newModel(tabInfo *TableInfo) table {
	model := reflect.New(tabInfo.ModelType).Interface().(table)
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

func getFieldByName(addr uintptr, ColName string, tabInfo *TableInfo) Field {
	colInfo, ok := tabInfo.ColumnMap[ColName]
	if !ok {
		panic(fmt.Errorf("nborm.getFieldByName() error: %s.%s.%s column not exist", tabInfo.DBName, tabInfo.TabName, ColName))
	}
	switch colInfo.OrmType {
	case TypeStringField:
		return (*StringField)(unsafe.Pointer(addr + colInfo.Offset))
	case TypeIntField:
		return (*IntField)(unsafe.Pointer(addr + colInfo.Offset))
	case TypeFloatField:
		return (*FloatField)(unsafe.Pointer(addr + colInfo.Offset))
	case TypeBoolField:
		return (*BoolField)(unsafe.Pointer(addr + colInfo.Offset))
	case TypeBinaryField:
		return (*BinaryField)(unsafe.Pointer(addr + colInfo.Offset))
	case TypeDateField:
		return (*DateField)(unsafe.Pointer(addr + colInfo.Offset))
	case TypeDatetimeField:
		return (*DatetimeField)(unsafe.Pointer(addr + colInfo.Offset))
	default:
		panic(fmt.Errorf("nborm.getFieldByName() error: unknown field type (%d)", colInfo.OrmType))
	}
}

func getFieldByColumnInfo(addr uintptr, colInfo *ColumnInfo) Field {
	fieldAddr := unsafe.Pointer(addr + colInfo.Offset)
	switch colInfo.OrmType {
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
		panic(fmt.Errorf("nborm.getFieldByColumnInfo() error: unknown field type (%d)", colInfo.OrmType))
	}

}

func getPrimaryKeyFieldsWithTableInfo(addr uintptr, tabInfo *TableInfo) []Field {
	validFields := getValidFieldsWithTableInfo(addr, tabInfo)
	if tabInfo.Pk.match(validFields...) {
		return tabInfo.Pk.getFields(addr)
	}
	return nil
}

func getIncWithTableInfo(addr uintptr, info *TableInfo) Field {
	if info.Inc == nil {
		return nil
	}
	return (*IntField)(unsafe.Pointer(addr + info.Inc.Offset))
}

func getIncAndOthers(addr uintptr, tabInfo *TableInfo) (Inc Field, others []Field) {
	for _, col := range tabInfo.Columns {
		if col.IsInc {
			Inc = getFieldByColumnInfo(addr, col)
		} else {
			others = append(others, getFieldByColumnInfo(addr, col))
		}
	}
	return
}

func getUniqueFieldsWithTableInfo(addr uintptr, tabInfo *TableInfo) []Field {
	if len(tabInfo.Unis) == 0 {
		return nil
	}
	validFields := getValidFieldsWithTableInfo(addr, tabInfo)
	if idx, match := tabInfo.Unis.match(validFields...); match {
		return tabInfo.Unis.getFields(addr, idx)
	}
	return nil
}

// func getUinsWithTableInfo(addr uintptr, tabInfo *TableInfo) []Field {
// 	if tabInfo.Unis == nil {
// 		return nil
// 	}

// 	l := make([][]Field, len(tabInfo.Unis))
// 	for i, key := range tabInfo.Unis {
// 		subL := make([]Field, len(key))
// 		for j, col := range key {
// 			subL[j] = getFieldByColumnInfo(addr, col)
// 		}
// 		// l[i] = getFieldByColumnInfo(addr, uniCol)
// 		l[i] = subL
// 	}
// 	return l
// }

// func getAllFieldsWithTableInfo(addr uintptr, tabInfo *TableInfo) []interface{} {
// 	l := make([]interface{}, len(tabInfo.Columns))
// 	for i, colInfo := range tabInfo.Columns {
// 		l[i] = getFieldByColumnInfo(addr, colInfo)
// 	}
// 	return l
// }

func getAllFieldsWithTableInfo(addr uintptr, tabInfo *TableInfo) []Field {
	l := make([]Field, len(tabInfo.Columns))
	for i, colInfo := range tabInfo.Columns {
		l[i] = getFieldByColumnInfo(addr, colInfo)
	}
	return l
}

func getValidFieldsWithTableInfo(addr uintptr, tabInfo *TableInfo) []Field {
	l := make([]Field, 0, len(tabInfo.Columns))
	allFields := getAllFieldsWithTableInfo(addr, tabInfo)
	for _, field := range allFields {
		if field.IsValid() {
			l = append(l, field)
		}
	}
	return l
}

func scanRow(addr uintptr, tabInfo *TableInfo, row *sql.Row) error {
	fields := getAllFieldsWithTableInfo(addr, tabInfo)
	addrList := make([]interface{}, len(fields))
	for i, field := range fields {
		addrList[i] = field.(interface{})
	}
	if err := row.Scan(addrList...); err != nil {
		return err
	}
	setSync(addr, tabInfo)
	return nil
}

func unionScanRow(addrs []uintptr, tabInfos []*TableInfo, row *sql.Row) error {
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

func scanRows(addr uintptr, tabInfo *TableInfo, rows *sql.Rows) error {
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

func unionScanRows(addrs []uintptr, tabInfos []*TableInfo, rows *sql.Rows) error {
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
