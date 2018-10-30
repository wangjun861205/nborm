package nborm

import (
	"database/sql"
	"encoding/json"
)

type table interface {
	DB() string
	Tab() string
}

//Model declare a Model interface
type Model interface {
	table
	Fields() []Field
	SetSync(bool)
}

//ModelList declare a ModelList interface
type ModelList interface {
	New() Model
	Len() int
	Index(int) Model
	Delete(int)
	Swap(i, j int)
	json.Marshaler
	Model() Model
}

//Field declare a Field interface
type Field interface {
	IsValid() bool
	IsNull() bool
	sql.Scanner
	Column() string
	IsPk() bool
	IsInc() bool
	IsUni() bool
	Where() *Where
	LessFunc() func(Model, Model) int
	SetByUpdateValue(*UpdateValue)
	Invalidate()
	Super() Model
	SortOrder(reverse bool) string
	value() interface{}
}

type relation interface {
	json.Marshaler
	srcDB() string
	srcTab() string
	srcCol() string
	dstDB() string
	dstTab() string
	dstCol() string
	joinClause() string
}
