package nborm

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
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
	DBName    string
	TabName   string
	ColName   string
	FieldName string
	Nullable  bool
	IsInc     bool
	IsPk      bool
	IsUni     bool
	DefVal    *DefaultValue
	Offset    uintptr
	SqlType   string
	Charset   string
	Collate   string
}

func (ci *ColumnInfo) dbName() string {
	return wrap(ci.DBName)
}

func (ci *ColumnInfo) tabName() string {
	return wrap(ci.TabName)
}

func (ci *ColumnInfo) colName() string {
	return wrap(ci.ColName)
}

func (ci *ColumnInfo) fullTabName() string {
	return fmt.Sprintf("%s.%s", wrap(ci.DBName), wrap(ci.TabName))
}

func (ci *ColumnInfo) fullColName() string {
	return fmt.Sprintf("%s.%s.%s", wrap(ci.DBName), wrap(ci.TabName), wrap(ci.ColName))
}

type OneToOneInfo struct {
	SrcCol      *ColumnInfo
	MidLeftCol  *ColumnInfo
	MidRightCol *ColumnInfo
	DstCol      *ColumnInfo
	Offset      uintptr
	FieldName   string
}

type ForeignKeyInfo struct {
	SrcCol    *ColumnInfo
	DstCol    *ColumnInfo
	Offset    uintptr
	FieldName string
}

type ReverseForeignKeyInfo struct {
	SrcCol    *ColumnInfo
	DstCol    *ColumnInfo
	Offset    uintptr
	FieldName string
}

type ManyToManyInfo struct {
	SrcCol      *ColumnInfo
	MidLeftCol  *ColumnInfo
	MidRightCol *ColumnInfo
	DstCol      *ColumnInfo
	Offset      uintptr
	FieldName   string
}

type PrimaryKey []*ColumnInfo

