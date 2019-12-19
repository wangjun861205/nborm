package nborm

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ValueField ValueField
type ValueField interface {
	BaseField
	Scan(interface{}) error
	value() *Expr
	toScan(m Model, selectors *[]interface{})
	set(v interface{}) ValueField
	update(v interface{}) ValueField
	setByReq(req *http.Request) error
	sqlLiteral() string
	getExpr() *Expr
	setByStr(str string) error
}

type stringValueField struct {
	baseField
	val string
	exp *Expr
}

func (f *stringValueField) init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.baseField.init(model, colName, fieldName, formName, uriName, index)
}

func (f *stringValueField) String() string {
	if !f.IsValid() || f.IsNull() {
		return fmt.Sprintf("%s: %v", f.fieldName(), nil)
	}
	return fmt.Sprintf("%s: %s", f.fieldName(), f.AnyValue())
}

func (f *stringValueField) toScan(m Model, selectors *[]interface{}) {
	fieldInfos := m.FieldInfos()
	for _, info := range fieldInfos {
		if info.Field.fieldName() == f.fieldName() {
			*selectors = append(*selectors, info.Field)
			return
		}
	}
	panic(fmt.Sprintf("cannot find field(%s) in %T", f.fieldName(), m))
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
		// return fmt.Errorf("invalid type for scan String(%T)", v)
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("String() invalid type for scan (value type: %T)", v), nil)
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

func (f *stringValueField) value() *Expr {
	f.mustValid()
	if f.IsNull() {
		return NewExpr("NULL")
	}
	if f.exp != nil {
		return f.exp
	}
	return NewExpr("?", f.val)
}

func (f *stringValueField) sqlLiteral() string {
	f.mustValid()
	if f.IsNull() {
		return "NULL"
	}
	if f.exp != nil {
		return f.exp.exp
	}
	return fmt.Sprintf(`"%s"`, f.val)
}

func (f *stringValueField) AnyValue() string {
	return f.val
}

func (f *stringValueField) getExpr() *Expr {
	return f.exp
}

func (f *stringValueField) SetExpr(exp *Expr) *stringValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = "", exp
	return f
}

// SetString 设置字符串值
func (f *stringValueField) SetString(v string) *stringValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = v, nil
	return f
}

func (f *stringValueField) set(v interface{}) ValueField {
	if exp, ok := v.(*Expr); ok {
		return f.SetExpr(exp)
	}
	return f.SetString(v.(string))
}

func (f *stringValueField) update(v interface{}) ValueField {
	f.appendUpdate(newUpdate(f, NewExpr("?", v)))
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
		f.SetNull()
		return nil
	}
	f.unsetNull()
	f.val = strings.Trim(string(b), "\"")
	return nil
}

func (f *stringValueField) setByReq(req *http.Request) error {
	if err := req.ParseForm(); err != nil {
		return newErr(ErrCodeIO, "setByReq() failed to parse request form", err)
	}
	m := map[string][]string(req.Form)
	key := f.formName()
	if len(m[key]) > 0 {
		f.SetString(m[key][0])
	}
	return nil
}

func (f *stringValueField) setByStr(str string) error {
	f.SetString(str)
	return nil
}

type intValueField struct {
	baseField
	val int
	exp *Expr
}

func (f *intValueField) init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.baseField.init(model, colName, fieldName, formName, uriName, index)
}

func (f *intValueField) String() string {
	if !f.IsValid() || f.IsNull() {
		return fmt.Sprintf("%s: %v", f.fieldName(), nil)
	}
	return fmt.Sprintf("%s: %d", f.fieldName(), f.AnyValue())
}

func (f *intValueField) toScan(m Model, selectors *[]interface{}) {
	fieldInfos := m.FieldInfos()
	for _, info := range fieldInfos {
		if info.Field.fieldName() == f.fieldName() {
			*selectors = append(*selectors, info.Field)
			return
		}
	}
	panic(fmt.Sprintf("cannot find field(%s) in %T", f.fieldName(), m))
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
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("intValueField invalid type for scan Int(%T)", v), nil)
	}
	return nil
}

func (f *intValueField) value() *Expr {
	f.mustValid()
	if f.IsNull() {
		return NewExpr("NULL")
	}
	if f.exp != nil {
		return f.exp
	}
	return NewExpr("?", f.val)
}

func (f *intValueField) sqlLiteral() string {
	f.mustValid()
	if f.IsNull() {
		return "NULL"
	}
	if f.exp != nil {
		return f.exp.exp
	}
	return fmt.Sprintf(`%d`, f.val)
}

