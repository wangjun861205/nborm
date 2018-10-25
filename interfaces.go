package nborm

import (
	"database/sql"
	"encoding/json"
)

type table interface {
	DB() string
	Tab() string
}

type Model interface {
	table
	Fields() []Field
}

type ModelList interface {
	New() Model
	Len() int
	Index(int) Model
	Delete(int)
	Swap(i, j int)
	json.Marshaler
	Model() Model
}

type Field interface {
	SQLVal() string
	IsValid() bool
	IsNull() bool
	sql.Scanner
	Column() string
	IsPk() bool
	IsInc() bool
	IsUni() bool
	InsertValuePair() [2]string
	UpdateValue() *UpdateValue
	Where() *Where
	LessFunc() func(Model, Model) int
	SetByUpdateValue(*UpdateValue)
	Invalidate()
	Super() Model
	SortOrder(reverse bool) string
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
