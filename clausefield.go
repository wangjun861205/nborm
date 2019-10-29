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
	U() ClauseField
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
	valueField.ExprUpdate(valueField, NewExpr("?", valueField.value()))
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

func (f *clauseField) AndEq(value interface{}) ClauseField {
	f.AndWhere("=", value)
	return f
}

func (f *clauseField) OrEq(value interface{}) ClauseField {
	f.OrWhere("=", value)
	return f
}

func (f *clauseField) AndNeq(value interface{}) ClauseField {
	f.AndWhere("<>", value)
	return f
}

func (f *clauseField) OrNeq(value interface{}) ClauseField {
	f.OrWhere("<>", value)
	return f
}

func (f *clauseField) AndLt(value interface{}) ClauseField {
	f.AndWhere("<", value)
	return f
}

func (f *clauseField) OrLt(value interface{}) ClauseField {
	f.OrWhere("<", value)
	return f
}

func (f *clauseField) AndLte(value interface{}) ClauseField {
	f.AndWhere("<=", value)
	return f
}

func (f *clauseField) OrLte(value interface{}) ClauseField {
	f.OrWhere("<=", value)
	return f
}

func (f *clauseField) AndGt(value interface{}) ClauseField {
	f.AndWhere(">", value)
	return f
}

func (f *clauseField) OrGt(value interface{}) ClauseField {
	f.OrWhere(">", value)
	return f
}

func (f *clauseField) AndGte(value interface{}) ClauseField {
	f.AndWhere(">=", value)
	return f
}

func (f *clauseField) OrGte(value interface{}) ClauseField {
	f.OrWhere(">=", value)
	return f
}

func (f *clauseField) AndIsNull() ClauseField {
	f.AndWhere("IS NULL", nil)
	return f
}

func (f *clauseField) OrIsNull() ClauseField {
	f.OrWhere("IS NULL", nil)
	return f
}

func (f *clauseField) AndIsNotNull() ClauseField {
	f.AndWhere("IS NOT NULL", nil)
	return f
}

func (f *clauseField) OrIsNotNull() ClauseField {
	f.OrWhere("IS NOT NULL", nil)
	return f
}

func (f *clauseField) AndIn(value interface{}) ClauseField {
	f.AndWhere("IN", value)
	return f
}

func (f *clauseField) OrIn(value interface{}) ClauseField {
	f.OrWhere("IN", value)
	return f
}

func (f *clauseField) AndLike(value interface{}) ClauseField {
	f.AndWhere("LIKE", value)
	return f
}

func (f *clauseField) OrLike(value interface{}) ClauseField {
	f.OrWhere("LIKE", value)
	return f
}

func (f *clauseField) AndNotLike(value interface{}) ClauseField {
	f.AndWhere("NOT LIKE", value)
	return f
}

func (f *clauseField) OrNotLike(value interface{}) ClauseField {
	f.OrWhere("NOT LIKE", value)
	return f
}

func (f *clauseField) AndBetween(startValue, endValue interface{}) ClauseField {
	f.valueField().AndExprWhere(NewExpr("@ BETWEEN ? AND ?", startValue, endValue))
	return f
}

func (f *clauseField) OrBetween(startValue, endValue interface{}) ClauseField {
	f.valueField().OrExprWhere(NewExpr("@ BETWEEN ? AND ?", f.valueField(), startValue, endValue))
	return f
}

func (f *clauseField) AndNotBetween(startValue, endValue interface{}) ClauseField {
	f.valueField().AndExprWhere(NewExpr("@ NOT BETWEEN ? AND ?", f.valueField(), startValue, endValue))
	return f
}

func (f *clauseField) OrNotBetween(startValue, endValue interface{}) ClauseField {
	f.valueField().OrExprWhere(NewExpr("@ NOT BETWEEN ? AND ?", f.valueField(), startValue, endValue))
	return f
}

