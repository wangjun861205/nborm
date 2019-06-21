package nborm

import "database/sql"

type RowScanner interface {
	ScanRow(*sql.Row) error
}

type RowsScanner interface {
	ScanRows(*sql.Rows) error
}

type Executor interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	Exec(string, ...interface{}) (sql.Result, error)
}

type BaseField interface {
	setModel(Model)
	dbName() string
	tabName() string
	colName() string
	setCol(string)
	fieldName() string
	setField(string)
	getStatus() fieldStatus
	setStatus(fieldStatus)
	addStatus(fieldStatus)
	removeStatus(fieldStatus)
	IsValid() bool
	setValid()
	unsetValid()
	IsNull() bool
	SetNull()
	unsetNull()
	isForWhere() bool
	setForWhere()
	unsetForWhere()
	isForUpdate() bool
	setForUpdate()
	unsetForUpdate()
	setPrimaryKey()
	unsetPrimaryKey()
	isPrimaryKey() bool
	setAutoInc()
	unsetAutoInc()
	isAutoInc() bool
	mustValid()
	rawFullTabName() string
	fullTabName() string
	fullColName() string
	ForSelect()
	ForSum()
	AscOrder()
	DscOrder()
	String() string
}

type Field interface {
	BaseField
	Scan(interface{}) error
	Value() interface{}
	Set(interface{})
	AndW() Field
	OrW() Field
	AndWhere(string, interface{}) Field
	OrWhere(string, interface{}) Field
	SetU()
	SetUpdate(interface{})
	whereList() whereList
	updateSet() *updateSet
	genAndWhere(string, interface{}) *where
	genOrWhere(string, interface{}) *where
}

type Model interface {
	DB() string
	Tab() string
	FieldInfos() FieldInfoList
	AutoIncField() Field
	PrimaryKey() FieldList
	Relations() RelationInfoList
	setRel(string, *where)
	getRelCols() string
	setRelCols(string)
	getRelJoin() string
	setRelJoin(string)
	getRelWhere() *where
	setRelWhere(*where)
	getAlias() string
	setAlias(string)
	getModelStatus() modelStatus
	addModelStatus(modelStatus)
	setModelStatus(modelStatus)
	removeModelStatus(modelStatus)
	SelectDistinct()
	setModel(Model)
	rawFullTabName() string
	fullTabName() string
}

type ModelList interface {
	Model
	NewModel() Model
	SetTotal(int)
	Len() int
}
