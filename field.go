package nborm

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
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
	f.appendSelectedFieldIndexes(f.index)
}

func (f *baseField) getFieldIndex() int {
	return f.index
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
func (f *baseField) Distinct() {
	f.Model.addModelStatus(distinct)
	f.addStatus(forSelect)
}

// CopyStatus CopyStatus
func (f *baseField) CopyStatus(dst Field) {
	dst.setStatus(f.status)
}

func (f *baseField) toRefClause() string {
	return f.fullColName()
}

func (f *baseField) toSimpleRefClause() string {
	return f.rawFullColName()
}

// GroupBy 设置为GroupBy字段
func (f *baseField) GroupBy() {
	f.appendGroupBys(f)
}

type clauseField struct {
	valueField func() ValueField
}

// AndW 按自身值来生成And Where
func (f *clauseField) AndW() ClauseField {
	valueField := f.valueField()
	valueField.AndExprWhere(NewExpr("@ = ?", valueField, valueField.value()))
	return f
}

// OrW 按自身值来生成Or Where
func (f *clauseField) OrW() ClauseField {
	valueField := f.valueField()
	valueField.OrExprWhere(NewExpr("@ = ?", valueField, valueField.value()))
	return f
}

// U 按自身值来生成更新表达式
func (f *clauseField) U() ClauseField {
	valueField := f.valueField()
	valueField.ExprUpdate(NewExpr("@ = ?", valueField, valueField.value()))
	return f
}

func (f *clauseField) AndWhere(op string, value interface{}) ClauseField {
	checkOp(op)
	valueField := f.valueField()
	valueField.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), valueField, value))
	return f
}

func (f *clauseField) OrWhere(op string, value interface{}) ClauseField {
	checkOp(op)
	valueField := f.valueField()
	valueField.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), valueField, value))
	return f
}

func (f *clauseField) AndEq(value interface{}) *condition {
	return newCondition(and, NewExpr("@ = ?", f.valueField(), value))
}

func (f *clauseField) OrEq(value interface{}) *condition {
	return newCondition(or, NewExpr("@ = ?", f.valueField(), value))
}

func (f *clauseField) AndNeq(value interface{}) *condition {
	return newCondition(and, NewExpr("@ <> ?", f.valueField(), value))
}

func (f *clauseField) OrNeq(value interface{}) *condition {
	return newCondition(or, NewExpr("@ <> ?", f.valueField(), value))
}

func (f *clauseField) AndLt(value interface{}) *condition {
	return newCondition(and, NewExpr("@ < ?", f.valueField(), value))
}

func (f *clauseField) OrLt(value interface{}) *condition {
	return newCondition(or, NewExpr("@ < ?", f.valueField(), value))
}

func (f *clauseField) AndLte(value interface{}) *condition {
	return newCondition(and, NewExpr("@ <= ?", f.valueField(), value))
}

func (f *clauseField) OrLte(value interface{}) *condition {
	return newCondition(or, NewExpr("@ <= ?", f.valueField(), value))
}

func (f *clauseField) AndGt(value interface{}) *condition {
	return newCondition(and, NewExpr("@ > ?", f.valueField(), value))
}

func (f *clauseField) OrGt(value interface{}) *condition {
	return newCondition(or, NewExpr("@ > ?", f.valueField(), value))
}

func (f *clauseField) AndGte(value interface{}) *condition {
	return newCondition(and, NewExpr("@ >= ?", f.valueField(), value))
}

func (f *clauseField) OrGte(value interface{}) *condition {
	return newCondition(or, NewExpr("@ >= ?", f.valueField(), value))
}

func (f *clauseField) AndIsNull() *condition {
	return newCondition(and, NewExpr("@ IS NULL", f))
}

func (f *clauseField) OrIsNull() *condition {
	return newCondition(or, NewExpr("@ IS NULL", f))
}

func (f *clauseField) AndIsNotNull() *condition {
	return newCondition(and, NewExpr("@ IS NOT NULL", f))
}

func (f *clauseField) OrIsNotNull() *condition {
	return newCondition(or, NewExpr("@ IS NOT NULL", f))
}

func (f *clauseField) AndIn(value interface{}) *condition {
	return newCondition(and, NewExpr("@ IN ?", f.valueField(), value))
}

func (f *clauseField) OrIn(value interface{}) *condition {
	return newCondition(or, NewExpr("@ NOT IN ?", f.valueField(), value))
}