func (Pk PrimaryKey) match(fields ...Field) bool {
	for _, PkCol := range Pk {
		get := false
		for _, field := range fields {
			if escap(field.columnName()) == PkCol.ColName {
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
				if col.ColName == escap(field.columnName()) {
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
	FieldMap           map[string]*ColumnInfo
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
	Charset            string
	Collate            string
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
		if ti.FieldMap == nil {
			ti.FieldMap = make(map[string]*ColumnInfo)
		}
		ti.FieldMap[col.FieldName] = col
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

func (ti *TableInfo) rawDBName() string {
	return ti.DBName
}

func (ti *TableInfo) tabName() string {
	return wrap(ti.TabName)
}

func (ti *TableInfo) rawTabName() string {
	return ti.TabName
}

func (ti *TableInfo) fullTabName() string {
	return fmt.Sprintf("%s.%s", ti.dbName(), ti.tabName())
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

func (si *SchemaInfo) getTableInfoByTableName(dbName, tabName string) *TableInfo {
	si.mux.RLock()
	defer si.mux.RUnlock()
	dbInfo, ok := si.DatabaseMap[dbName]
	if !ok {
		panic(fmt.Errorf("nborm.SchemaInfo.getTableInfoByTableName() error: database not exist (%s)", dbName))
	}
	tabInfo, ok := dbInfo.TableMap[tabName]
	if !ok {
		panic(fmt.Errorf("nborm.SchemaInfo.getTableInfoByTableName() error: table not exists (%s.%s)", dbName, tabName))
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
var sqlFloatRe = regexp.MustCompile(`^f'(\d+\.?\d*)'$`)
var sqlBoolRe = regexp.MustCompile(`^b'(true|false)'$`)
var sqlBinaryRe = regexp.MustCompile(`^x'([a-fA-F0-9])'$`)
var sqlDateRe = regexp.MustCompile(`^dt'(\d{4}-\d{2}-\d{2})'$`)
var sqlDatetimeRe = regexp.MustCompile(`^tm'(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`)

type DefaultValueType int

const (
	ConstDefault DefaultValueType = iota
	StringDefault
	IntDefault
	FloatDefault
	BoolDefault
	BinaryDefault
	DateDefault
	DatetimeDefault
)

type DefaultValue struct {
	Type  DefaultValueType
	Value interface{}
}

func (dv *DefaultValue) createTableClause() string {
	if dv == nil {
		return "DEFAULT NULL"
	}
	switch dv.Type {
	case ConstDefault:
		return fmt.Sprintf("DEFAULT %s", dv.Value)
	case StringDefault:
		return fmt.Sprintf("DEFAULT %q", dv.Value)
	case IntDefault:
		return fmt.Sprintf("DEFAULT %d", dv.Value)
	case FloatDefault:
		return fmt.Sprintf("DEFAULT %f", dv.Value)
	case BoolDefault:
		return fmt.Sprintf("DEFAULT %t", dv.Value)
	case BinaryDefault:
		return fmt.Sprintf("DEFAULT x'%x'", dv.Value)
	case DateDefault:
		return fmt.Sprintf("DEFAULT %q", dv.Value.(time.Time).Format("2006-01-02"))
	case DatetimeDefault:
		return fmt.Sprintf("DEFAULT %q", dv.Value.(time.Time).Format("2006-01-02 15:04:05"))
	default:
		panic(fmt.Errorf("nborm.DefaultValue.createTableClause() error: unknown DefaultValue type (%v)", dv))
	}
}

func (dv *DefaultValue) value() interface{} {
	if dv == nil {
		return nil
	}
	// switch dv.Type {
	// case ConstDefault:
	// 	if strings.HasPrefix(dv.Value.(string), "CURRENT_TIMESTAMP") {
	// 		return time.Now()
	// 	} else {
	// 		return nil
	// 	}
	// case StringDefault:
	// 	return dv.Value.(string)
	// case IntDefault:
	// 	return dv.Value.(int64)
	// case FloatDefault:
	// 	return dv.Value.(float64)
	// case BoolDefault:
	// 	return dv.Value.(bool)
	// case BinaryDefault:
	// 	return dv.Value.([]byte)
	// case DateDefault, DatetimeDefault:
	// 	return dv.Value.(time.Time)
	// default:
	// 	panic(fmt.Errorf("nborm.DefaultValue.value() unsupported default type %v", dv))
	// }
	return dv.Value
}

func (dv *DefaultValue) UnmarshalJSON(b []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	switch DefaultValueType(int(m["Type"].(float64))) {
	case ConstDefault:
		if strings.HasPrefix(m["Value"].(string), "CURRENT_TIMESTAMP") {
			dv.Value = time.Now()
		} else {
			dv.Value = nil
		}
	case StringDefault:
		dv.Value = m["Value"].(string)
	case IntDefault:
		dv.Value = int64(m["Value"].(float64))
	case FloatDefault:
		dv.Value = m["Value"].(float64)
	case BoolDefault:
		dv.Value = m["Value"].(bool)
	case BinaryDefault:
		dv.Value = m["Value"].([]byte)
	case DateDefault, DatetimeDefault:
		dv.Value = m["Value"].(time.Time)
	default:
		return fmt.Errorf("nborm.DefaultValue.UnmarshalJSON() error: unsupported DefaultValue type (%v)", m)
	}
	return nil
}

func parseDefaultValue(s string) *DefaultValue {
	switch {
	case sqlConstRe.MatchString(s):
		return &DefaultValue{ConstDefault, sqlConstRe.FindStringSubmatch(s)[1]}
	case sqlStringRe.MatchString(s):
		return &DefaultValue{StringDefault, sqlStringRe.FindStringSubmatch(s)[1]}
	case sqlIntRe.MatchString(s):
		val, _ := strconv.ParseInt(sqlIntRe.FindStringSubmatch(s)[1], 10, 64)
		return &DefaultValue{IntDefault, val}
	case sqlFloatRe.MatchString(s):
		val, _ := strconv.ParseFloat(sqlFloatRe.FindStringSubmatch(s)[1], 64)
		return &DefaultValue{FloatDefault, val}
	case sqlBoolRe.MatchString(s):
		val, _ := strconv.ParseBool(sqlBoolRe.FindStringSubmatch(s)[1])
		return &DefaultValue{BoolDefault, val}
	case sqlBinaryRe.MatchString(s):
		val, _ := hex.DecodeString(sqlBinaryRe.FindStringSubmatch(s)[1])
		return &DefaultValue{BinaryDefault, val}
	case sqlDateRe.MatchString(s):
		val, _ := time.Parse("2006-01-02", sqlDateRe.FindStringSubmatch(s)[1])
		return &DefaultValue{DateDefault, val}
	case sqlDatetimeRe.MatchString(s):
		val, _ := time.Parse("2006-01-02 15:04:05", sqlDatetimeRe.FindStringSubmatch(s)[1])
		return &DefaultValue{DatetimeDefault, val}
	default:
		panic(fmt.Errorf("nborm.parseDefaultValue() error: unknown default value type (%s)", s))
	}
}

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
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset, field.validators, field.fieldName,
				field.modelName = db, tab, col.ColName, col.Nullable, col.IsPk, col.IsUni, col.DefVal, col.Offset, []Validator{nullValidator},
				col.FieldName, tabInfo.ModelName
		case TypeIntField:
			field := (*IntField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.inc, field.pk, field.uni, field.defVal, field.offset, field.validators,
				field.fieldName, field.modelName = db, tab, col.ColName, col.Nullable, col.IsInc, col.IsPk, col.IsUni, col.DefVal, col.Offset,
				[]Validator{nullValidator}, col.FieldName, tabInfo.ModelName
		case TypeFloatField:
			field := (*FloatField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset, field.validators, field.fieldName,
				field.modelName = db, tab, col.ColName, col.Nullable, col.IsPk, col.IsUni, col.DefVal, col.Offset, []Validator{nullValidator},
				col.FieldName, tabInfo.ModelName
		case TypeBoolField:
			field := (*BoolField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset, field.validators, field.fieldName,
				field.modelName = db, tab, col.ColName, col.Nullable, col.IsPk, col.IsUni, col.DefVal, col.Offset, []Validator{nullValidator},
				col.FieldName, tabInfo.ModelName
		case TypeBinaryField:
			field := (*BinaryField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset, field.validators, field.fieldName,
				field.modelName = db, tab, col.ColName, col.Nullable, col.IsPk, col.IsUni, col.DefVal, col.Offset, []Validator{nullValidator},
				col.FieldName, tabInfo.ModelName
		case TypeDateField:
			field := (*DateField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset, field.validators, field.fieldName,
				field.modelName = db, tab, col.ColName, col.Nullable, col.IsPk, col.IsUni, col.DefVal, col.Offset, []Validator{nullValidator},
				col.FieldName, tabInfo.ModelName
		case TypeDatetimeField:
			field := (*DatetimeField)(unsafe.Pointer(baseAddr + col.Offset))
			field.db, field.tab, field.column, field.nullable, field.pk, field.uni, field.defVal, field.offset, field.validators, field.fieldName,
				field.modelName = db, tab, col.ColName, col.Nullable, col.IsPk, col.IsUni, col.DefVal, col.Offset, []Validator{nullValidator},
				col.FieldName, tabInfo.ModelName
		}

	}
	for _, oto := range tabInfo.OneToOnes {
		srcField := getFieldByName(baseAddr, oto.SrcCol.ColName, tabInfo)
		relField := (*OneToOne)(unsafe.Pointer(baseAddr + oto.Offset))
		relField.srcDB = oto.SrcCol.DBName
		relField.srcTab = oto.SrcCol.TabName
		relField.srcCol = oto.SrcCol.ColName
		relField.midDB = oto.MidLeftCol.DBName
		relField.midTab = oto.MidLeftCol.TabName
		relField.midLeftCol = oto.MidLeftCol.ColName
		relField.midRightCol = oto.MidRightCol.ColName
		relField.dstDB = oto.DstCol.DBName
		relField.dstTab = oto.DstCol.TabName
		relField.dstCol = oto.DstCol.ColName
		relField.srcValF = srcField.value
	}
	for _, fk := range tabInfo.ForeignKeys {
		srcField := getFieldByName(baseAddr, fk.SrcCol.ColName, tabInfo)
		relField := (*ForeignKey)(unsafe.Pointer(baseAddr + fk.Offset))
		relField.srcDB = fk.SrcCol.DBName
		relField.srcTab = fk.SrcCol.TabName
		relField.srcCol = fk.SrcCol.ColName
		relField.dstDB = fk.DstCol.DBName
		relField.dstTab = fk.DstCol.TabName
		relField.dstCol = fk.DstCol.ColName
		relField.srcValF = srcField.value
	}
	for _, rfk := range tabInfo.ReverseForeignKeys {
		srcField := getFieldByName(baseAddr, rfk.SrcCol.ColName, tabInfo)
		relField := (*ReverseForeignKey)(unsafe.Pointer(baseAddr + rfk.Offset))
		relField.srcDB = rfk.SrcCol.DBName
		relField.srcTab = rfk.SrcCol.TabName
		relField.srcCol = rfk.SrcCol.ColName
		relField.dstDB = rfk.DstCol.DBName
		relField.dstTab = rfk.DstCol.TabName
		relField.dstCol = rfk.DstCol.ColName
		relField.srcValF = srcField.value
	}
	for _, mtm := range tabInfo.ManyToManys {
		srcField := getFieldByName(baseAddr, mtm.SrcCol.ColName, tabInfo)
		relField := (*ManyToMany)(unsafe.Pointer(baseAddr + mtm.Offset))
		relField.srcDB = mtm.SrcCol.DBName
		relField.srcTab = mtm.SrcCol.TabName
		relField.srcCol = mtm.SrcCol.ColName
		relField.midDB = mtm.MidLeftCol.DBName
		relField.midTab = mtm.MidLeftCol.TabName
		relField.midLeftCol = mtm.MidLeftCol.ColName
		relField.midRightCol = mtm.MidRightCol.ColName
		relField.dstDB = mtm.DstCol.DBName
		relField.dstTab = mtm.DstCol.TabName
		relField.dstCol = mtm.DstCol.ColName
		relField.srcValF = srcField.value
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
	ColName = escap(ColName)
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

func getFieldByFieldName(addr uintptr, fieldName string, tabInfo *TableInfo) (Field, error) {
	colInfo, ok := tabInfo.FieldMap[fieldName]
	if !ok {
		return nil, fmt.Errorf("nborm.getFieldByFieldName() error: field not exists (%s)", fieldName)
	}
	fieldAddr := unsafe.Pointer(addr + colInfo.Offset)
	switch colInfo.OrmType {
	case TypeStringField:
		return (*StringField)(fieldAddr), nil
	case TypeIntField:
		return (*IntField)(fieldAddr), nil
	case TypeFloatField:
		return (*FloatField)(fieldAddr), nil
	case TypeBoolField:
		return (*BoolField)(fieldAddr), nil
	case TypeBinaryField:
		return (*BinaryField)(fieldAddr), nil
	case TypeDateField:
		return (*DateField)(fieldAddr), nil
	case TypeDatetimeField:
		return (*DatetimeField)(fieldAddr), nil
	default:
		return nil, fmt.Errorf("nborm.getFieldByFieldName() error: unknown field type (%d)", colInfo.OrmType)
	}
}

func getFieldByOffset(addr, offset uintptr, ormType OrmType) (Field, error) {
	fieldAddr := unsafe.Pointer(addr + offset)
	switch ormType {
	case TypeStringField:
		return (*StringField)(fieldAddr), nil
	case TypeIntField:
		return (*IntField)(fieldAddr), nil
	case TypeFloatField:
		return (*FloatField)(fieldAddr), nil
	case TypeBoolField:
		return (*BoolField)(fieldAddr), nil
	case TypeBinaryField:
		return (*BinaryField)(fieldAddr), nil
	case TypeDateField:
		return (*DateField)(fieldAddr), nil
	case TypeDatetimeField:
		return (*DatetimeField)(fieldAddr), nil
	default:
		return nil, fmt.Errorf("nborm.getFieldByFieldName() error: unknown field type (%d)", ormType)
	}
}

func getPrimaryKeyFieldsWithTableInfo(addr uintptr, tabInfo *TableInfo) []Field {
	fields := getAllFieldsWithTableInfo(addr, tabInfo)
	validFields := make([]Field, 0, len(fields))
	for _, field := range fields {
		if field.IsValid() && !field.IsNull() {
			validFields = append(validFields, field)
		}
	}
	if len(validFields) == 0 {
		return nil
	}
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
	fields := getAllFieldsWithTableInfo(addr, tabInfo)
	validFields := make([]Field, 0, len(fields))
	for _, field := range fields {
		if field.IsValid() && !field.IsNull() {
			validFields = append(validFields, field)
		}
	}
	if len(validFields) == 0 {
		return nil
	}
	if idx, match := tabInfo.Unis.match(validFields...); match {
		return tabInfo.Unis.getFields(addr, idx)
	}
	return nil
}

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

func getFieldsWithoutInc(addr uintptr, tabInfo *TableInfo) []Field {
	l := make([]Field, 0, len(tabInfo.Columns))
	for _, col := range tabInfo.Columns {
		if col.ColName != tabInfo.Inc.ColName {
			l = append(l, getFieldByColumnInfo(addr, col))
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
			fieldAddrList := make([]interface{}, len(tabInfo.Columns))
			for j, field := range getAllFieldsWithTableInfo(modelAddr, tabInfo) {
				fieldAddrList[j] = field.(interface{})
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

func initModelForInsert(addr uintptr, tabInfo *TableInfo) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("nborm.initModelForInsert() error: %v", e)
		}
	}()
	fields := getFieldsWithoutInc(addr, tabInfo)
	for _, field := range fields {
		if !field.IsValid() {
			field.setVal(field.getDefVal())
		}
	}
	return nil
}

func processInsertModel(addr uintptr, tabInfo *TableInfo) ([]string, []string, []interface{}, error) {
	colLen := len(tabInfo.Columns)
	colList, placeHolderList, valList := make([]string, colLen), make([]string, colLen), make([]interface{}, colLen)
	for i, col := range tabInfo.Columns {
		field := getFieldByColumnInfo(addr, col)
		colList[i] = field.fullColName()
		if !field.IsValid() {
			if field.isInc() || field.isNullable() {
				field.setVal(nil)
				placeHolderList[i] = "?"
				valList[i] = nil
			} else {
				if defVal := field.getDefVal(); defVal == nil {
					return nil, nil, nil, fmt.Errorf("nborm.genInsertStmt() error: column cannot be null (%s)", field.fullColName())
				} else {
					if binDef, ok := defVal.([]byte); ok {
						field.setVal(binDef)
						placeHolderList[i] = fmt.Sprintf("x'%x'", binDef)
					} else {
						field.setVal(defVal)
						placeHolderList[i] = "?"
						valList[i] = defVal
					}
				}
			}
		} else {
			if field.IsNull() {
				if field.isNullable() {
					placeHolderList[i] = "?"
					valList[i] = nil
				} else {
					return nil, nil, nil, fmt.Errorf("nborm.genInsertStmt() error: column cannot be null (%s)", field.fullColName())
				}
			} else {
				if binField, ok := field.(*BinaryField); ok {
					placeHolderList[i] = binField.strVal().(string)
				} else {
					placeHolderList[i] = "?"
					valList[i] = field.value()
				}
			}
		}
	}
	return colList, placeHolderList, valList, nil
}
