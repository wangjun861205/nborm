package nborm

// Field Field
type Field interface {
	ClauseField
	ValueField
	dup() Field
	Init(Model Model, colName, fieldName, formName string, index int)
}

// FieldInfo Field的基本信息
type FieldInfo struct {
	ColName   string
	FieldName string
	Field     Field
	Index     int
}

// FieldInfoList FieldInfo的列表
type FieldInfoList []FieldInfo

// FieldList Field的列表
type FieldList []Field

// String 字符串Field
type String struct {
	clauseField
	stringValueField
}

// Init 初始化方法
func (f *String) Init(model Model, colName, fieldName, formName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.stringValueField
	}
	f.stringValueField.init(model, colName, fieldName, formName, index)
}

// String 实现Stringer接口
func (f String) String() string {
	return f.stringValueField.String()
}

func (f *String) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *String) dup() Field {
	nf := *f
	return &nf
}

//=============================================================================================================

// Int 整数型字段
type Int struct {
	clauseField
	intValueField
}

// Init 初始化方法
func (f *Int) Init(model Model, colName, fieldName, formName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.intValueField
	}
	f.intValueField.init(model, colName, fieldName, formName, index)
}

func (f Int) String() string {
	return f.intValueField.String()
}

func (f *Int) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *Int) dup() Field {
	nf := *f
	return &nf
}

//=======================================================================================================

// Date 日期型字段
type Date struct {
	clauseField
	dateValueField
}

// Init 初始化方法
func (f *Date) Init(model Model, colName, fieldName, formName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.dateValueField
	}
	f.dateValueField.init(model, colName, fieldName, formName, index)
}

func (f Date) String() string {
	return f.dateValueField.String()
}

func (f *Date) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *Date) dup() Field {
	nf := *f
	return &nf
}

//=========================================================================================

// Datetime 日期时间型字段
type Datetime struct {
	clauseField
	datetimeValueField
}

// Init 初始化方法
func (f *Datetime) Init(model Model, colName, fieldName, formName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.datetimeValueField
	}
	f.datetimeValueField.init(model, colName, fieldName, formName, index)
}

func (f Datetime) String() string {
	return f.datetimeValueField.String()
}

func (f *Datetime) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *Datetime) dup() Field {
	nf := *f
	return &nf
}

//=============================================================================================================

// Decimal 浮点数字段
type Decimal struct {
	clauseField
	decimalValueField
}

// Init 初始化方法
func (f *Decimal) Init(model Model, colName, fieldName, formName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.decimalValueField
	}
	f.decimalValueField.init(model, colName, fieldName, formName, index)
}

func (f Decimal) String() string {
	return f.decimalValueField.String()
}

func (f *Decimal) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *Decimal) dup() Field {
	nf := *f
	return &nf
}

// ===================================================================================================

// Time 日期型字段
type Time struct {
	clauseField
	timeValueField
}

// Init 初始化方法
func (f *Time) Init(model Model, colName, fieldName, formName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.timeValueField
	}
	f.timeValueField.init(model, colName, fieldName, formName, index)
}

func (f Time) String() string {
	return f.timeValueField.String()
}

func (f *Time) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *Time) dup() Field {
	nf := *f
	return &nf
}

// Bytes 二进制字段
type Bytes struct {
	clauseField
	byteValueField
}

// Init 初始化方法
func (f *Bytes) Init(model Model, colName, fieldName, formName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.byteValueField
	}
	f.byteValueField.init(model, colName, fieldName, formName, index)
}

func (f Bytes) String() string {
	return f.byteValueField.String()
}

func (f *Bytes) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *Bytes) dup() Field {
	nf := *f
	return &nf
}
