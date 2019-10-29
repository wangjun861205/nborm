package nborm

import (
	"fmt"
	"io"
)

// BaseField 基础字段
type BaseField interface {
	Model
	id() string
	colName() string
	setCol(string)
	fieldName() string
	setField(string)
	getStatus() fieldStatus
	setStatus(fieldStatus)
	addStatus(fieldStatus)
	removeStatus(fieldStatus)
	checkFieldStatus(status fieldStatus) bool
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
	isAutoInc() bool
	isPk() bool
	isUk() bool
	clauser
	// referencer
}

type fieldStatus int

const (
	invalid fieldStatus = 0
	valid   fieldStatus = 1
	// notNull     fieldStatus = 1 << 1
	null        fieldStatus = 1 << 1
	primaryKey  fieldStatus = 1 << 2
	autoInc     fieldStatus = 1 << 3
	forSelect   fieldStatus = 1 << 4
	forSum      fieldStatus = 1 << 5
	forAscOrder fieldStatus = 1 << 6
	forDscOrder fieldStatus = 1 << 7
	forGroup    fieldStatus = 1 << 9
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
	// return !(f.status&notNull == notNull)
	return f.checkFieldStatus(null)
}

// SetNull 设置为空值
func (f *baseField) SetNull() {
	f.addStatus(valid)
	f.addStatus(null)
	// f.removeStatus(notNull)
}

func (f *baseField) unsetNull() {
	// f.addStatus(notNull)
	f.status &^= null

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

func (f *baseField) isAutoInc() bool {
	return f.checkFieldStatus(autoInc)
}

func (f *baseField) isPk() bool {
	for _, pk := range f.PrimaryKey() {
		if pk.colName() == f.colName() {
			return true
		}
	}
	return false
}

func (f *baseField) isUk() bool {
	for _, uks := range f.UniqueKeys() {
		for _, uk := range uks {
			if uk.colName() == f.colName() {
				return true
			}
		}
	}
	return false
}