func (f *intValueField) Value() (int, bool) {
	f.mustValid()
	if f.IsNull() {
		return 0, true
	}
	return f.val, false
}

func (f *intValueField) getExpr() *Expr {
	return f.exp
}

func (f *intValueField) SetExpr(exp *Expr) *intValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = 0, exp
	return f
}

func (f *intValueField) SetInt(v int) *intValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = v, nil
	return f
}

func (f *intValueField) set(v interface{}) ValueField {
	if exp, ok := v.(*Expr); ok {
		return f.SetExpr(exp)
	}
	return f.SetInt(v.(int))
}

func (f *intValueField) update(v interface{}) ValueField {
	f.appendUpdate(newUpdate(f, NewExpr("?", v)))
	return f
}

func (f intValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%d", f.val)), nil
}

func (f *intValueField) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		return nil
	}
	v, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	f.val = int(v)
	f.unsetNull()
	return nil
}

func (f *intValueField) setByReq(req *http.Request) error {
	if err := req.ParseForm(); err != nil {
		return newErr(ErrCodeIO, "intValueField failed to parse request form", err)
	}
	m := map[string][]string(req.Form)
	key := f.formName()
	if len(m[key]) > 0 {
		i64, err := strconv.ParseInt(m[key][0], 10, 64)
		if err != nil {
			return err
		}
		f.SetInt(int(i64))
	}
	return nil
}

func (f *intValueField) setByStr(str string) error {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("intValueField invalid value format for parser (value: %s)", str), err)
	}
	f.SetInt(int(i))
	return nil
}

func (f *intValueField) AnyValue() int {
	return f.val
}

type dateValueField struct {
	baseField
	val time.Time
	exp *Expr
}

func (f *dateValueField) init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.baseField.init(model, colName, fieldName, formName, uriName, index)
}

func (f *dateValueField) String() string {
	if !f.IsValid() || f.IsNull() {
		return fmt.Sprintf("%s: %v", f.fieldName(), nil)
	}
	return fmt.Sprintf("%s: %s", f.fieldName(), f.AnyValue().In(time.Local).Format("2006-01-02"))
}

func (f *dateValueField) toScan(m Model, selectors *[]interface{}) {
	fieldInfos := m.FieldInfos()
	for _, info := range fieldInfos {
		if info.Field.fieldName() == f.fieldName() {
			*selectors = append(*selectors, info.Field)
			return
		}
	}
	panic(fmt.Sprintf("cannot find field(%s) in %T", f.fieldName(), m))
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
		t, err := time.ParseInLocation("2006-01-02", string(val), time.Local)
		if err != nil {
			return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("dateValueField failed to parser bytes (value: %s)", string(val)), err)
		}
		f.val = t
	case string:
		t, err := time.ParseInLocation("2006-01-02", val, time.Local)
		if err != nil {
			return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("dateValueField failed to parse string (value: %s)", val), err)
		}
		f.val = t
	case time.Time:
		f.val = val
	default:
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("invalid type for scan Date(%T)", v), nil)
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

func (f *dateValueField) value() *Expr {
	f.mustValid()
	if f.IsNull() {
		return NewExpr("NULL")
	}
	if f.exp != nil {
		return f.exp
	}
	return NewExpr("?", f.val.In(time.Local).Format("2006-01-02"))
}

func (f *dateValueField) sqlLiteral() string {
	f.mustValid()
	if f.IsNull() {
		return "NULL"
	}
	if f.exp != nil {
		return f.exp.exp
	}
	return fmt.Sprintf(`"%s"`, f.val.Format("2006-01-02"))
}

func (f *dateValueField) getExpr() *Expr {
	return f.exp
}

func (f *dateValueField) SetExpr(exp *Expr) *dateValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = time.Time{}, exp
	return f
}

func (f *dateValueField) SetDate(v time.Time) *dateValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = v, nil
	return f
}

func (f *dateValueField) set(v interface{}) ValueField {
	if exp, ok := v.(*Expr); ok {
		return f.SetExpr(exp)
	}
	return f.SetDate(v.(time.Time))
}

func (f *dateValueField) update(v interface{}) ValueField {
	if t, ok := v.(time.Time); ok {
		f.appendUpdate(newUpdate(f, NewExpr("?", t.In(time.Local).Format("2006-01-02"))))
		return f
	}
	f.appendUpdate(newUpdate(f, NewExpr("?", v)))
	return f
}

