package nborm

import (
	"database/sql"
)

type table interface {
	DB() string
	Tab() string
}

//Field declare a Field interface
type Field interface {
	IsValid() bool
	IsNull() bool
	sql.Scanner
	dbName() string
	tabName() string
	columnName() string
	fullTabName() string
	fullColName() string
	isPk() bool
	isInc() bool
	isUni() bool
	where() *Where
	LessFunc(reverse bool) func(iaddr, jaddr uintptr) int
	SetByUpdateValue(*UpdateValue)
	Invalidate()
	SortOrder(reverse bool) string
	value() interface{}
	setVal(interface{}, bool)
	updateValue() *UpdateValue
	isNullable() bool
	getDefVal() interface{}
	check() error
	getFieldName() string
	getModelName() string
}

type relation interface {
	// joinClause() string
	where() *Where
	getSrcDB() string
	getRawSrcDB() string
	getSrcTab() string
	getRawSrcTab() string
	getSrcCol() string
	getRawSrcCol() string
	getDstDB() string
	getRawDstDB() string
	getDstTab() string
	getRawDstTab() string
	getDstCol() string
	getRawDstCol() string
	getFullSrcTab() string
	getFullDstTab() string
	getFullSrcCol() string
	getFullDstCol() string
	getSrcVal() interface{}
}

type complexRelation interface {
	relation
	getMidDB() string
	getRawMidDB() string
	getMidTab() string
	getRawMidTab() string
	getFullMidTab() string
	getMidLeftCol() string
	getRawMidLeftCol() string
	getMidRightCol() string
	getRawMidRightCol() string
	getFullMidLeftCol() string
	getFullMidRightCol() string
}
