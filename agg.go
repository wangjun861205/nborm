package nborm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
)

type aggregator interface {
	referencer
	clauser
	getField() Field
	getExpr() *Expr
	getName() string
	MarshalJSON() ([]byte, error)
	GetField() Field
}

// IntAgg 整数汇总
type IntAgg struct {
	expr  *Expr
	name  string
	field *Int
}

func newIntAgg(expr *Expr, name string) *IntAgg {
	f := new(Int)
	f.Init(nil, name, "", "", "", -1)
	return &IntAgg{expr, name, f}
}

func (a *IntAgg) String() string {
	return fmt.Sprintf("%s: %s", a.name, a.field.String())
}

func (a *IntAgg) toScan(m Model, selectors *[]interface{}) {
	// if m.getConList() != nil {
	// 	na := newIntAgg(a.expr, a.name)
	// 	m.appendAgg(na)
	// 	*selectors = append(*selectors, na)
	// } else {
	// 	for _, agg := range m.getAggs() {
	// 		if agg.getName() == a.name {
	// 			*selectors = append(*selectors, agg)
	// 			return
	// 		}
	// 	}
	// 	panic(fmt.Sprintf("IntAgg.toScan() error: cannot find agg(name: %s)", a.name))
	// }
	for _, agg := range m.getAggs() {
		if agg.getName() == a.name {
			*selectors = append(*selectors, agg)
			return
		}
	}
	na := newStrAgg(a.expr, a.name)
	m.appendAgg(na)
	*selectors = append(*selectors, na)
}

// Scan 实现Scanner接口
func (a *IntAgg) Scan(v interface{}) error {
	return a.field.Scan(v)
}

func (a *IntAgg) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	a.expr.toClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte("AS "))
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *IntAgg) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	a.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte("AS "))
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *IntAgg) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *IntAgg) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *IntAgg) getField() Field {
	return a.field
}

func (a *IntAgg) getExpr() *Expr {
	return a.expr
}

func (a *IntAgg) getName() string {
	return a.name
}

// MarshalJSON 实现JSON接口
func (a *IntAgg) MarshalJSON() ([]byte, error) {
	return marshalAgg(a)
}

// AnyValue 直接返回数据
func (a *IntAgg) AnyValue() int {
	return a.field.AnyValue()
}

func (a *IntAgg) GetField() Field {
	return a.field
}

// StrAgg 字符串汇总
type StrAgg struct {
	expr  *Expr
	name  string
	field *String
}

func newStrAgg(expr *Expr, name string) *StrAgg {
	f := new(String)
	f.Init(nil, name, "", "", "", -1)
	return &StrAgg{expr, name, f}
}

func (a *StrAgg) String() string {
	return fmt.Sprintf("%s: %s", a.name, a.field.String())
}

func (a *StrAgg) toScan(m Model, selectors *[]interface{}) {
	// if m.getConList() != nil {
	// 	na := newStrAgg(a.expr, a.name)
	// 	m.appendAgg(na)
	// 	*selectors = append(*selectors, na)

	// } else {
	for _, agg := range m.getAggs() {
		if agg.getName() == a.name {
			*selectors = append(*selectors, agg)
			return
		}
	}
	na := newStrAgg(a.expr, a.name)
	m.appendAgg(na)
	*selectors = append(*selectors, na)
	// panic(fmt.Sprintf("StrAgg.toScan() error: cannot find agg(name: %s)", a.name))
	// }
}

// Scan 实现Scanner接口
func (a *StrAgg) Scan(v interface{}) error {
	return a.field.Scan(v)
}

func (a *StrAgg) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	a.expr.toClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte("AS "))
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *StrAgg) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	a.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte("AS "))
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *StrAgg) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *StrAgg) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *StrAgg) getField() Field {
	return a.field
}

func (a *StrAgg) getExpr() *Expr {
	return a.expr
}

func (a *StrAgg) getName() string {
	return a.name
}

// MarshalJSON 实现JSON接口
func (a *StrAgg) MarshalJSON() ([]byte, error) {
	return marshalAgg(a)
}

