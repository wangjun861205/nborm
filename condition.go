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

// AndCondition AndCondition
func AndCondition(exprStr string, values ...interface{}) *condition {
	return &condition{
		and,
		NewExpr(exprStr, values...),
	}
}

// OrCondition OrCondition
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

// AndEq AndEq
func AndEq(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ = ?", f.valueField(), value))
	}
}

// OrEq OrEq
func OrEq(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ = ?", f.valueField(), value))
	}
}

// AndNeq AndNeq
func AndNeq(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ <> ?", f.valueField(), value))
	}
}

// OrNeq OrNeq
func OrNeq(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ <> ?", f.valueField(), value))
	}
}

// AndLt AndLt
func AndLt(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ < ?", f.valueField(), value))
	}
}

// OrLt OrLt
func OrLt(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ < ?", f.valueField(), value))
	}
}

// AndLte AndLte
func AndLte(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ <= ?", f.valueField(), value))
	}
}

// OrLte OrLte
func OrLte(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ <= ?", f.valueField(), value))
	}
}

// AndGt AndGt
func AndGt(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ > ?", f.valueField(), value))
	}
}

// OrGt OrGt
func OrGt(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ > ?", f.valueField(), value))
	}
}

// AndGte AndGte
func AndGte(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ >= ?", f.valueField(), value))
	}
}

// OrGte OrGte
func OrGte(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ >= ?", f.valueField(), value))
	}
}

// AndIsNull AndIsNull
func AndIsNull() func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ IS NULL", f))
	}
}

// OrIsNull OrIsNull
func OrIsNull() func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ IS NULL", f))
	}
}

// AndIsNotNull AndIsNotNull
func AndIsNotNull() func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ IS NOT NULL", f))
	}
}

// OrIsNotNull OrIsNotNull
func OrIsNotNull() func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ IS NOT NULL", f))
	}
}

// AndIn AndIn
func AndIn(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ IN ?", f.valueField(), value))
	}
}

// OrIn OrIn
func OrIn(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ NOT IN ?", f.valueField(), value))
	}
}

// AndLike AndLike
func AndLike(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ LIKE ?", f.valueField(), value))
	}
}

// OrLike OrLike
func OrLike(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ LIKE ?", f.valueField(), value))
	}
}

// AndNotLike AndNotLike
func AndNotLike(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ NOT LIKE ?", f.valueField(), value))
	}
}

// OrNotLike OrNotLike
func OrNotLike(value interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ NOT LIKE ?", f.valueField(), value))
	}
}

// AndBetween AndBetween
func AndBetween(startValue, endValue interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ BETWEEN ? AND ?", f.valueField(), startValue, endValue))
	}
}

// OrBetween OrBetween
func OrBetween(startValue, endValue interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ BETWEEN ? AND ?", f.valueField(), startValue, endValue))
	}
}

// AndNotBetween AndNotBetween
func AndNotBetween(startValue, endValue interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(and, NewExpr("@ NOT BETWEEN ? AND ?", f.valueField(), startValue, endValue))
	}
}

// OrNotBetween OrNotBetween
func OrNotBetween(startValue, endValue interface{}) func(*clauseField) *condition {
	return func(f *clauseField) *condition {
		return newCondition(or, NewExpr("@ NOT BETWEEN ? AND ?", f.valueField(), startValue, endValue))
	}
}