func (f dateValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%q", f.val.In(time.Local).Format("2006-01-02"))), nil
}

func (f *dateValueField) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		return nil
	}
	if t, err := time.ParseInLocation("2006-01-02", strings.Trim(string(b), "\""), time.Local); err == nil {
		f.val = t
		f.unsetNull()
		return nil
	}
	it, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("dateValueField.UnmarshalJSON() failed to parse value (%s)", b), err)
	}
	f.val = time.Unix(it, 0)
	f.unsetNull()
	return nil
}

func (f *dateValueField) setByReq(req *http.Request) error {
	if err := req.ParseForm(); err != nil {
		return newErr(ErrCodeIO, "dateValueField.setByReq() failed to parse request body", err)
	}
	m := map[string][]string(req.Form)
	key := f.formName()
	if len(m[key]) > 0 && m[key][0] != "" {
		if t, err := time.ParseInLocation("2006-01-02", m[key][0], time.Local); err == nil {
			f.SetDate(t)
			return nil
		}
		it, err := strconv.ParseInt(m[key][0], 10, 64)
		if err != nil {
			return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("dateValueField.setByReq() failed to parse value (%s)", m[key][0]), err)
		}
		f.SetDate(time.Unix(it, 0))
	}
	return nil
}

func (f *dateValueField) setByStr(str string) error {
	if t, err := time.ParseInLocation("2006-01-02", str, time.Local); err == nil {
		f.SetDate(t)
		return nil
	}
	it, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("dateValueField.setByStr() failed to parse value (%s)", str), err)
	}
	f.SetDate(time.Unix(it, 0))
	return nil
}

func (f *dateValueField) AnyValue() time.Time {
	return f.val
}

type datetimeValueField struct {
	baseField
	val time.Time
	exp *Expr
}

func (f *datetimeValueField) init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.baseField.init(model, colName, fieldName, formName, uriName, index)
}

func (f *datetimeValueField) String() string {
	if !f.IsValid() || f.IsNull() {
		return fmt.Sprintf("%s: %v", f.fieldName(), nil)
	}
	return fmt.Sprintf("%s: %s", f.fieldName(), f.AnyValue().In(time.Local).Format("2006-01-02 15:04:05"))
}

func (f *datetimeValueField) toScan(m Model, selectors *[]interface{}) {
	fieldInfos := m.FieldInfos()
	for _, info := range fieldInfos {
		if info.Field.fieldName() == f.fieldName() {
			*selectors = append(*selectors, info.Field)
			return
		}
	}
	panic(fmt.Sprintf("cannot find field(%s) in %T", f.fieldName(), m))
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
		t, err := time.ParseInLocation("2006-01-02 15:04:05", string(val), time.Local)
		if err != nil {
			return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("datetimeValueField.Scan() failed to parse string value (value: %s)", val), err)
		}
		f.val = t
	case string:
		t, err := time.ParseInLocation("2006-01-02 15:04:05", val, time.Local)
		if err != nil {
			return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("datetimeValueField.Scan() failed to parse string value (value: %s)", val), err)
		}
		f.val = t
	case time.Time:
		f.val = val
	default:
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("datetimeValueField.Scan() invalid value type (%T)", v), nil)
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

func (f *datetimeValueField) value() *Expr {
	f.mustValid()
	if f.IsNull() {
		return NewExpr("NULL")
	}
	if f.exp != nil {
		return f.exp
	}
	return NewExpr("?", f.val.In(time.Local).Format("2006-01-02 15:04:05"))
}

func (f *datetimeValueField) sqlLiteral() string {
	f.mustValid()
	if f.IsNull() {
		return "NULL"
	}
	if f.exp != nil {
		return f.exp.exp
	}
	return fmt.Sprintf(`"%s"`, f.val.Format("2006-01-02 15:04:05"))
}

func (f *datetimeValueField) getExpr() *Expr {
	return f.exp
}

func (f *datetimeValueField) SetExpr(exp *Expr) *datetimeValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = time.Time{}, exp
	return f
}

func (f *datetimeValueField) SetDatetime(v time.Time) *datetimeValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = v, nil
	return f
}

func (f *datetimeValueField) set(v interface{}) ValueField {
	if exp, ok := v.(*Expr); ok {
		return f.SetExpr(exp)
	}
	return f.SetDatetime(v.(time.Time))
}

