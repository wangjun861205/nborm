package nborm

import (
	"fmt"
	"io"
)

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

type fieldStatus int

const (
	invalid     fieldStatus = 0
	valid       fieldStatus = 1
	notNull     fieldStatus = 1 << 1
	primaryKey  fieldStatus = 1 << 2
	autoInc     fieldStatus = 1 << 3
	forSelect   fieldStatus = 1 << 4
	forSum      fieldStatus = 1 << 5
	forAscOrder fieldStatus = 1 << 6
	forDscOrder fieldStatus = 1 << 7
	// forAgg      fieldStatus = 1 << 8
	forGroup fieldStatus = 1 << 9
)

type baseField struct {
	Model
	col    string
	field  string
	index  int
	status fieldStatus
}

func (f *baseField) init(model Model, colName, fieldName string, index int) {
	f.Model = model
	f.col = colName
	f.field = fieldName
	f.index = index
}

// func (f *baseField) String() string {
// 	if !f.IsValid() || f.IsNull() {
// 		return fmt.Sprintf("%s: %v", f.fieldName(), nil)
// 	}
// 	return fmt.Sprintf("%s: %v", f.fieldName(), f.Model.FieldInfos()[f.index].Field.value())
// }

func (f *baseField) colName() string {
	return f.col
}

func (f *baseField) setCol(col string) {
	f.col = col
}

func (f *baseField) fieldName() string {
	return f.field
}

func (f *baseField) setField(field string) {
	f.field = field
}

func (f *baseField) getStatus() fieldStatus {
	return f.status
}

func (f *baseField) setStatus(status fieldStatus) {
	f.status = status
}

func (f *baseField) addStatus(status fieldStatus) {
	f.status |= status
}

func (f *baseField) removeStatus(status fieldStatus) {
	f.status &^= status
}

func (f *baseField) checkFieldStatus(status fieldStatus) bool {
	return f.status&status > 0
}

// IsValid 是否有值
func (f *baseField) IsValid() bool {
	return f.status&valid == valid
}

func (f *baseField) setValid() {
	f.addStatus(valid)
}

func (f *baseField) unsetValid() {
	f.removeStatus(valid)
}

// IsNull 是否为空值
func (f *baseField) IsNull() bool {
	return !(f.status&notNull == notNull)
}

// SetNull 设置为空值
func (f *baseField) SetNull() {
	f.addStatus(valid)
	f.removeStatus(notNull)
}

func (f *baseField) unsetNull() {
	f.addStatus(notNull)
}

func (f *baseField) mustValid() {
	if !f.IsValid() {
		panic(fmt.Sprintf("invalid field (%s.%s.%s(%s))", f.DB(), f.Tab(), f.col, f.field))
	}
}

func (f *baseField) fullColName() string {
	if f.Model == nil {
		return f.col
	}
	if f.Model.getAlias() != "" {
		return fmt.Sprintf("%s.`%s`", f.Model.getAlias(), f.col)
	}
	return fmt.Sprintf("%s.`%s`", f.rawFullTabName(), f.col)
}

func (f *baseField) rawFullColName() string {
	return fmt.Sprintf("%s.`%s`", f.rawFullTabName(), f.col)
}

func (f *baseField) ForSelect() {
	f.appendSelector(f.Model.FieldInfos()[f.index].Field)
}

func (f *baseField) getFieldIndex() int {
	return f.index
}

func (f *baseField) id() string {
	return f.fieldName()
}

// ForSum 设置为总和字段
func (f *baseField) ForSum() {
	f.addStatus(forSelect)
	f.addStatus(forSum)
	f.addModelStatus(forModelAgg)
}

// AscOrder 设置为正序排序字段
func (f *baseField) AscOrder() {
	f.appendOrderBys(newOrderBy(f, asc))

}

// DscOrder 设置为倒序排序字段
func (f *baseField) DescOrder() {
	f.appendOrderBys(newOrderBy(f, desc))
}

// Distinct 设置为去重字段
// func (f *baseField) Distinct() {
// 	f.Model.addModelStatus(distinct)
// 	f.addStatus(forSelect)
// }

// CopyStatus CopyStatus
func (f *baseField) CopyStatus(dst Field) {
	dst.setStatus(f.status)
}

func (f *baseField) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(f.fullColName()))
	w.Write([]byte(" "))
}

func (f *baseField) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(f.rawFullColName()))
	w.Write([]byte(" "))
}

func (f *baseField) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(f.fullColName()))
	w.Write([]byte(" "))
}

func (f *baseField) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(f.rawFullColName()))
	w.Write([]byte(" "))
}

// GroupBy 设置为GroupBy字段
func (f *baseField) GroupBy() {
	f.appendGroupBys(f)
}

// String 字符串Field
type String struct {
	clauseField
	stringValueField
}

// Init 初始化方法
func (f *String) Init(model Model, colName, fieldName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.stringValueField
	}
	f.stringValueField.init(model, colName, fieldName, index)
}

// String 实现Stringer接口
func (f String) String() string {
	return f.stringValueField.String()
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
func (f *Int) Init(model Model, colName, fieldName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.intValueField
	}
	f.intValueField.init(model, colName, fieldName, index)
}

func (f Int) String() string {
	return f.intValueField.String()
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
func (f *Date) Init(model Model, colName, fieldName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.dateValueField
	}
	f.dateValueField.init(model, colName, fieldName, index)
}

func (f Date) String() string {
	return f.dateValueField.String()
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
func (f *Datetime) Init(model Model, colName, fieldName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.datetimeValueField
	}
	f.datetimeValueField.init(model, colName, fieldName, index)
}

func (f Datetime) String() string {
	return f.datetimeValueField.String()
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
func (f *Decimal) Init(model Model, colName, fieldName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.decimalValueField
	}
	f.decimalValueField.init(model, colName, fieldName, index)
}

func (f Decimal) String() string {
	return f.decimalValueField.String()
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
func (f *Time) Init(model Model, colName, fieldName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.timeValueField
	}
	f.timeValueField.init(model, colName, fieldName, index)
}

func (f Time) String() string {
	return f.timeValueField.String()
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
func (f *Bytes) Init(model Model, colName, fieldName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.byteValueField
	}
	f.byteValueField.init(model, colName, fieldName, index)
}

func (f Bytes) String() string {
	return f.byteValueField.String()
}

func (f *Bytes) dup() Field {
	nf := *f
	return &nf
}
