package nborm

import (
	"fmt"
	"io"
)

// ClauseField ClauseField
type ClauseField interface {
	// AndW() ClauseField
	// AndWhere(string, interface{}) ClauseField
	// OrWhere(string, interface{}) ClauseField
	// U() ClauseField
	// Update(interface{}) ClauseField
	toInsert(inserts *insertList)
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
	valueField.ExprUpdate(valueField, NewExpr("?", valueField, valueField.value()))
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

func (f *clauseField) toInsert(inserts *insertList) {
	valueField := f.valueField()
	if valueField.checkFieldStatus(valid) {
		*inserts = append(*inserts, newInsert(valueField, NewExpr("?", valueField.value())))
	}
}

func (f *clauseField) Update(value interface{}) ClauseField {
	f.valueField().update(value)
	return f
}

func (f *clauseField) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(f.valueField().fullColName()))
	w.Write([]byte(" "))
}

func (f *clauseField) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(f.valueField().rawFullColName()))
	w.Write([]byte(" "))
}
