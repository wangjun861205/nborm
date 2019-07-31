package nborm

import (
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
	forAgg      fieldStatus = 1 << 8
	forGroup    fieldStatus = 1 << 9
)

type baseField struct {
	Model
	col    string
	field  string
	status fieldStatus
}

func (f *baseField) init(model Model, colName, fieldName string) {
	f.Model = model
	f.col = colName
	f.field = fieldName
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

// ForSelect 设置为选择字段
func (f *baseField) ForSelect() {
	f.addStatus(forSelect)
}

// ForSum 设置为总和字段
func (f *baseField) ForSum() {
	f.addStatus(forSelect)
	f.addStatus(forSum)
	f.addModelStatus(forModelAgg)
}

// AscOrder 设置为正序排序字段
func (f *baseField) AscOrder() {
	f.removeStatus(forDscOrder)
	f.addStatus(forAscOrder)
	f.addModelStatus(forModelOrder)
	for parent := f.getParent(); parent != nil; parent = parent.getParent() {
		parent.addModelStatus(containSubOrder)
	}
}

// DscOrder 设置为倒序排序字段
func (f *baseField) DscOrder() {
	f.removeStatus(forAscOrder)
	f.addStatus(forDscOrder)
	f.addModelStatus(forModelOrder)
	for parent := f.getParent(); parent != nil; parent = parent.getParent() {
		parent.addModelStatus(containSubOrder)
	}
}

// Distinct 设置为去重字段
func (f *baseField) Distinct() {
	f.Model.addModelStatus(distinct)
	f.addStatus(forSelect)
}

// GroupBy 分组
func (f *baseField) GroupBy() {
	f.addStatus(forGroup | forSelect)
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

func (f *clauseField) Update(value interface{}) ClauseField {
	valueField := f.valueField()
	valueField.ExprUpdate(NewExpr("@ = ?", valueField, value))
	return f
}

type stringValueField struct {
	baseField
	val string
}

func (f *stringValueField) init(model Model, colName, fieldName string) {
	f.baseField.init(model, colName, fieldName)
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
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.val = v
	return f
}

func (f *stringValueField) set(v interface{}) ValueField {
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.val = v.(string)
	return f
}

func (f stringValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() {
		return nil, fmt.Errorf("invalid field (%s)", f.rawFullColName())
	}
	if f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%q", f.val)), nil
}

func (f *stringValueField) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		// f.SetNull()
		f.removeStatus(notNull)
		return nil
	}
	// f.SetString(strings.Trim(string(b), "\""))
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

func (f *String) Init(model Model, colName, fieldName string) {
	f.clauseField.valueField = func() ValueField {
		return &f.stringValueField
	}
	f.stringValueField.init(model, colName, fieldName)
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

func (f *intValueField) init(model Model, colName, fieldName string) {
	f.baseField.init(model, colName, fieldName)
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
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.val = v
	return f
}

func (f *intValueField) set(v interface{}) ValueField {
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.val = v.(int)
	return f
}

func (f intValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() {
		return nil, fmt.Errorf("invalid field (%s)", f.rawFullColName())
	}
	if f.IsNull() {
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

type Int struct {
	clauseField
	intValueField
}

func (f *Int) Init(model Model, colName, fieldName string) {
	f.clauseField.valueField = func() ValueField {
		return &f.intValueField
	}
	f.intValueField.init(model, colName, fieldName)
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

func (f *dateValueField) init(model Model, colName, fieldName string) {
	f.baseField.init(model, colName, fieldName)
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
	f.setValid()
	f.unsetNull()
	f.val = v
	f.addModelStatus(containValue)
	return f
}

func (f *dateValueField) set(v interface{}) ValueField {
	switch val := v.(type) {
	case string:
		t, err := time.Parse("2006-01-02", val)
		if err != nil {
			panic(err)
		}
		f.val = t
	case time.Time:
		f.val = val
	default:
		panic(fmt.Errorf("invalid value for Date.Set(): want string or time.Time got (%T)", v))
	}
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	return f
}

func (f dateValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() {
		return nil, fmt.Errorf("invalid field (%s)", f.rawFullColName())
	}
	if f.IsNull() {
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

type Date struct {
	clauseField
	dateValueField
}

func (f *Date) Init(model Model, colName, fieldName string) {
	f.clauseField.valueField = func() ValueField {
		return &f.dateValueField
	}
	f.dateValueField.init(model, colName, fieldName)
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

func (f *datetimeValueField) init(model Model, colName, fieldName string) {
	f.baseField.init(model, colName, fieldName)
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
	f.setValid()
	f.unsetNull()
	f.val = v
	f.addModelStatus(containValue)
	return f
}

func (f *datetimeValueField) set(v interface{}) ValueField {
	switch val := v.(type) {
	case string:
		t, err := time.Parse("2006-01-02 15:04:05", val)
		if err != nil {
			panic(err)
		}
		f.val = t
	case time.Time:
		f.val = val
	default:
		panic(fmt.Errorf("invalid value for Date.Set(): want string or time.Time got (%T)", v))
	}
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	return f
}

func (f datetimeValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() {
		return nil, fmt.Errorf("invalid field (%s)", f.rawFullColName())
	}
	if f.IsNull() {
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

type Datetime struct {
	clauseField
	datetimeValueField
}

func (f *Datetime) Init(model Model, colName, fieldName string) {
	f.clauseField.valueField = func() ValueField {
		return &f.datetimeValueField
	}
	f.datetimeValueField.init(model, colName, fieldName)
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

func (f *decimalValueField) init(model Model, colName, fieldName string) {
	f.baseField.init(model, colName, fieldName)
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
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.val = v
	return f
}

func (f *decimalValueField) set(v interface{}) ValueField {
	f.setValid()
	f.unsetNull()
	f.val = v.(float64)
	f.addModelStatus(containValue)
	return f
}

func (f decimalValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() {
		return nil, fmt.Errorf("invalid field (%s)", f.rawFullColName())
	}
	if f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%d", f.val)), nil
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

type Decimal struct {
	clauseField
	decimalValueField
}

func (f *Decimal) Init(model Model, colName, fieldName string) {
	f.clauseField.valueField = func() ValueField {
		return &f.decimalValueField
	}
	f.decimalValueField.init(model, colName, fieldName)
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

func (f *timeValueField) value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.val.Format("15:04:05")
}

func (f *timeValueField) SetTime(v time.Time) *timeValueField {
	f.setValid()
	f.unsetNull()
	f.val = v
	f.addModelStatus(containValue)
	return f
}

func (f *timeValueField) set(v interface{}) ValueField {
	switch val := v.(type) {
	case string:
		t, err := time.Parse("15:04:05", val)
		if err != nil {
			panic(err)
		}
		f.val = t
	case time.Time:
		f.val = val
	default:
		panic(fmt.Errorf("invalid value for Date.Set(): want string or time.Time got (%T)", v))
	}
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	return f
}