func (f *clauseField) AndLike(value interface{}) *condition {
	return newCondition(and, NewExpr("@ LIKE ?", f.valueField(), value))
}

func (f *clauseField) OrLike(value interface{}) *condition {
	return newCondition(or, NewExpr("@ LIKE ?", f.valueField(), value))
}

func (f *clauseField) AndNotLike(value interface{}) *condition {
	return newCondition(and, NewExpr("@ NOT LIKE ?", f.valueField(), value))
}

func (f *clauseField) OrNotLike(value interface{}) *condition {
	return newCondition(or, NewExpr("@ NOT LIKE ?", f.valueField(), value))
}

func (f *clauseField) AndBetween(startValue, endValue interface{}) *condition {
	return newCondition(and, NewExpr("@ BETWEEN ? AND ?", f.valueField(), startValue, endValue))
}

func (f *clauseField) OrBetween(startValue, endValue interface{}) *condition {
	return newCondition(or, NewExpr("@ BETWEEN ? AND ?", f.valueField(), startValue, endValue))
}

func (f *clauseField) AndNotBetween(startValue, endValue interface{}) *condition {
	return newCondition(and, NewExpr("@ NOT BETWEEN ? AND ?", f.valueField(), startValue, endValue))
}

func (f *clauseField) OrNotBetween(startValue, endValue interface{}) *condition {
	return newCondition(or, NewExpr("@ NOT BETWEEN ? AND ?", f.valueField(), startValue, endValue))
}

func (f *clauseField) AndWhereGroup(funcs ...func(*clauseField) *condition) {
	switch len(funcs) {
	case 0:
		return
	case 1:
		f.valueField().appendWheres(funcs[0](f).toExpr())
	default:
		l := make(conditionList, 0, len(funcs))
		for _, fn := range funcs {
			l = append(l, fn(f))
		}
		f.valueField().appendWheres(l.group(and).toExpr())
	}
}

func (f *clauseField) OrWhereGroup(funcs ...func(*clauseField) *condition) {
	switch len(funcs) {
	case 0:
		return
	case 1:
		f.valueField().appendWheres(funcs[0](f).toExpr())
	default:
		l := make(conditionList, 0, len(funcs))
		for _, fn := range funcs {
			l = append(l, fn(f))
		}
		f.valueField().appendWheres(l.group(or).toExpr())
	}
}

func (f *clauseField) Update(value interface{}) ClauseField {
	valueField := f.valueField()
	valueField.ExprUpdate(NewExpr("@ = ?", valueField, value))
	return f
}

func (f *clauseField) Set(value interface{}) ClauseField {
	f.valueField().appendInserts(NewExpr("@ = ?", f, value))
	return f
}

func (f *clauseField) toRefClause() string {
	return f.valueField().fullColName()
}

func (f *clauseField) toSimpleRefClause() string {
	return f.valueField().rawFullColName()
}

type stringValueField struct {
	baseField
	val string
}

func (f *stringValueField) init(model Model, colName, fieldName string, index int) {
	f.baseField.init(model, colName, fieldName, index)
}

// Scan 实现Scan接口
func (f *stringValueField) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case string:
		f.val = val
	case []byte:
		f.val = string(val)
	default:
		return fmt.Errorf("invalid type for scan String(%T)", v)
	}
	return nil
}

func (f *stringValueField) Value() (string, bool) {
	f.mustValid()
	if f.IsNull() {
		return "", true
	}
	return f.val, false
}

func (f *stringValueField) value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.val
}

// SetString 设置字符串值
func (f *stringValueField) SetString(v string) *stringValueField {
	f.appendInserts(NewExpr("@ = ?", f, v))
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.val = v
	return f
}

func (f stringValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%q", f.val)), nil
}

func (f *stringValueField) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.removeStatus(notNull)
		return nil
	}
	f.addStatus(notNull)
	f.val = strings.Trim(string(b), "\"")
	return nil
}

func (f *stringValueField) AnyValue() string {
	return f.val
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

func (f *String) dup() Field {
	nf := *f
	return &nf
}

//=============================================================================================================

type intValueField struct {
	baseField
	val int
}

func (f *intValueField) init(model Model, colName, fieldName string, index int) {
	f.baseField.init(model, colName, fieldName, index)
}

func (f *intValueField) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case int64:
		f.val = int(val)
	case []byte:
		i, err := strconv.ParseInt(string(val), 10, 64)
		if err != nil {
			return err
		}
		f.val = int(i)
	default:
		return fmt.Errorf("invalid type for scan Int(%T)", v)
	}
	return nil
}

