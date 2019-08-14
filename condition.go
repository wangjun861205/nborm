package nborm

import (
	"fmt"
	"regexp"
	"strings"
)

type condRelType string

const (
	and condRelType = "AND"
	or  condRelType = "OR"
)

type condition struct {
	rel  condRelType
	expr *Expr
}

func newCondition(rel condRelType, expr *Expr) *condition {
	return &condition{
		rel,
		expr,
	}
}

func AndCondition(exprStr string, values ...interface{}) *condition {
	return &condition{
		and,
		NewExpr(exprStr, values...),
	}
}

func OrCondition(exprStr string, values ...interface{}) *condition {
	return &condition{
		or,
		NewExpr(exprStr, values...),
	}
}

func (w *condition) toExpr() *Expr {
	return &Expr{
		fmt.Sprintf("%s %s", w.rel, w.expr.exp),
		w.expr.values,
	}
}

type conditionList []*condition

func (l conditionList) toExpr() *Expr {
	if len(l) == 0 {
		return nil
	}
	var builder strings.Builder
	values := make([]interface{}, 0, 16)
	for _, c := range l {
		exp := c.toExpr()
		builder.WriteString(exp.exp)
		values = append(values, exp.values...)
	}
	return &Expr{
		builder.String(),
		values,
	}
}

func (l conditionList) group(rel condRelType) *condition {
	if len(l) == 0 {
		return nil
	}
	var builder strings.Builder
	values := make([]interface{}, 0, 16)
	for _, c := range l {
		exp := c.toExpr()
		builder.WriteString(exp.exp)
		builder.WriteString(" ")
		values = append(values, exp.values...)
	}
	return newCondition(rel, NewExpr(fmt.Sprintf("(%s)", regexp.MustCompile(`^(AND|OR) `).ReplaceAllString(builder.String(), "")), values...))
}

func AndEq(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ = ?", f.valueField(), value))
	}
}

func OrEq(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ = ?", f.valueField(), value))
	}
}

func AndNeq(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ <> ?", f.valueField(), value))
	}
}

func OrNeq(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ <> ?", f.valueField(), value))
	}
}

func AndLt(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ < ?", f.valueField(), value))
	}
}

func OrLt(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ < ?", f.valueField(), value))
	}
}

func AndLte(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ <= ?", f.valueField(), value))
	}
}

func OrLte(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ <= ?", f.valueField(), value))
	}
}

func AndGt(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ > ?", f.valueField(), value))
	}
}

func OrGt(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ > ?", f.valueField(), value))
	}
}

func AndGte(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ >= ?", f.valueField(), value))
	}
}

func OrGte(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ >= ?", f.valueField(), value))
	}
}

func AndIsNull() func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ IS NULL", f))
	}
}

func OrIsNull() func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ IS NULL", f))
	}
}

func AndIsNotNull() func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ IS NOT NULL", f))
	}
}

func OrIsNotNull() func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ IS NOT NULL", f))
	}
}

func AndIn(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ IN ?", f.valueField(), value))
	}
}

func OrIn(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ NOT IN ?", f.valueField(), value))
	}
}

func AndLike(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ LIKE ?", f.valueField(), value))
	}
}

func OrLike(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ LIKE ?", f.valueField(), value))
	}
}

func AndNotLike(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ NOT LIKE ?", f.valueField(), value))
	}
}

func OrNotLike(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ NOT LIKE ?", f.valueField(), value))
	}
}

func AndBetween(startValue, endValue interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ BETWEEN ? AND ?", f.valueField(), startValue, endValue))
	}
}

func OrBetween(startValue, endValue interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ BETWEEN ? AND ?", f.valueField(), startValue, endValue))
	}
}

func AndNotBetween(startValue, endValue interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ NOT BETWEEN ? AND ?", f.valueField(), startValue, endValue))
	}
}

func OrNotBetween(startValue, endValue interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ NOT BETWEEN ? AND ?", f.valueField(), startValue, endValue))
	}
}