// AnyValue 直接返回数据
func (a *StrAgg) AnyValue() string {
	return a.field.AnyValue()
}

func (a *StrAgg) GetField() Field {
	return a.field
}

// DateAgg 日期汇总
type DateAgg struct {
	expr  *Expr
	name  string
	field *Date
}

func newDateAgg(expr *Expr, name string) *DateAgg {
	f := new(Date)
	f.Init(nil, name, "", "", "", -1)
	return &DateAgg{expr, name, f}
}

func (a *DateAgg) String() string {
	return fmt.Sprintf("%s: %s", a.name, a.field.String())
}

func (a *DateAgg) toScan(m Model, selectors *[]interface{}) {
	// if m.getConList() != nil {
	// 	na := newDateAgg(a.expr, a.name)
	// 	m.appendAgg(na)
	// 	*selectors = append(*selectors, na)
	// } else {
	// 	for _, agg := range m.getAggs() {
	// 		if agg.getName() == a.name {
	// 			*selectors = append(*selectors, agg)
	// 			return
	// 		}
	// 	}
	// 	panic(fmt.Sprintf("DateAgg.toScan() error: cannot find agg(name: %s)", a.name))
	// }
	for _, agg := range m.getAggs() {
		if agg.getName() == a.name {
			*selectors = append(*selectors, agg)
			return
		}
	}
	na := newStrAgg(a.expr, a.name)
	m.appendAgg(na)
	*selectors = append(*selectors, na)
}

// Scan 实现Scanner接口
func (a *DateAgg) Scan(v interface{}) error {
	return a.field.Scan(v)
}

func (a *DateAgg) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	a.expr.toClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte("AS "))
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DateAgg) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	a.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte("AS "))
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DateAgg) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DateAgg) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DateAgg) getField() Field {
	return a.field
}

func (a *DateAgg) getExpr() *Expr {
	return a.expr
}

func (a *DateAgg) getName() string {
	return a.name
}

// MarshalJSON 实现JSON接口
func (a *DateAgg) MarshalJSON() ([]byte, error) {
	return marshalAgg(a)
}

// AnyValue 直接返回数据
func (a *DateAgg) AnyValue() time.Time {
	return a.field.AnyValue()
}

func (a *DateAgg) GetField() Field {
	return a.field
}

// DatetimeAgg 日期时间汇总
type DatetimeAgg struct {
	expr  *Expr
	name  string
	field *Datetime
}

func newDatetimeAgg(expr *Expr, name string) *DatetimeAgg {
	f := new(Datetime)
	f.Init(nil, name, "", "", "", -1)
	return &DatetimeAgg{expr, name, f}
}

func (a *DatetimeAgg) String() string {
	return fmt.Sprintf("%s: %s", a.name, a.field.String())
}

func (a *DatetimeAgg) toScan(m Model, selectors *[]interface{}) {
	// if m.getConList() != nil {
	// 	na := newDatetimeAgg(a.expr, a.name)
	// 	m.appendAgg(na)
	// 	*selectors = append(*selectors, na)
	// } else {
	// 	for _, agg := range m.getAggs() {
	// 		if agg.getName() == a.name {
	// 			*selectors = append(*selectors, agg)
	// 			return
	// 		}
	// 	}
	// 	panic(fmt.Sprintf("DateTimeAgg.toScan() error: cannot find agg(name: %s)", a.name))
	// }
	for _, agg := range m.getAggs() {
		if agg.getName() == a.name {
			*selectors = append(*selectors, agg)
			return
		}
	}
	na := newStrAgg(a.expr, a.name)
	m.appendAgg(na)
	*selectors = append(*selectors, na)
}

// Scan 实现Scanner接口
func (a *DatetimeAgg) Scan(v interface{}) error {
	return a.field.Scan(v)
}

func (a *DatetimeAgg) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	a.expr.toClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte("AS "))
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DatetimeAgg) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	a.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte("AS "))
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DatetimeAgg) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DatetimeAgg) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DatetimeAgg) getField() Field {
	return a.field
}