func (f *intValueField) value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.val
}

func (f *intValueField) Value() (int, bool) {
	f.mustValid()
	if f.IsNull() {
		return 0, true
	}
	return f.val, false
}

func (f *intValueField) SetInt(v int) *intValueField {
	f.appendInserts(NewExpr("@ = ?", f, v))
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.val = v
	return f
}

// func (f *intValueField) Set(value interface{}) ValueField {
// 	f.appendInserts(NewExpr("@ = ?", f, value))
// 	if v, ok := value.(int); ok {
// 		f.setValid()
// 		f.unsetNull()
// 		f.addModelStatus(containValue)
// 		f.val = v
// 	}
// 	return f
// }

func (f intValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%d", f.val)), nil
}

func (f intValueField) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		// f.SetNull()
		f.removeStatus(notNull)
		return nil
	}
	v, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	// f.SetInt(int(v))
	f.val = int(v)
	f.addStatus(notNull)
	return nil
}

func (f *intValueField) AnyValue() int {
	return f.val
}

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

func (f *Int) dup() Field {
	nf := *f
	return &nf
}

//=======================================================================================================

type dateValueField struct {
	baseField
	val time.Time
}

func (f *dateValueField) init(model Model, colName, fieldName string, index int) {
	f.baseField.init(model, colName, fieldName, index)
}

func (f *dateValueField) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case []byte:
		t, err := time.Parse("2006-01-02", string(val))
		if err != nil {
			return err
		}
		f.val = t
	case string:
		t, err := time.Parse("2006-01-02", val)
		if err != nil {
			return err
		}
		f.val = t
	case time.Time:
		f.val = val
	default:
		return fmt.Errorf("invalid type for scan Date(%T)", v)
	}
	return nil
}

func (f *dateValueField) Value() (time.Time, bool) {
	f.mustValid()
	if f.IsNull() {
		return time.Time{}, true
	}
	return f.val, false
}

func (f *dateValueField) value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.val.Format("2006-01-02")
}

func (f *dateValueField) SetDate(v time.Time) *dateValueField {
	f.appendInserts(NewExpr("@ = ?", f, v.Format("2006-01-02")))
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.val = v
	return f
}

// func (f *dateValueField) Set(value interface{}) ValueField {
// 	f.appendInserts(NewExpr("@ = ?", f, value))
// 	if v, ok := value.(time.Time); ok {
// 		f.setValid()
// 		f.unsetNull()
// 		f.addModelStatus(containValue)
// 		f.val = v
// 	}
// 	return f
// }

func (f dateValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%q", f.val.Format("2006-01-02"))), nil
}

func (f *dateValueField) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		// f.SetNull()
		f.removeStatus(notNull)
		return nil
	}
	t, err := time.Parse("2006-01-02", strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}
	// f.SetDate(t)
	f.val = t
	f.addStatus(notNull)
	return nil
}

func (f *dateValueField) AnyValue() time.Time {
	return f.val
}

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

func (f *Date) dup() Field {
	nf := *f
	return &nf
}

//=========================================================================================

type datetimeValueField struct {
	baseField
	val time.Time
}

func (f *datetimeValueField) init(model Model, colName, fieldName string, index int) {
	f.baseField.init(model, colName, fieldName, index)
}

func (f *datetimeValueField) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case []byte:
		t, err := time.Parse("2006-01-02 15:04:05", string(val))
		if err != nil {
			return err
		}
		f.val = t
	case string:
		t, err := time.Parse("2006-01-02 15:04:05", val)
		if err != nil {
			return err
		}
		f.val = t
	case time.Time:
		f.val = val
	default:
		return fmt.Errorf("invalid type for scan Date(%T)", v)
	}
	return nil
}

func (f *datetimeValueField) Value() (time.Time, bool) {
	f.mustValid()
	if f.IsNull() {
		return time.Time{}, true
	}
	return f.val, false
}

func (f *datetimeValueField) value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.val.Format("2006-01-02 15:04:05")
}

func (f *datetimeValueField) SetDatetime(v time.Time) *datetimeValueField {
	f.appendInserts(NewExpr("@ = ?", f, v.Format("2006-01-02 15:04:05")))
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.val = v
	return f
}

// func (f *datetimeValueField) Set(value interface{}) ValueField {
// 	f.appendInserts(NewExpr("@ = ?", f, value))
// 	if v, ok := value.(time.Time); ok {
// 		f.setValid()
// 		f.unsetNull()
// 		f.addModelStatus(containValue)
// 		f.val = v
// 	}
// 	return f
// }

