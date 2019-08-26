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
	getFieldIndex() int
}

type ClauseField interface {
	AndW() ClauseField
	AndWhere(string, interface{}) ClauseField
	OrWhere(string, interface{}) ClauseField
	U() ClauseField
	Update(interface{}) ClauseField
	Set(interface{}) ClauseField
}

// ValueField ValueField
type ValueField interface {
	BaseField
	Scan(interface{}) error
	value() interface{}
}

// Field Field
type Field interface {
	ClauseField
	ValueField
	dup() Field
	Init(Model, string, string, int)
	refClauser
}

type baseModel interface {
	SelectDistinct() Model
	SetForJoin() Model
	SetForLeftJoin() Model
	SetForRightJoin() Model
}

type clauseModel interface {
	AscOrderBy(refClauser) Model
	DescOrderBy(refClauser) Model
	AndExprWhere(*Expr) Model
	OrExprWhere(*Expr) Model
	AndModelWhereGroup(...*condition) Model
	OrModelWhereGroup(...*condition) Model
	AndHaving(*Expr) Model
	OrHaving(*Expr) Model
	AndHavingGroup(...*condition) Model
	OrHavingGroup(...*condition) Model
	ExprUpdate(*Expr) Model
	ExprGroupBy(refClauser) Model
	SelectAll() Model
	SelectFields(...Field) Model
	SelectExcept(...Field) Model
}

// Model Model
type Model interface {
	baseModel
	clauseModel
	DB() string
	Tab() string
	FieldInfos() FieldInfoList
	AutoIncField() Field
	PrimaryKey() FieldList
	UniqueKeys() []FieldList
	relations() RelationInfoList
	getAlias() string
	genAlias()
	getModelStatus() modelStatus
	addModelStatus(modelStatus)
	setModelStatus(modelStatus)
	removeModelStatus(modelStatus)
	checkStatus(modelStatus) bool
	rawFullTabName() string
	fullTabName() string
	getParent() Model
	setParent(Model)
	getIndex() int
	setIndex(int)
	genIndex() int
	appendWheres(*Expr)
	getWheres() exprList
	InitRel()
	SetLimit(int, int)
	getLimit() (int, int)
	getAggs() aggList
	setConList(ModelList)
	getConList() ModelList
	Collapse()
	setAggs(aggList)
	GetCache(string, time.Duration) bool
	SetCache(string)
	getGroupBys() []refClauser
	appendGroupBys(refClauser)
	appendSelectedFieldIndexes(int)
	getSelectedFieldIndexes() []int
	getOrderBys() []refClauser
	appendOrderBys(refClauser)
	appendHavings(*Expr)
	getHavings() exprList
	getInserts() exprList
	appendInserts(*Expr)
	getUpdates() exprList
}

type refClauser interface {
	toRefClause() string
	toSimpleRefClause() string
}

type clauser interface {
	toClause() (string, []interface{})
	toSimpleClause() (string, []interface{})
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