func (f *clauseField) AndEqWhere(value interface{}) *where {
	return newWhere(NewExpr("@ = ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrEqWhere(value interface{}) *where {
	return newWhere(NewExpr("@ = ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndNeqWhere(value interface{}) *where {
	return newWhere(NewExpr("@ <> ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrNeqWhere(value interface{}) *where {
	return newWhere(NewExpr("@ <> ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndLtWhere(value interface{}) *where {
	return newWhere(NewExpr("@ < ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrLtWhere(value interface{}) *where {
	return newWhere(NewExpr("@ < ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndLteWhere(value interface{}) *where {
	return newWhere(NewExpr("@ <= ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrLteWhere(value interface{}) *where {
	return newWhere(NewExpr("@ <= ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndGtWhere(value interface{}) *where {
	return newWhere(NewExpr("@ > ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrGtWhere(value interface{}) *where {
	return newWhere(NewExpr("@ > ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndGteWhere(value interface{}) *where {
	return newWhere(NewExpr("@ >= ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrGteWhere(value interface{}) *where {
	return newWhere(NewExpr("@ >= ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndIsNullWhere() *where {
	return newWhere(NewExpr("@ IS NULL", f), whereAnd)
}

func (f *clauseField) OrIsNullWhere() *where {
	return newWhere(NewExpr("@ IS NULL", f), whereOr)
}

func (f *clauseField) AndIsNotNullWhere() *where {
	return newWhere(NewExpr("@ IS NOT NULL", f), whereAnd)
}

func (f *clauseField) OrIsNotNullWhere() *where {
	return newWhere(NewExpr("@ IS NOT NULL", f), whereOr)
}

func (f *clauseField) AndInWhere(value interface{}) *where {
	return newWhere(NewExpr("@ IN ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrInWhere(value interface{}) *where {
	return newWhere(NewExpr("@ NOT IN ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndLikeWhere(value interface{}) *where {
	return newWhere(NewExpr("@ LIKE ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrLikeWhere(value interface{}) *where {
	return newWhere(NewExpr("@ LIKE ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndNotLikeWhere(value interface{}) *where {
	return newWhere(NewExpr("@ NOT LIKE ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrNotLikeWhere(value interface{}) *where {
	return newWhere(NewExpr("@ NOT LIKE ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndBetweenWhere(startValue, endValue interface{}) *where {
	return newWhere(NewExpr("@ BETWEEN ? AND ?", f.valueField(), startValue, endValue), whereAnd)
}

func (f *clauseField) OrBetweenWhere(startValue, endValue interface{}) *where {
	return newWhere(NewExpr("@ BETWEEN ? AND ?", f.valueField(), startValue, endValue), whereOr)
}

func (f *clauseField) AndNotBetweenWhere(startValue, endValue interface{}) *where {
	return newWhere(NewExpr("@ NOT BETWEEN ? AND ?", f.valueField(), startValue, endValue), whereAnd)
}

func (f *clauseField) OrNotBetweenWhere(startValue, endValue interface{}) *where {
	return newWhere(NewExpr("@ NOT BETWEEN ? AND ?", f.valueField(), startValue, endValue), whereOr)
}

func (f *clauseField) AndEqHaving(value interface{}) *having {
	return newHaving(NewExpr("@ = ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrEqHaving(value interface{}) *having {
	return newHaving(NewExpr("@ = ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndNeqHaving(value interface{}) *having {
	return newHaving(NewExpr("@ <> ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrNeqHaving(value interface{}) *having {
	return newHaving(NewExpr("@ <> ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndLtHaving(value interface{}) *having {
	return newHaving(NewExpr("@ < ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrLtHaving(value interface{}) *having {
	return newHaving(NewExpr("@ < ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndLteHaving(value interface{}) *having {
	return newHaving(NewExpr("@ <= ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrLteHaving(value interface{}) *having {
	return newHaving(NewExpr("@ <= ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndGtHaving(value interface{}) *having {
	return newHaving(NewExpr("@ > ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrGtHaving(value interface{}) *having {
	return newHaving(NewExpr("@ > ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndGteHaving(value interface{}) *having {
	return newHaving(NewExpr("@ >= ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrGteHaving(value interface{}) *having {
	return newHaving(NewExpr("@ >= ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndIsNullHaving() *having {
	return newHaving(NewExpr("@ IS NULL", f), whereAnd)
}

func (f *clauseField) OrIsNullHaving() *having {
	return newHaving(NewExpr("@ IS NULL", f), whereOr)
}

func (f *clauseField) AndIsNotNullHaving() *having {
	return newHaving(NewExpr("@ IS NOT NULL", f), whereAnd)
}

func (f *clauseField) OrIsNotNullHaving() *having {
	return newHaving(NewExpr("@ IS NOT NULL", f), whereOr)
}

func (f *clauseField) AndInHaving(value interface{}) *having {
	return newHaving(NewExpr("@ IN ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrInHaving(value interface{}) *having {
	return newHaving(NewExpr("@ NOT IN ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndLikeHaving(value interface{}) *having {
	return newHaving(NewExpr("@ LIKE ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrLikeHaving(value interface{}) *having {
	return newHaving(NewExpr("@ LIKE ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndNotLikeHaving(value interface{}) *having {
	return newHaving(NewExpr("@ NOT LIKE ?", f.valueField(), value), whereAnd)
}

func (f *clauseField) OrNotLikeHaving(value interface{}) *having {
	return newHaving(NewExpr("@ NOT LIKE ?", f.valueField(), value), whereOr)
}

func (f *clauseField) AndBetweenHaving(startValue, endValue interface{}) *having {
	return newHaving(NewExpr("@ BETWEEN ? AND ?", f.valueField(), startValue, endValue), whereAnd)
}

func (f *clauseField) OrBetweenHaving(startValue, endValue interface{}) *having {
	return newHaving(NewExpr("@ BETWEEN ? AND ?", f.valueField(), startValue, endValue), whereOr)
}

func (f *clauseField) AndNotBetweenHaving(startValue, endValue interface{}) *having {
	return newHaving(NewExpr("@ NOT BETWEEN ? AND ?", f.valueField(), startValue, endValue), whereAnd)
}

func (f *clauseField) OrNotBetweenHaving(startValue, endValue interface{}) *having {
	return newHaving(NewExpr("@ NOT BETWEEN ? AND ?", f.valueField(), startValue, endValue), whereOr)
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