func (f datetimeValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%q", f.val.Format("2006-01-02 15:04:05"))), nil
}

func (f *datetimeValueField) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		// f.SetNull()
		f.removeStatus(notNull)
		return nil
	}
	t, err := time.Parse("2006-01-02 15:04:05", strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}
	// f.SetDatetime(t)
	f.removeStatus(notNull)
	f.val = t
	return nil

}

func (f *datetimeValueField) AnyValue() time.Time {
	return f.val
}

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

func (f *Datetime) dup() Field {
	nf := *f
	return &nf
}

//=============================================================================================================

type decimalValueField struct {
	baseField
	val float64
}

func (f *decimalValueField) init(model Model, colName, fieldName string, index int) {
	f.baseField.init(model, colName, fieldName, index)
}

func (f *decimalValueField) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case float64:
		f.val = val
	case float32:
		f.val = float64(val)
	case []byte:
		fv, err := strconv.ParseFloat(string(val), 64)
		if err != nil {
			return err
		}
		f.val = fv
	default:
		return fmt.Errorf("invalid type for scan Decimal(%T)", v)
	}
	return nil
}

func (f *decimalValueField) Value() (float64, bool) {
	f.mustValid()
	if f.IsNull() {
		return 0, true
	}
	return f.val, false
}

func (f *decimalValueField) value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.val
}

func (f *decimalValueField) SetDecimal(v float64) *decimalValueField {
	f.appendInserts(NewExpr("@ = ?", v))
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.val = v
	return f
}

func (f decimalValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%f", f.val)), nil
}

func (f *decimalValueField) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		// f.SetNull()
		f.removeStatus(notNull)
		return nil
	}
	d, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return err
	}
	// f.SetDecimal(d)
	f.addStatus(notNull)
	f.val = d
	return nil
}

func (f *decimalValueField) AnyValue() float64 {
	return f.val
}

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

func (f *Decimal) dup() Field {
	nf := *f
	return &nf
}

// ===================================================================================================

type timeValueField struct {
	baseField
	val time.Time
}

func (f *timeValueField) init(model Model, colName, fieldName string, index int) {
	f.baseField.init(model, colName, fieldName, index)
}

func (f *timeValueField) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case []byte:
		t, err := time.Parse("15:04:05", string(val))
		if err != nil {
			return err
		}
		f.val = t
	case string:
		t, err := time.Parse("15:04:05", val)
		if err != nil {
			return err
		}
		f.val = t
	case time.Time:
		f.val = val
	default:
		return fmt.Errorf("invalid type for scan timeValueField(%T)", v)
	}
	return nil
}

func (f *timeValueField) Value() (time.Time, bool) {
	f.mustValid()
	if f.IsNull() {
		return time.Time{}, true
	}
	return f.val, false
}

func (f *timeValueField) AnyValue() time.Time {
	return f.val
}

func (f *timeValueField) value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.val.Format("15:04:05")
}

func (f *timeValueField) SetTime(v time.Time) *timeValueField {
	f.appendInserts(NewExpr("@ = ?", f, v.Format("15:04:05")))
	f.setValid()
	f.unsetNull()
	f.val = v
	f.addModelStatus(containValue)
	return f
}

func (f *timeValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(f.val.Format("15:04:05")), nil
}

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

func (f *Time) dup() Field {
	nf := *f
	return &nf
}

type byteValueField struct {
	baseField
	val []byte
}

func (f *byteValueField) init(model Model, colName, fieldName string, index int) {
	f.baseField.init(model, colName, fieldName, index)
}

func (f *byteValueField) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case []byte:
		f.val = val
	case string:
		bs, err := hex.DecodeString(val)
		if err != nil {
			return err
		}
		f.val = bs
	default:
		return fmt.Errorf("invalid type for scan Decimal(%T)", v)
	}
	return nil
}

func (f *byteValueField) Value() ([]byte, bool) {
	f.mustValid()
	if f.IsNull() {
		return nil, true
	}
	return f.val, false
}

func (f *byteValueField) value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.val
}

func (f *byteValueField) SetBytes(v []byte) *byteValueField {
	f.appendInserts(NewExpr("@ = ?", v))
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.val = v
	return f
}

func (f byteValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%x", f.val)), nil
}

func (f *byteValueField) AnyValue() []byte {
	return f.val
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

func (f *Bytes) dup() Field {
	nf := *f
	return &nf
}
