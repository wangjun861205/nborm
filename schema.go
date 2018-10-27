package nborm

import "encoding/json"

//Column information_schema.columns
type Column struct {
	TABLE_CATALOG            *StringField
	TABLE_SCHEMA             *StringField
	TABLE_NAME               *StringField
	COLUMN_NAME              *StringField
	ORDINAL_POSITION         *IntField
	COLUMN_DEFAULT           *StringField
	IS_NULLABLE              *StringField
	DATA_TYPE                *StringField
	CHARACTER_MAXIMUM_LENGTH *IntField
	CHARACTER_OCTET_LENGTH   *IntField
	NUMERIC_PRECISION        *IntField
	NUMERIC_SCALE            *IntField
	DATETIME_PRECISION       *IntField
	CHARACTER_SET_NAME       *StringField
	COLLATION_NAME           *StringField
	COLUMN_TYPE              *StringField
	COLUMN_KEY               *StringField
	EXTRA                    *StringField
	PRIVILEGES               *StringField
	COLUMN_COMMENT           *StringField
	GENERATION_EXPRESSION    *StringField
	SRS_ID                   *IntField
	_isSync                  bool
}

//NewColumn create Column
func NewColumn() *Column {
	col := &Column{}
	col.TABLE_CATALOG = NewStringField(col, "TABLE_CATALOG", false, false, false)
	col.TABLE_SCHEMA = NewStringField(col, "TABLE_SCHEMA", false, false, false)
	col.TABLE_NAME = NewStringField(col, "TABLE_NAME", false, false, false)
	col.COLUMN_NAME = NewStringField(col, "COLUMN_NAME", false, false, false)
	col.ORDINAL_POSITION = NewIntField(col, "ORDINAL_POSITION", false, false, false)
	col.COLUMN_DEFAULT = NewStringField(col, "COLUMN_DEFAULT", false, false, false)
	col.IS_NULLABLE = NewStringField(col, "IS_NULLABLE", false, false, false)
	col.DATA_TYPE = NewStringField(col, "DATA_TYPE", false, false, false)
	col.CHARACTER_MAXIMUM_LENGTH = NewIntField(col, "CHARACTER_MAXIMUM_LENGTH", false, false, false)
	col.CHARACTER_OCTET_LENGTH = NewIntField(col, "CHARACTER_OCTET_LENGTH", false, false, false)
	col.NUMERIC_PRECISION = NewIntField(col, "NUMERIC_PRECISION", false, false, false)
	col.NUMERIC_SCALE = NewIntField(col, "NUMERIC_SCALE", false, false, false)
	col.DATETIME_PRECISION = NewIntField(col, "DATETIME_PRECISION", false, false, false)
	col.CHARACTER_SET_NAME = NewStringField(col, "CHARACTER_SET_NAME", false, false, false)
	col.COLLATION_NAME = NewStringField(col, "COLLATION_NAME", false, false, false)
	col.COLUMN_TYPE = NewStringField(col, "COLUMN_TYPE", false, false, false)
	col.COLUMN_KEY = NewStringField(col, "COLUMN_KEY", false, false, false)
	col.EXTRA = NewStringField(col, "EXTRA", false, false, false)
	col.PRIVILEGES = NewStringField(col, "PRIVILEGES", false, false, false)
	col.COLUMN_COMMENT = NewStringField(col, "COLUMN_COMMENT", false, false, false)
	col.GENERATION_EXPRESSION = NewStringField(col, "GENERATION_EXPRESSION", false, false, false)
	col.SRS_ID = NewIntField(col, "SRS_ID", false, false, false)
	return col
}

//DB database name
func (c *Column) DB() string {
	return "information_schema"
}

//Tab table name
func (c *Column) Tab() string {
	return "columns"
}

//Fields all data Field
func (c *Column) Fields() []Field {
	return []Field{
		c.TABLE_CATALOG,
		c.TABLE_SCHEMA,
		c.TABLE_NAME,
		c.COLUMN_NAME,
		c.ORDINAL_POSITION,
		c.COLUMN_DEFAULT,
		c.IS_NULLABLE,
		c.DATA_TYPE,
		c.CHARACTER_MAXIMUM_LENGTH,
		c.CHARACTER_OCTET_LENGTH,
		c.NUMERIC_PRECISION,
		c.NUMERIC_SCALE,
		c.DATETIME_PRECISION,
		c.CHARACTER_SET_NAME,
		c.COLLATION_NAME,
		c.COLUMN_TYPE,
		c.COLUMN_KEY,
		c.EXTRA,
		c.PRIVILEGES,
		c.COLUMN_COMMENT,
		c.GENERATION_EXPRESSION,
		c.SRS_ID,
	}
}

