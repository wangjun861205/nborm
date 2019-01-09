package nborm

import (
	"database/sql"
)

type table interface {
	DB() string
	Tab() string
}

//Model declare a Model interface
// type Model interface {
// 	table
// 	Fields() []Field
// 	SetSync(bool)
// }

//ModelList declare a ModelList interface
// type ModelList interface {
// 	table
// 	New() Model
// 	Len() int
// 	// Index(int) Model
// 	// Delete(int)
// 	Swap(i, j int)
// 	json.Marshaler
// 	Model() Model
// }

//Field declare a Field interface
type Field interface {
	IsValid() bool
	IsNull() bool
	sql.Scanner
	dbName() string
	tabName() string
	columnName() string
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
}

type relation interface {
	// joinClause() string
	where() *Where
	getSrcDB() string
	getSrcTab() string
	getDstDB() string
	getDstTab() string
}
