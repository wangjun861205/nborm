package nborm

import (
	"encoding/json"
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

func (f *baseField) setModel(model Model) {
	f.Model = model
}

func (f *baseField) dbName() string {
	return f.DB()
}

func (f *baseField) tabName() string {
	return f.Tab()
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

// String 字符串Field
type String struct {
	baseField
	value string
}

// Scan 实现Scan接口
func (f *String) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case string:
		f.value = val
	case []byte:
		f.value = string(val)
	default:
		return fmt.Errorf("invalid type for scan String(%T)", v)
	}
	return nil
}

// MarshalJSON 实现接口
func (f String) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return json.Marshal(nil)
	}
	return json.Marshal(f.value)
}

// UnmarshalJSON 实现接口
func (f *String) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		f.value = ""
		return nil
	}
	f.unsetNull()
	f.addModelStatus(containValue)
	return json.Unmarshal(b, &f.value)
}

// Value 取值
func (f *String) Value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.value
}

// JSONValue JSON Marshal
func (f *String) JSONValue() interface{} {
	if !f.IsValid() || f.IsNull() {
		return nil
	}
	return f.value
}

// SetString 设置字符串值
func (f *String) SetString(v string) *String {
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.value = v
	return f
}

// Set 设置值
func (f *String) Set(v interface{}) Field {
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.value = v.(string)
	return f
}

// String 取值
func (f *String) String() string {
	return f.value
}

// AndW 按自身值来生成And Where
func (f *String) AndW() Field {
	f.AndExprWhere(NewExpr("@ = ?", f, f.Value()))
	return f
}

// OrW 按自身值来生成Or Where
func (f *String) OrW() Field {
	f.OrExprWhere(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *String) AndWhere(op string, value interface{}) Field {
	checkOp(op)
	f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, value))
	return f
}

func (f *String) OrWhere(op string, value interface{}) Field {
	checkOp(op)
	f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, value))
	return f
}

func (f *String) U() Field {
	f.ExprUpdate(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *String) Update(value interface{}) Field {
	f.ExprUpdate(NewExpr("@ = ?", f, value))
	return f
}

func (f *String) dup() Field {
	nf := *f
	return &nf
}

//=============================================================================================================

type Int struct {
	baseField
	value int
}

func (f *Int) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case int64:
		f.value = int(val)
	case []byte:
		i, err := strconv.ParseInt(string(val), 10, 64)
		if err != nil {
			return err
		}
		f.value = int(i)
	default:
		return fmt.Errorf("invalid type for scan Int(%T)", v)
	}
	return nil
}

func (f Int) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return json.Marshal(nil)
	}
	return json.Marshal(f.value)
}

func (f *Int) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		f.value = 0
		return nil
	}
	f.unsetNull()
	f.addModelStatus(containValue)
	return json.Unmarshal(b, &f.value)
}

func (f *Int) Value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.value
}

func (f *Int) JSONValue() interface{} {
	if !f.IsValid() || f.IsNull() {
		return nil
	}
	return f.value
}

func (f *Int) SetInt(v int) *Int {
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.value = v
	return f
}

func (f *Int) Set(v interface{}) Field {
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.value = v.(int)
	return f
}

func (f *Int) Int() int {
	return f.value
}

func (f *Int) AndW() Field {
	f.AndExprWhere(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *Int) OrW() Field {
	f.AndExprWhere(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *Int) AndWhere(op string, value interface{}) Field {
	checkOp(op)
	f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, value))
	return f
}

func (f *Int) OrWhere(op string, value interface{}) Field {
	checkOp(op)
	f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, value))
	return f
}

func (f *Int) U() Field {
	f.ExprUpdate(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *Int) Update(value interface{}) Field {
	f.ExprUpdate(NewExpr("@ = ?", f, value))
	return f
}

func (f *Int) dup() Field {
	nf := *f
	return &nf
}

//=======================================================================================================

type Date struct {
	baseField
	value time.Time
}

func (f *Date) Scan(v interface{}) error {
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
		f.value = t
	case string:
		t, err := time.Parse("2006-01-02", val)
		if err != nil {
			return err
		}
		f.value = t
	case time.Time:
		f.value = val
	default:
		return fmt.Errorf("invalid type for scan Date(%T)", v)
	}
	return nil
}