func (a *DatetimeAgg) getExpr() *Expr {
	return a.expr
}

func (a *DatetimeAgg) getName() string {
	return a.name
}

// MarshalJSON 实现JSON接口
func (a *DatetimeAgg) MarshalJSON() ([]byte, error) {
	return marshalAgg(a)
}

// AnyValue 直接返回数据
func (a *DatetimeAgg) AnyValue() time.Time {
	return a.field.AnyValue()
}

func (a *DatetimeAgg) GetField() Field {
	return a.field
}

// TimeAgg 时间汇总
type TimeAgg struct {
	expr  *Expr
	name  string
	field *Time
}

func newTimeAgg(expr *Expr, name string) *TimeAgg {
	f := new(Time)
	f.Init(nil, name, "", "", "", -1)
	return &TimeAgg{expr, name, f}
}

func (a *TimeAgg) String() string {
	return fmt.Sprintf("%s: %s", a.name, a.field.String())
}

func (a *TimeAgg) toScan(m Model, selectors *[]interface{}) {
	// if m.getConList() != nil {
	// 	na := newTimeAgg(a.expr, a.name)
	// 	m.appendAgg(na)
	// 	*selectors = append(*selectors, na)
	// } else {
	// 	for _, agg := range m.getAggs() {
	// 		if agg.getName() == a.name {
	// 			*selectors = append(*selectors, agg)
	// 			return
	// 		}
	// 	}
	// 	panic(fmt.Sprintf("TimeAgg.toScan() error: cannot find agg(name: %s)", a.name))
	// }
	for _, agg := range m.getAggs() {
		if agg.getName() == a.name {
			*selectors = append(*selectors, agg)
			return
		}
	}
	na := newStrAgg(a.expr, a.name)
	m.appendAgg(na)
	*selectors = append(*selectors, na)
}

// Scan 实现Scanner接口
func (a *TimeAgg) Scan(v interface{}) error {
	return a.field.Scan(v)
}

func (a *TimeAgg) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	a.expr.toClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte("AS "))
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *TimeAgg) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	a.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte("AS "))
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *TimeAgg) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *TimeAgg) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *TimeAgg) getField() Field {
	return a.field
}

func (a *TimeAgg) getExpr() *Expr {
	return a.expr
}

func (a *TimeAgg) getName() string {
	return a.name
}

// MarshalJSON 实现JSON接口
func (a *TimeAgg) MarshalJSON() ([]byte, error) {
	return marshalAgg(a)
}

// AnyValue 直接返回数据
func (a *TimeAgg) AnyValue() time.Time {
	return a.field.AnyValue()
}

func (a *TimeAgg) GetField() Field {
	return a.field
}

// DecimalAgg 浮点数汇总
type DecimalAgg struct {
	expr  *Expr
	name  string
	field *Decimal
}

func newDecAgg(expr *Expr, name string) *DecimalAgg {
	f := new(Decimal)
	f.Init(nil, name, "", "", "", -1)
	return &DecimalAgg{expr, name, f}
}

func (a *DecimalAgg) String() string {
	return fmt.Sprintf("%s: %s", a.name, a.field.String())
}

func (a *DecimalAgg) toScan(m Model, selectors *[]interface{}) {
	// if m.getConList() != nil {
	// 	na := newDecAgg(a.expr, a.name)
	// 	m.appendAgg(na)
	// 	*selectors = append(*selectors, na)
	// } else {
	// 	for _, agg := range m.getAggs() {
	// 		if agg.getName() == a.name {
	// 			*selectors = append(*selectors, agg)
	// 			return
	// 		}
	// 	}
	// 	panic(fmt.Sprintf("DecimalAgg.toScan() error: cannot find agg(name: %s)", a.name))
	// }
	for _, agg := range m.getAggs() {
		if agg.getName() == a.name {
			*selectors = append(*selectors, agg)
			return
		}
	}
	na := newStrAgg(a.expr, a.name)
	m.appendAgg(na)
	*selectors = append(*selectors, na)
}

