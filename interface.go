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

type baseModel interface {
	SetForJoin() Model
	SetForLeftJoin() Model
	SetForRightJoin() Model
	SetForDelete() Model
	referencer
}

type clauseModel interface {
	SelectDistinct() Model
	AscOrderBy(referencer) Model
	DescOrderBy(referencer) Model
	AndExprWhere(*Expr) Model
	OrExprWhere(*Expr) Model
	AndModelWhereGroup(wheres ...wherer) Model
	OrModelWhereGroup(wheres ...wherer) Model
	AndExprHaving(*Expr) Model
	OrExprHaving(*Expr) Model
	AndHavingGroup(...havinger) Model
	OrHavingGroup(...havinger) Model
	ExprUpdate(referencer, *Expr) Model
	ModelGroupBy(referencer) Model
	SelectAll() Model
	SelectFields(...Field) Model
	SelectExcept(...Field) Model
	GroupBySelectedFields() Model
	appendSelector(s selector)
	getSelectors() *selectorList
	appendAgg(agg aggregator)
	appendUpdate(update *update) Model
	ClearSelect() Model
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
	getWheres() wherer
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
	getGroupBys() groupByList
	appendGroupBys(referencer)
	getOrderBys() orderByList
	appendOrderBys(*orderBy)
	appendHavings(havinger)
	getHavings() havinger
	getInserts() insertList
	getUpdates() updateList
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
