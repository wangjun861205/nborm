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
	mustValid()
	rawFullTabName() string
	fullTabName() string
	rawFullColName() string
	fullColName() string
	ForSelect()
	ForSum()
	AscOrder()
	DscOrder()
	String() string
	addModelStatus(modelStatus)
	GroupBy()
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
	SetU() Field
	SetUpdate(interface{}) Field
	// updateSet() *updateSet
	dup() Field
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
	getAlias() string
	setAlias()
	getModelStatus() modelStatus
	addModelStatus(modelStatus)
	setModelStatus(modelStatus)
	removeModelStatus(modelStatus)
	SelectDistinct()
	setModel(Model)
	rawFullTabName() string
	fullTabName() string
	AndExprWhere(*Expr, ...interface{}) Model
	OrExprWhere(*Expr, ...interface{}) Model
	GetParent() Model
	SetParent(Model)
	getIndex() int
	getWhere() *where
	getHaving() *where
	InitRel()
	SetLimit(int, int)
	getLimit() (int, int)
	getAggExps() []*aggExp
	appendUpdate(*update)
	getUpdateList() updateList
}

type ModelList interface {
	Model
	NewModel() Model
	SetTotal(int)
	Len() int
	Collapse()
}