//SetSync set synchronized status
func (c *Column) SetSync(b bool) {
	c._isSync = b
}

//GetSync get synchronized status
func (c *Column) GetSync() bool {
	return c._isSync
}

//KeyColumnUsage information_schema.key_column_usage
type KeyColumnUsage struct {
	CONSTRAINT_CATALOG            *StringField
	CONSTRAINT_SCHEMA             *StringField
	CONSTRAINT_NAME               *StringField
	TABLE_CATALOG                 *StringField
	TABLE_SCHEMA                  *StringField
	TABLE_NAME                    *StringField
	COLUMN_NAME                   *StringField
	ORDINAL_POSITION              *IntField
	POSITION_IN_UNIQUE_CONSTRAINT *StringField
	REFERENCED_TABLE_SCHEMA       *StringField
	REFERENCED_TABLE_NAME         *StringField
	REFERENCED_COLUMN_NAME        *StringField
	_isSync                       bool
}

//NewKeyColumnUsage create KeyColumnUsage
func NewKeyColumnUsage() *KeyColumnUsage {
	keyColUsg := &KeyColumnUsage{}
	keyColUsg.CONSTRAINT_CATALOG = NewStringField(keyColUsg, "CONSTRAINT_CATALOG", false, false, false)
	keyColUsg.CONSTRAINT_SCHEMA = NewStringField(keyColUsg, "CONSTRAINT_SCHEMA", false, false, false)
	keyColUsg.CONSTRAINT_NAME = NewStringField(keyColUsg, "CONSTRAINT_NAME", false, false, false)
	keyColUsg.TABLE_CATALOG = NewStringField(keyColUsg, "TABLE_CATALOG", false, false, false)
	keyColUsg.TABLE_SCHEMA = NewStringField(keyColUsg, "TABLE_SCHEMA", false, false, false)
	keyColUsg.TABLE_NAME = NewStringField(keyColUsg, "TABLE_NAME", false, false, false)
	keyColUsg.COLUMN_NAME = NewStringField(keyColUsg, "COLUMN_NAME", false, false, false)
	keyColUsg.ORDINAL_POSITION = NewIntField(keyColUsg, "ORDINAL_POSITION", false, false, false)
	keyColUsg.POSITION_IN_UNIQUE_CONSTRAINT = NewStringField(keyColUsg, "POSITION_IN_UNIQUE_CONSTRAINT", false, false, false)
	keyColUsg.REFERENCED_TABLE_SCHEMA = NewStringField(keyColUsg, "REFERENCED_TABLE_SCHEMA", false, false, false)
	keyColUsg.REFERENCED_TABLE_NAME = NewStringField(keyColUsg, "REFERENCED_TABLE_NAME", false, false, false)
	keyColUsg.REFERENCED_COLUMN_NAME = NewStringField(keyColUsg, "REFERENCED_COLUMN_NAME", false, false, false)
	return keyColUsg
}

//DB database name
func (u *KeyColumnUsage) DB() string {
	return "information_schema"
}

//Tab table name
func (u *KeyColumnUsage) Tab() string {
	return "key_column_usage"
}

//Fields all data Fields
func (u *KeyColumnUsage) Fields() []Field {
	return []Field{
		u.CONSTRAINT_CATALOG,
		u.CONSTRAINT_SCHEMA,
		u.CONSTRAINT_NAME,
		u.TABLE_CATALOG,
		u.TABLE_SCHEMA,
		u.TABLE_NAME,
		u.COLUMN_NAME,
		u.ORDINAL_POSITION,
		u.POSITION_IN_UNIQUE_CONSTRAINT,
		u.REFERENCED_TABLE_SCHEMA,
		u.REFERENCED_TABLE_NAME,
		u.REFERENCED_COLUMN_NAME,
	}
}

//SetSync set synchronized status
func (u *KeyColumnUsage) SetSync(b bool) {
	u._isSync = b
}

//GetSync get synchronized status
func (u *KeyColumnUsage) GetSync() bool {
	return u._isSync
}

