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
	mustValid()
	rawFullTabName() string
	fullTabName() string
	rawFullColName() string
	fullColName() string
	ForSelect()
	ForSum()
	AscOrder()
	DscOrder()
	// String() string
	addModelStatus(modelStatus)
	GroupBy()
}

// Field Field
type Field interface {
	BaseField
	Scan(interface{}) error
	Value() interface{}
	Set(interface{}) Field
	AndW() Field
	OrW() Field
	AndWhere(string, interface{}) Field
	OrWhere(string, interface{}) Field
	U() Field
	Update(interface{}) Field
	dup() Field
}

// Model Model
type Model interface {
	DB() string
	Tab() string
	FieldInfos() FieldInfoList
	AutoIncField() Field
	PrimaryKey() FieldList
	Relations() RelationInfoList
	setRel(string, *Expr)
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
	checkStatus(modelStatus) bool
	SelectDistinct()
	setModel(Model)
	rawFullTabName() string
	fullTabName() string
	GetParent() Model
	SetParent(Model)
	getIndex() int
	getWheres() exprList
	getJoinWheres() exprList
	getHavings() exprList
	InitRel()
	SetLimit(int, int)
	getLimit() (int, int)
	getAggExps() []*aggExp
	ExprUpdate(*Expr)
	getUpdateList() exprList
	getConList() ModelList
	Collapse()
	setJoinClause(string)
	getJoinClause() string
	getRevJoinWheres() exprList
	appendWhere(...*Expr)
	appendJoinWhere(...*Expr)
	AndExprWhere(*Expr) Model
	OrExprWhere(*Expr) Model
	getOnCond() *Expr
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