// Scan 实现Scanner接口
func (a *DecimalAgg) Scan(v interface{}) error {
	return a.field.Scan(v)
}

func (a *DecimalAgg) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	a.expr.toClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte("AS "))
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DecimalAgg) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	a.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte("AS "))
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DecimalAgg) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DecimalAgg) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DecimalAgg) getField() Field {
	return a.field
}

func (a *DecimalAgg) getExpr() *Expr {
	return a.expr
}

func (a *DecimalAgg) getName() string {
	return a.name
}

// MarshalJSON 实现JSON接口
func (a *DecimalAgg) MarshalJSON() ([]byte, error) {
	return marshalAgg(a)
}

// AnyValue 直接返回数据
func (a *DecimalAgg) AnyValue() float64 {
	return a.field.AnyValue()
}

func (a *DecimalAgg) GetField() Field {
	return a.field
}

func marshalAgg(agg aggregator) ([]byte, error) {
	if !agg.getField().IsValid() || agg.getField().IsNull() {
		return []byte(fmt.Sprintf("{\"%s\": null}", agg.getField().colName())), nil
	}
	buffer := bytes.NewBuffer(make([]byte, 0, 64))
	fmt.Fprintf(buffer, "{ \"%s\": ", agg.getField().colName())
	b, err := json.Marshal(agg.getField())
	if err != nil {
		return nil, err
	}
	buffer.Write(b)
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

type aggList []aggregator

func (l aggList) copy() aggList {
	nl := make(aggList, 0, 8)
	for _, agg := range l {
		switch a := agg.(type) {
		case *StrAgg:
			nl = append(nl, newStrAgg(a.expr, a.field.colName()))
		case *IntAgg:
			nl = append(nl, newIntAgg(a.expr, a.field.colName()))
		case *DateAgg:
			nl = append(nl, newDateAgg(a.expr, a.field.colName()))
		case *DatetimeAgg:
			nl = append(nl, newDatetimeAgg(a.expr, a.field.colName()))
		case *TimeAgg:
			nl = append(nl, newTimeAgg(a.expr, a.field.colName()))
		case *DecimalAgg:
			nl = append(nl, newDecAgg(a.expr, a.field.colName()))
		default:
			panic(fmt.Errorf("unsupported field type (%T)", agg))
		}

	}
	return nl
}

func (l aggList) forCheckDup() string {
	var builder strings.Builder
	for _, agg := range l {
		switch a := agg.(type) {
		case *StrAgg:
			builder.WriteString(fmt.Sprintf("%v", a.AnyValue()))
		case *IntAgg:
			builder.WriteString(fmt.Sprintf("%v", a.AnyValue()))
		case *DateAgg:
			builder.WriteString(fmt.Sprintf("%v", a.AnyValue()))
		case *DatetimeAgg:
			builder.WriteString(fmt.Sprintf("%v", a.AnyValue()))
		case *TimeAgg:
			builder.WriteString(fmt.Sprintf("%v", a.AnyValue()))
		case *DecimalAgg:
			builder.WriteString(fmt.Sprintf("%v", a.AnyValue()))
		default:
			panic(fmt.Errorf("unsupported field type (%T)", agg))
		}
	}
	return builder.String()
}

// 尝试将Aggs作为Model的Field来进行JSON Marshal
// func (l aggList) MarshalJSON() ([]byte, error) {
// 	buf := &bytes.Buffer{}
// 	if len(l) == 0 {
// 		return []byte("{}"), nil
// 	}
// 	for _, agg := range l {
// 		b, err := marshalAgg(agg)
// 		if err != nil {
// 			return nil, err
// 		}
// 		buf.Write(bytes.Trim(b, "{}"))
// 		buf.WriteString(",")
// 	}
// 	b := []byte(fmt.Sprintf("{%s}", bytes.Trim(buf.Bytes(), ",")))
// 	fmt.Println(string(b))
// 	return b, nil
// }