//Table information_schema.tables
type Table struct {
	TABLE_CATALOG   *StringField
	TABLE_SCHEMA    *StringField
	TABLE_NAME      *StringField
	TABLE_TYPE      *StringField
	ENGINE          *StringField
	VERSION         *IntField
	ROW_FORMAT      *StringField
	TABLE_ROWS      *IntField
	AVG_ROW_LENGTH  *IntField
	DATA_LENGTH     *IntField
	MAX_DATA_LENGTH *IntField
	INDEX_LENGTH    *IntField
	DATA_FREE       *IntField
	AUTO_INCREMENT  *IntField
	CREATE_TIME     *DatetimeField
	UPDATE_TIME     *DatetimeField
	CHECK_TIME      *DatetimeField
	TABLE_COLLATION *StringField
	CHECKSUM        *IntField
	CREATE_OPTIONS  *StringField
	TABLE_COMMENT   *StringField
	_isSync         bool
}

//NewTable create Table
func NewTable() *Table {
	tab := &Table{}
	tab.TABLE_CATALOG = NewStringField(tab, "TABLE_CATALOG", false, false, false)
	tab.TABLE_SCHEMA = NewStringField(tab, "TABLE_SCHEMA", false, false, false)
	tab.TABLE_NAME = NewStringField(tab, "TABLE_NAME", false, false, false)
	tab.TABLE_TYPE = NewStringField(tab, "TABLE_TYPE", false, false, false)
	tab.ENGINE = NewStringField(tab, "ENGINE", false, false, false)
	tab.VERSION = NewIntField(tab, "VERSION", false, false, false)
	tab.ROW_FORMAT = NewStringField(tab, "ROW_FORMAT", false, false, false)
	tab.TABLE_ROWS = NewIntField(tab, "TABLE_ROWS", false, false, false)
	tab.AVG_ROW_LENGTH = NewIntField(tab, "AVG_ROW_LENGTH", false, false, false)
	tab.DATA_LENGTH = NewIntField(tab, "DATA_LENGTH", false, false, false)
	tab.MAX_DATA_LENGTH = NewIntField(tab, "MAX_DATA_LENGTH", false, false, false)
	tab.INDEX_LENGTH = NewIntField(tab, "INDEX_LENGTH", false, false, false)
	tab.DATA_FREE = NewIntField(tab, "DATA_FREE", false, false, false)
	tab.AUTO_INCREMENT = NewIntField(tab, "AUTO_INCREMENT", false, false, false)
	tab.CREATE_TIME = NewDatetimeField(tab, "CREATE_TIME", false, false, false)
	tab.UPDATE_TIME = NewDatetimeField(tab, "UPDATE_TIME", false, false, false)
	tab.CHECK_TIME = NewDatetimeField(tab, "CHECK_TIME", false, false, false)
	tab.TABLE_COLLATION = NewStringField(tab, "TABLE_COLLATION", false, false, false)
	tab.CHECKSUM = NewIntField(tab, "CHECKSUM", false, false, false)
	tab.CREATE_OPTIONS = NewStringField(tab, "CREATE_OPTIONS", false, false, false)
	tab.TABLE_COMMENT = NewStringField(tab, "TABLE_COMMENT", false, false, false)
	return tab
}

//DB database name
func (t *Table) DB() string {
	return "information_schema"
}

//Tab table name
func (t *Table) Tab() string {
	return "tables"
}

//Fields all data Fields
func (t *Table) Fields() []Field {
	return []Field{
		t.TABLE_CATALOG,
		t.TABLE_SCHEMA,
		t.TABLE_NAME,
		t.TABLE_TYPE,
		t.ENGINE,
		t.VERSION,
		t.ROW_FORMAT,
		t.TABLE_ROWS,
		t.AVG_ROW_LENGTH,
		t.DATA_LENGTH,
		t.MAX_DATA_LENGTH,
		t.INDEX_LENGTH,
		t.DATA_FREE,
		t.AUTO_INCREMENT,
		t.CREATE_TIME,
		t.UPDATE_TIME,
		t.CHECK_TIME,
		t.TABLE_COLLATION,
		t.CHECKSUM,
		t.CREATE_OPTIONS,
		t.TABLE_COMMENT,
	}
}

//SetSync set synchronized status
func (t *Table) SetSync(b bool) {
	t._isSync = b
}