func (f *Date) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return json.Marshal(nil)
	}
	s := f.value.Format("2006-01-02")
	return json.Marshal(s)
}

func (f *Date) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		f.value = time.Time{}
		return nil
	}
	f.unsetNull()
	t, err := time.Parse("2006-01-02", strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}
	f.value = t
	f.addModelStatus(containValue)
	return nil
}

func (f *Date) Value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.value.Format("2006-01-02")
}

func (f *Date) JSONValue() interface{} {
	if !f.IsValid() || f.IsNull() {
		return nil
	}
	return f.value.Format("2006-01-02")
}

func (f *Date) SetDate(v time.Time) *Date {
	f.setValid()
	f.unsetNull()
	f.value = v
	f.addModelStatus(containValue)
	return f
}

func (f *Date) Set(v interface{}) Field {
	switch val := v.(type) {
	case string:
		t, err := time.Parse("2006-01-02", val)
		if err != nil {
			panic(err)
		}
		f.value = t
	case time.Time:
		f.value = val
	default:
		panic(fmt.Errorf("invalid value for Date.Set(): want string or time.Time got (%T)", v))
	}
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	return f
}

func (f *Date) Date() time.Time {
	return f.value
}

func (f *Date) AndW() Field {
	f.AndExprWhere(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *Date) OrW() Field {
	f.OrExprWhere(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *Date) AndWhere(op string, value interface{}) Field {
	checkOp(op)
	switch v := value.(type) {
	case time.Time:
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, v.Format("2006-01-02")))
	case []time.Time:
		strList := make([]string, 0, len(v))
		for _, t := range v {
			strList = append(strList, t.Format("2006-01-02"))
		}
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, strList))
	default:
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, v))
	}
	return f
}

func (f *Date) OrWhere(op string, value interface{}) Field {
	checkOp(op)
	switch v := value.(type) {
	case time.Time:
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, v.Format("2006-01-02")))
	case []time.Time:
		strList := make([]string, 0, len(v))
		for _, t := range v {
			strList = append(strList, t.Format("2006-01-02"))
		}
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, strList))
	default:
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, v))
	}
	return f
}

func (f *Date) U() Field {
	f.ExprUpdate(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *Date) Update(value interface{}) Field {
	f.ExprUpdate(NewExpr("@ = ?", f, value))
	return f
}

func (f *Date) dup() Field {
	nf := *f
	return &nf
}

//=========================================================================================

type Datetime struct {
	baseField
	value time.Time
}

func (f *Datetime) Scan(v interface{}) error {
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
		f.value = t
	case string:
		t, err := time.Parse("2006-01-02 15:04:05", val)
		if err != nil {
			return err
		}
		f.value = t
	case time.Time:
		f.value = val
	default:
		return fmt.Errorf("invalid type for scan Date(%T)", v)
	}
	return nil
}

func (f *Datetime) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return json.Marshal(nil)
	}
	s := f.value.Format("2006-01-02 15:04:05")
	return json.Marshal(s)
}

func (f *Datetime) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		f.value = time.Time{}
		return nil
	}
	f.unsetNull()
	t, err := time.Parse("2006-01-02 15:04:05", strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}
	f.value = t
	f.addModelStatus(containValue)
	return nil
}

func (f *Datetime) Value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.value.Format("2006-01-02 15:04:05")
}

func (f *Datetime) JSONValue() interface{} {
	if !f.IsValid() || f.IsNull() {
		return nil
	}
	return f.value.Format("2006-01-02 15:04:05")
}

