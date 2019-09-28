package nborm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type aggregator interface {
	referencer
	clauser
	getField() Field
	getExpr() *Expr
	getName() string
	MarshalJSON() ([]byte, error)
}

// IntAgg 整数汇总
type IntAgg struct {
	expr  *Expr
	name  string
	field *Int
}

func newIntAgg(expr *Expr, name string) *IntAgg {
	f := new(Int)
	f.Init(nil, name, "", -1)
	return &IntAgg{expr, name, f}
}

func (a *IntAgg) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	a.expr.toClause(w, vals, isFirstGroup, isFirstNode)
}

func (a *IntAgg) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	a.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}

func (a *IntAgg) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *IntAgg) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
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

// StrAgg 字符串汇总
type StrAgg struct {
	expr  *Expr
	name  string
	field *String
}

func newStrAgg(expr *Expr, name string) *StrAgg {
	f := new(String)
	f.Init(nil, name, "", -1)
	return &StrAgg{expr, name, f}
}

func (a *StrAgg) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	a.expr.toClause(w, vals, isFirstGroup, isFirstNode)
}

func (a *StrAgg) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	a.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}

func (a *StrAgg) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *StrAgg) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
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

// DateAgg 日期汇总
type DateAgg struct {
	expr  *Expr
	name  string
	field *Date
}

func newDateAgg(expr *Expr, name string) *DateAgg {
	f := new(Date)
	f.Init(nil, name, "", -1)
	return &DateAgg{expr, name, f}
}

func (a *DateAgg) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	a.expr.toClause(w, vals, isFirstGroup, isFirstNode)
}

func (a *DateAgg) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	a.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}

func (a *DateAgg) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DateAgg) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
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

// DatetimeAgg 日期时间汇总
type DatetimeAgg struct {
	expr  *Expr
	name  string
	field *Datetime
}

func newDatetimeAgg(expr *Expr, name string) *DatetimeAgg {
	f := new(Datetime)
	f.Init(nil, name, "", -1)
	return &DatetimeAgg{expr, name, f}
}

func (a *DatetimeAgg) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	a.expr.toClause(w, vals, isFirstGroup, isFirstNode)
}

func (a *DatetimeAgg) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	a.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}

func (a *DatetimeAgg) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DatetimeAgg) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
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

// TimeAgg 时间汇总
type TimeAgg struct {
	expr  *Expr
	name  string
	field *Time
}

func newTimeAgg(expr *Expr, name string) *TimeAgg {
	f := new(Time)
	f.Init(nil, name, "", -1)
	return &TimeAgg{expr, name, f}
}

func (a *TimeAgg) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	a.expr.toClause(w, vals, isFirstGroup, isFirstNode)
}

func (a *TimeAgg) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	a.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}

func (a *TimeAgg) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *TimeAgg) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
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

// DecimalAgg 浮点数汇总
type DecimalAgg struct {
	expr  *Expr
	name  string
	field *Decimal
}

func newDecAgg(expr *Expr, name string) *DecimalAgg {
	f := new(Decimal)
	f.Init(nil, name, "", -1)
	return &DecimalAgg{expr, name, f}
}

func (a *DecimalAgg) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	a.expr.toClause(w, vals, isFirstGroup, isFirstNode)
}

func (a *DecimalAgg) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	a.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}

func (a *DecimalAgg) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	w.Write([]byte(a.name))
	w.Write([]byte(" "))
}

func (a *DecimalAgg) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
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