//GetSync get synchronized status
func (t *Table) GetSync() bool {
	return t._isSync
}

//ColumnList Column list
type ColumnList struct {
	M    *Column
	List []*Column
}

//NewColumnList create ColumnList
func NewColumnList() *ColumnList {
	return &ColumnList{
		NewColumn(),
		make([]*Column, 0, 128),
	}
}

//New create Column and append it to list
func (cl *ColumnList) New() Model {
	c := NewColumn()
	cl.List = append(cl.List, c)
	return c
}

//Len length of list
func (cl *ColumnList) Len() int {
	return len(cl.List)
}

//Index index operation
func (cl *ColumnList) Index(i int) Model {
	return (cl.List)[i]
}

//Delete delete operation
func (cl *ColumnList) Delete(i int) {
	switch i {
	case 0:
		cl.List = (cl.List)[1:]
	case cl.Len() - 1:
		cl.List = (cl.List)[:cl.Len()-1]
	default:
		cl.List = append((cl.List)[:i], (cl.List)[i+1:]...)
	}
}

//Swap swap element
func (cl *ColumnList) Swap(i, j int) {
	(cl.List)[i], (cl.List)[j] = (cl.List)[j], (cl.List)[i]
}

//MarshalJSON implement json.Marshaler interface
func (cl *ColumnList) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(cl.List, "\t", "\t")
}

//Model return example Model
func (cl *ColumnList) Model() Model {
	return cl.M
}

//KeyColumnUsageList KeyColumnUsage list
type KeyColumnUsageList struct {
	M    *KeyColumnUsage
	List []*KeyColumnUsage
}

//NewKeyColumnUsageList create KeyColumnUsageList
func NewKeyColumnUsageList() *KeyColumnUsageList {
	return &KeyColumnUsageList{
		NewKeyColumnUsage(),
		make([]*KeyColumnUsage, 0, 128),
	}
}

//New create a Model and append it to list
func (ul *KeyColumnUsageList) New() Model {
	u := NewKeyColumnUsage()
	ul.List = append(ul.List, u)
	return u
}

//Len length of list
func (ul *KeyColumnUsageList) Len() int {
	return len(ul.List)
}

//Index index operation
func (ul *KeyColumnUsageList) Index(i int) Model {
	return (ul.List)[i]
}

//Delete delete element
func (ul *KeyColumnUsageList) Delete(i int) {
	switch i {
	case 0:
		ul.List = (ul.List)[1:]
	case ul.Len() - 1:
		ul.List = (ul.List)[:ul.Len()-1]
	default:
		ul.List = append((ul.List)[:i], (ul.List)[i+1:]...)
	}
}

//Swap swap elements
func (ul *KeyColumnUsageList) Swap(i, j int) {
	(ul.List)[i], (ul.List)[j] = (ul.List)[j], (ul.List)[i]
}

//MarshalJSON implement json.Marshaler interface
func (ul *KeyColumnUsageList) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(ul.List, "\t", "\t")
}

//Model return example Model
func (ul *KeyColumnUsageList) Model() Model {
	return ul.M
}

//TableList Table list
type TableList struct {
	M    *Table
	List []*Table
}

//NewTableList create TableList
func NewTableList() *TableList {
	return &TableList{
		NewTable(),
		make([]*Table, 0, 128),
	}
}

//Len length of list
func (l *TableList) Len() int {
	return len(l.List)
}

//New create Model and append it to list
func (l *TableList) New() Model {
	m := NewTable()
	l.List = append(l.List, m)
	return m
}

//Index index operation
func (l *TableList) Index(i int) Model {
	return (l.List)[i]
}

//Delete delete element
func (l *TableList) Delete(i int) {
	switch i {
	case 0:
		l.List = (l.List)[1:]
	case l.Len() - 1:
		l.List = (l.List)[:l.Len()-1]
	default:
		l.List = append((l.List)[:i], (l.List)[i+1:]...)
	}
}

//Swap swap elements
func (l *TableList) Swap(i, j int) {
	(l.List)[i], (l.List)[j] = (l.List)[j], (l.List)[i]
}

//MarshalJSON implement json.Marshaler interface
func (l *TableList) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(l.List, "\t", "\t")
}

//Model return example Model
func (l *TableList) Model() Model {
	return l.M
}