func (f *datetimeValueField) update(v interface{}) ValueField {
	if t, ok := v.(time.Time); ok {
		f.appendUpdate(newUpdate(f, NewExpr("?", t.In(time.Local).Format("2006-01-02 15:04:05"))))
		return f
	}
	f.appendUpdate(newUpdate(f, NewExpr("?", v)))
	return f
}

func (f datetimeValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%q", f.val.In(time.Local).Format("2006-01-02 15:04:05"))), nil
}

func (f *datetimeValueField) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		return nil
	}
	if t, err := time.ParseInLocation("2006-01-02 15:04:05", strings.Trim(string(b), "\""), time.Local); err == nil {
		f.unsetNull()
		f.val = t
		return nil
	}
	ti, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("datetimeValueField.UnmarshalJSON() invalid value format (%s)", b), err)
	}
	f.unsetNull()
	f.val = time.Unix(ti, 0)
	return nil
}

func (f *datetimeValueField) setByReq(req *http.Request) error {
	if err := req.ParseForm(); err != nil {
		return newErr(ErrCodeIO, "datetimeValueField.setByReq() failed to parse request form", err)
	}
	m := map[string][]string(req.Form)
	key := f.formName()
	if len(m[key]) > 0 && m[key][0] != "" {
		if t, err := time.ParseInLocation("2006-01-02 15:04:05", m[key][0], time.Local); err == nil {
			f.SetDatetime(t)
			return nil
		}
		ti, err := strconv.ParseInt(m[key][0], 10, 64)
		if err != nil {
			return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("datetimeValueField.setByReq() invalid value format (%s)", m[key][0]), err)
		}
		f.SetDatetime(time.Unix(ti, 0))
	}
	return nil
}

func (f *datetimeValueField) setByStr(str string) error {
	if t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local); err == nil {
		f.SetDatetime(t)
		return nil
	}
	ti, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("datetimeValueField.setByStr() invalid value format (%s)", str), err)
	}
	f.SetDatetime(time.Unix(ti, 0))
	return nil
}

func (f *datetimeValueField) AnyValue() time.Time {
	return f.val
}

type timeValueField struct {
	baseField
	val time.Time
	exp *Expr
}

func (f *timeValueField) init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.baseField.init(model, colName, fieldName, formName, uriName, index)
}

func (f *timeValueField) String() string {
	if !f.IsValid() || f.IsNull() {
		return fmt.Sprintf("%s: %v", f.fieldName(), nil)
	}
	return fmt.Sprintf("%s: %s", f.fieldName(), f.AnyValue().In(time.Local).Format("15:04:05"))
}

func (f *timeValueField) toScan(m Model, selectors *[]interface{}) {
	fieldInfos := m.FieldInfos()
	for _, info := range fieldInfos {
		if info.Field.fieldName() == f.fieldName() {
			*selectors = append(*selectors, info.Field)
			return
		}
	}
	panic(fmt.Sprintf("cannot find field(%s) in %T", f.fieldName(), m))
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
		t, err := time.ParseInLocation("15:04:05", string(val), time.Local)
		if err != nil {
			return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("timeValueField.Scan() invalid value format (value: %s)", val), err)
		}
		f.val = t
	case string:
		t, err := time.ParseInLocation("15:04:05", val, time.Local)
		if err != nil {
			return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("timeValueField.Scan() invalid value format (value: %s)", val), err)
		}
		f.val = t
	case time.Time:
		f.val = val
	default:
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("timeValueField.Scan() invalid value type (value type: %T)", v), nil)
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

func (f *timeValueField) value() *Expr {
	f.mustValid()
	if f.IsNull() {
		return NewExpr("NULL")
	}
	if f.exp != nil {
		return f.exp
	}
	return NewExpr("?", f.val.In(time.Local).Format("15:04:05"))
}

func (f *timeValueField) sqlLiteral() string {
	f.mustValid()
	if f.IsNull() {
		return "NULL"
	}
	if f.exp != nil {
		return f.exp.exp
	}
	return fmt.Sprintf(`"%s"`, f.val.Format("15:04:05"))
}

func (f *timeValueField) AnyValue() time.Time {
	return f.val
}

func (f *timeValueField) getExpr() *Expr {
	return f.exp
}

func (f *timeValueField) SetExpr(exp *Expr) *timeValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = time.Time{}, exp
	return f
}

func (f *timeValueField) SetTime(v time.Time) *timeValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = v, nil
	return f
}

