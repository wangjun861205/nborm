package nborm

import "database/sql"

type table interface {
	DB() string
	Tab() string
}

type Model interface {
	table
	Fields() []Field
}

type modelList interface {
	table
	New() Model
	Len() int
	Index(int) Model
	Delete(int)
	Swap(i, j int)
}

type where interface {
	String() string
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
}

type relation interface {
	DstDB() string
	DstTab() string
	DstCol() string
	SrcCol() string
}
