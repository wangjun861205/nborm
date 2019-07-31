package nborm

import (
	"database/sql"
	"time"
)

// RowScanner RowScanner
type RowScanner interface {
	ScanRow(*sql.Row) error
}

// RowsScanner RowsScanner
type RowsScanner interface {
	ScanRows(*sql.Rows) error
}

// Executor 可执行sql语句的对象
type Executor interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	Exec(string, ...interface{}) (sql.Result, error)
}

// BaseField 基础字段
type BaseField interface {
	Model
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
	mustValid()
	rawFullColName() string
	fullColName() string
	ForSelect()
	ForSum()
	AscOrder()
	DscOrder()
	GroupBy()
}

type ClauseField interface {
	AndW() ClauseField
	AndWhere(string, interface{}) ClauseField
	OrWhere(string, interface{}) ClauseField
	U() ClauseField
	Update(interface{}) ClauseField
}

// ValueField ValueField
type ValueField interface {
	BaseField
	Scan(interface{}) error
	value() interface{}
	set(interface{}) ValueField
}

// Field Field
type Field interface {
	ClauseField
	ValueField
	dup() Field
}

// Model Model
type Model interface {
	DB() string
	Tab() string
	FieldInfos() FieldInfoList
	AutoIncField() Field
	PrimaryKey() FieldList
	UniqueKeys() []FieldList
	relations() RelationInfoList
	getAlias() string
	setAlias()
	getModelStatus() modelStatus
	addModelStatus(modelStatus)
	setModelStatus(modelStatus)
	removeModelStatus(modelStatus)
	checkStatus(modelStatus) bool
	SelectDistinct()
	// setModel(Model)
	rawFullTabName() string
	fullTabName() string
	getParent() Model
	setParent(Model)
	getIndex() int
	getWheres() exprList
	getHavings() exprList
	InitRel()
	SetLimit(int, int)
	getLimit() (int, int)
	getAggExps() []*aggExp
	ExprUpdate(*Expr)
	getUpdateList() exprList
	setConList(ModelList)
	getConList() ModelList
	Collapse()
	appendWhere(...*Expr)
	AndExprWhere(*Expr) Model
	OrExprWhere(*Expr) Model
	GetCache(string, time.Duration) bool
	SetCache(string)
}

// ModelList ModelList
type ModelList interface {
	Model
	NewModel() Model
	SetTotal(int)
	GetTotal() int
	Len() int
	GetList() []Model
	Slice(int, int)
	GetListCache(string, time.Duration) bool
	SetListCache(string)
}