func (f *timeValueField) set(v interface{}) ValueField {
	if exp, ok := v.(*Expr); ok {
		return f.SetExpr(exp)
	}
	return f.SetTime(v.(time.Time))
}

func (f *timeValueField) update(v interface{}) ValueField {
	if t, ok := v.(time.Time); ok {
		f.appendUpdate(newUpdate(f, NewExpr("?", t.In(time.Local).Format("2006-01-02 15:04:05"))))
		return f
	}
	f.appendUpdate(newUpdate(f, NewExpr("?", v)))
	return f
}

func (f timeValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(f.val.In(time.Local).Format("15:04:05")), nil
}

func (f *timeValueField) UnmarshalJSON(b []byte) error {
	f.setValid()
	if string(b) == "null" {
		f.SetNull()
		return nil
	}
	if t, err := time.ParseInLocation("15:04:05", strings.Trim(string(b), "\""), time.Local); err == nil {
		f.unsetNull()
		f.val = t
		return nil
	}
	ti, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("timeValueField.UnmarshalJSON() invalid value format (value: %s)", b), err)
	}
	f.val = time.Unix(ti, 0)
	f.unsetNull()
	return nil
}

func (f *timeValueField) setByReq(req *http.Request) error {
	if err := req.ParseForm(); err != nil {
		return newErr(ErrCodeIO, "timeValueField.setByReq() failed to parse request form", err)
	}
	m := map[string][]string(req.Form)
	key := f.formName()
	if len(m[key]) > 0 && m[key][0] != "" {
		if t, err := time.ParseInLocation("15:04:05", m[key][0], time.Local); err == nil {
			f.SetTime(t)
			return nil
		}
		ti, err := strconv.ParseInt(m[key][0], 10, 64)
		if err != nil {
			return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("timeValueField.setByReq() invalid value format (value: %s)", m[key][0]), err)
		}
		f.SetTime(time.Unix(ti, 0))
	}
	return nil
}

func (f *timeValueField) setByStr(str string) error {
	if t, err := time.ParseInLocation("15:04:05", str, time.Local); err == nil {
		f.SetTime(t)
		return nil
	}
	ti, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("timeValueField.setByStr() invalid value format (value: %s)", str), err)
	}
	f.SetTime(time.Unix(ti, 0))
	return nil
}

type decimalValueField struct {
	baseField
	val float64
	exp *Expr
}

func (f *decimalValueField) init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.baseField.init(model, colName, fieldName, formName, uriName, index)
}

func (f *decimalValueField) String() string {
	if !f.IsValid() || f.IsNull() {
		return fmt.Sprintf("%s: %v", f.fieldName(), nil)
	}
	return fmt.Sprintf("%s: %f", f.fieldName(), f.AnyValue())
}

func (f *decimalValueField) toScan(m Model, selectors *[]interface{}) {
	fieldInfos := m.FieldInfos()
	for _, info := range fieldInfos {
		if info.Field.fieldName() == f.fieldName() {
			*selectors = append(*selectors, info.Field)
			return
		}
	}
	panic(fmt.Sprintf("cannot find field(%s) in %T", f.fieldName(), m))
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
			return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("decimalValueField.Scan() invalid value format (value: %s)", val), err)
		}
		f.val = fv
	default:
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("invalid type for scan Decimal(%T)", v), nil)
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

func (f *decimalValueField) value() *Expr {
	f.mustValid()
	if f.IsNull() {
		return NewExpr("NULL")
	}
	if f.exp != nil {
		return f.exp
	}
	return NewExpr("?", f.val)
}

func (f *decimalValueField) sqlLiteral() string {
	f.mustValid()
	if f.IsNull() {
		return "NULL"
	}
	if f.exp != nil {
		return f.exp.exp
	}
	return fmt.Sprintf(`%f`, f.val)
}

func (f *decimalValueField) getExpr() *Expr {
	return f.exp
}

func (f *decimalValueField) SetExpr(exp *Expr) *decimalValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = 0, exp
	return f
}

func (f *decimalValueField) SetDecimal(v float64) *decimalValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = v, nil
	return f
}

func (f *decimalValueField) set(v interface{}) ValueField {
	if exp, ok := v.(*Expr); ok {
		return f.SetExpr(exp)
	}
	return f.SetDecimal(v.(float64))
}