func (f *Datetime) SetDatetime(v time.Time) *Datetime {
	f.setValid()
	f.unsetNull()
	f.value = v
	f.addModelStatus(containValue)
	return f
}

func (f *Datetime) Set(v interface{}) Field {
	switch val := v.(type) {
	case string:
		t, err := time.Parse("2006-01-02", val)
		if err != nil {
			panic(err)
		}
		f.value = t
	case time.Time:
		f.value = val
	default:
		panic(fmt.Errorf("invalid value for Date.Set(): want string or time.Time got (%T)", v))
	}
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	return f
}

func (f *Datetime) Datetime() time.Time {
	return f.value
}

func (f *Datetime) AndW() Field {
	f.mustValid()
	f.AndExprWhere(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *Datetime) OrW() Field {
	f.mustValid()
	f.OrExprWhere(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *Datetime) AndWhere(op string, value interface{}) Field {
	checkOp(op)
	switch v := value.(type) {
	case time.Time:
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, v.Format("2006-01-02 15:04:05")))
	case []time.Time:
		strList := make([]string, 0, len(v))
		for _, t := range v {
			strList = append(strList, t.Format("2006-01-02 15:04:05"))
		}
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, strList))
	default:
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, v))
	}
	return f
}

func (f *Datetime) OrWhere(op string, value interface{}) Field {
	checkOp(op)
	switch v := value.(type) {
	case time.Time:
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, v.Format("2006-01-02 15:04:05")))
	case []time.Time:
		strList := make([]string, 0, len(v))
		for _, t := range v {
			strList = append(strList, t.Format("2006-01-02 15:04:05"))
		}
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, strList))
	default:
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, v))
	}
	return f
}

func (f *Datetime) U() Field {
	f.ExprUpdate(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *Datetime) Update(value interface{}) Field {
	switch v := value.(type) {
	case time.Time:
		f.ExprUpdate(NewExpr("@ = ?", f, v.Format("2006-01-02 15:04:05")))
	default:
		f.ExprUpdate(NewExpr("@ = ?", f, v))
	}
	return f
}

func (f *Datetime) dup() Field {
	nf := *f
	return &nf
}

//=============================================================================================================

type Decimal struct {
	baseField
	value   float64
	exprVal *Expr
}

func (f *Decimal) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case float64:
		f.value = val
	case float32:
		f.value = float64(val)
	default:
		return fmt.Errorf("invalid type for scan Decimal(%T)", v)
	}
	return nil
}

func (f *Decimal) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return json.Marshal(nil)
	}
	return json.Marshal(f.value)
}

func (f *Decimal) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		f.value = 0.0
		return nil
	}
	f.unsetNull()
	f.addModelStatus(containValue)
	return json.Unmarshal(b, &f.value)

}

func (f *Decimal) Value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.value
}

func (f *Decimal) JSONValue() interface{} {
	if !f.IsValid() || f.IsNull() {
		return nil
	}
	return f.value
}

func (f *Decimal) SetDecimal(v float64) *Decimal {
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.value = v
	return f
}

func (f *Decimal) Set(v interface{}) Field {
	f.setValid()
	f.unsetNull()
	f.value = v.(float64)
	f.addModelStatus(containValue)
	return f
}

func (f *Decimal) Decimal() float64 {
	return f.value
}

func (f *Decimal) AndW() Field {
	f.AndExprWhere(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *Decimal) OrW() Field {
	f.OrExprWhere(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *Decimal) AndWhere(op string, value interface{}) Field {
	checkOp(op)
	f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, value))
	return f
}

func (f *Decimal) OrWhere(op string, value interface{}) Field {
	checkOp(op)
	f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f, value))
	return f
}

func (f *Decimal) U() Field {
	f.ExprUpdate(NewExpr("@ = ?", f, f.Value()))
	return f
}

func (f *Decimal) Update(value interface{}) Field {
	f.ExprUpdate(NewExpr("@ = ?", f, value))
	return f
}

func (f *Decimal) dup() Field {
	nf := *f
	return &nf
}
