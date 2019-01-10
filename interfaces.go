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
}

type relation interface {
	// joinClause() string
	where() *Where
	getSrcDB() string
	getSrcTab() string
	getDstDB() string
	getDstTab() string
	getFullSrcTab() string
	getFullDstTab() string
}