func (f *decimalValueField) update(v interface{}) ValueField {
	f.appendUpdate(newUpdate(f, NewExpr("?", v)))
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
		f.SetNull()
		// f.removeStatus(notNull)
		return nil
	}
	d, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("decimalValueField.UnmarshalJSON() invalid value format (value: %s)", b), err)
	}
	// f.SetDecimal(d)
	// f.addStatus(notNull)
	f.unsetNull()
	f.val = d
	return nil
}

func (f *decimalValueField) setByReq(req *http.Request) error {
	if err := req.ParseForm(); err != nil {
		return newErr(ErrCodeIO, "decimalValueField.setByReq() failed to parser request form", err)
	}
	m := map[string][]string(req.Form)
	key := f.formName()
	if len(m[key]) > 0 {
		val, err := strconv.ParseFloat(m[key][0], 64)
		if err != nil {
			return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("decimalValueField.setByReq() invalid value format (value: %s)", m[key][0]), err)
		}
		f.SetDecimal(val)
	}
	return nil
}

func (f *decimalValueField) setByStr(str string) error {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return newErr(ErrCodeInvalidValueFormat, fmt.Sprintf("decimalValueField.setByStr() invalid value format (value: %s)", str), err)
	}
	f.SetDecimal(val)
	return nil
}

func (f *decimalValueField) AnyValue() float64 {
	return f.val
}

type byteValueField struct {
	baseField
	val []byte
	exp *Expr
}

func (f *byteValueField) init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.baseField.init(model, colName, fieldName, formName, uriName, index)
}

func (f *byteValueField) String() string {
	if !f.IsValid() || f.IsNull() {
		return fmt.Sprintf("%s: %v", f.fieldName(), nil)
	}
	return fmt.Sprintf("%s: %x", f.fieldName(), f.AnyValue())
}

func (f *byteValueField) toScan(m Model, selectors *[]interface{}) {
	fieldInfos := m.FieldInfos()
	for _, info := range fieldInfos {
		if info.Field.fieldName() == f.fieldName() {
			*selectors = append(*selectors, info.Field)
			return
		}
	}
	panic(fmt.Sprintf("cannot find field(%s) in %T", f.fieldName(), m))
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
			return newErr(ErrCodeInvalidValueFormat, "byteValueField.Scan() invalid value format", err)
		}
		f.val = bs
	default:
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("invalid type for scan Decimal(%T)", v), nil)
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

func (f *byteValueField) value() *Expr {
	f.mustValid()
	if f.IsNull() {
		return NewExpr("NULL")
	}
	if f.exp != nil {
		return f.exp
	}
	return NewExpr("?", fmt.Sprintf("x'%x'", f.val))
}

func (f *byteValueField) sqlLiteral() string {
	f.mustValid()
	if f.IsNull() {
		return "NULL"
	}
	if f.exp != nil {
		return f.exp.exp
	}
	return fmt.Sprintf(`x"%x"`, f.val)
}

func (f *byteValueField) AnyValue() []byte {
	return f.val
}

func (f *byteValueField) getExpr() *Expr {
	return f.exp
}

func (f *byteValueField) SetExpr(exp *Expr) *byteValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = nil, exp
	return f
}

func (f *byteValueField) SetBytes(v []byte) *byteValueField {
	f.setValid()
	f.unsetNull()
	f.val, f.exp = v, nil
	return f
}

func (f *byteValueField) set(v interface{}) ValueField {
	if exp, ok := v.(*Expr); ok {
		return f.SetExpr(exp)
	}
	return f.SetBytes(v.([]byte))
}

func (f *byteValueField) update(v interface{}) ValueField {
	f.appendUpdate(newUpdate(f, NewExpr("?", v)))
	return f
}

func (f byteValueField) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%x", f.val)), nil
}

func (f *byteValueField) setByReq(req *http.Request) error {
	if err := req.ParseForm(); err != nil {
		return newErr(ErrCodeIO, "byteValueField.setByReq() failed to parse request form", err)
	}
	m := map[string][]string(req.Form)
	key := f.formName()
	if len(m[key]) > 0 {
		val, err := hex.DecodeString(m[key][0])
		if err != nil {
			return newErr(ErrCodeInvalidValueFormat, "byteValueField.setByReq() invalid value format", err)
		}
		f.SetBytes(val)
	}
	return nil
}

func (f *byteValueField) setByStr(str string) error {
	val, err := hex.DecodeString(str)
	if err != nil {
		return newErr(ErrCodeInvalidValueFormat, "byteValueField.setByStr() invalid value format", err)
	}
	f.SetBytes(val)
	return nil
}
