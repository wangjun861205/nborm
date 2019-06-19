package nborm

import (
	"fmt"
	"strings"
	"time"
)

type clauseStatus int

const (
	byValue clauseStatus = iota
	byExpr
)

type updateSet struct {
	field  Field
	value  interface{}
	expr   *Expr
	status clauseStatus
}

func newUpdateSet(field Field, value interface{}) *updateSet {
	set := &updateSet{field: field}
	if v, ok := value.(*Expr); ok {
		set.expr = v
		set.status = byExpr
	} else {
		set.value = value
		set.status = byValue
	}
	return set
}

func (u *updateSet) toClause(cl *[]string, vl *[]interface{}) {
	if u.status == byValue {
		*cl = append(*cl, fmt.Sprintf("%s.%s.%s = ?"), u.field.dbName(), u.field.tabName(), u.field.fieldName())
		*vl = append(*vl, u.value)
		return
	}
	*cl = append(*cl, fmt.Sprintf("%s.%s.%s = %s", u.field.dbName(), u.field.tabName(), u.field.fieldName()))
}

type updateSetList []*updateSet

func (l updateSetList) toClause() (string, []interface{}) {
	cl := make([]string, 0, len(l))
	vl := make([]interface{}, 0, len(l))
	for _, u := range l {
		u.toClause(&cl, &vl)
	}
	return fmt.Sprintf("SET %s", strings.Join(cl, ",")), vl
}

type operator string

const (
	eq   operator = "="
	neq  operator = "<>"
	lt   operator = "<"
	lte  operator = "<="
	gt   operator = ">"
	gte  operator = ">="
	is   operator = "IS"
	nis  operator = "IS NOT"
	like operator = "LIKE"
	in   operator = "IN"
)

type rel string

const (
	and rel = "AND"
	or  rel = "OR"
)

func checkOp(op string) {
	switch op {
	case "=", "<>", "<", "<=", ">", ">=", "IS", "IS NOT", "LIKE", "IN":
		return
	default:
		panic(fmt.Sprintf("unsupported operator (%s)", op))
	}
}

type where struct {
	rel    rel
	field  Field
	op     operator
	value  interface{}
	expr   *Expr
	status clauseStatus
}

func newWhere(rel rel, field Field, op string, value interface{}) *where {
	checkOp(op)
	w := &where{rel: rel, field: field, op: operator(op)}
	if v, ok := value.(*Expr); ok {
		w.expr = v
		w.status = byExpr
		return w
	}
	w.value = value
	w.status = byValue
	return w
}

func (w *where) toClause(cl *[]string, vl *[]interface{}) {
	switch w.status {
	case byValue:
		if w.op == in {
			switch l := w.value.(type) {
			case []string:
				var builder strings.Builder
				for _, v := range l {
					builder.WriteString(fmt.Sprintf(`'%s', `, v))
				}
				*cl = append(*cl, fmt.Sprintf("%s %s IN (%s)", w.rel, w.field.fullColName(), strings.Trim(builder.String(), ", ")))
			case []int:
				var builder strings.Builder
				for _, v := range l {
					builder.WriteString(fmt.Sprintf("%d, ", v))
				}
				*cl = append(*cl, fmt.Sprintf("%s %s IN (%s)", w.rel, w.field.fullColName(), strings.Trim(builder.String(), ", ")))
			case []float64:
				var builder strings.Builder
				for _, v := range l {
					builder.WriteString(fmt.Sprintf("%f, ", v))
				}
				*cl = append(*cl, fmt.Sprintf("%s %s IN (%s)", w.rel, w.field.fullColName(), strings.Trim(builder.String(), ", ")))
			case []time.Time:
				var builder strings.Builder
				switch w.field.(type) {
				case *Date:
					for _, v := range l {
						builder.WriteString(fmt.Sprintf(`'%s', `, v.Format("2006-01-02")))
					}
				case *Datetime:
					for _, v := range l {
						builder.WriteString(fmt.Sprintf(`'%s', `, v.Format("2006-01-02 15:04:05")))
					}
				}
				*cl = append(*cl, fmt.Sprintf("%s %s IN (%s)", w.rel, w.field.fullColName(), strings.Trim(builder.String(), ", ")))
			}
			return
		}
		*cl = append(*cl, fmt.Sprintf("%s %s %s ?", w.rel, w.field.fullColName(), w.op))
		*vl = append(*vl, w.value)
	case byExpr:
		*cl = append(*cl, fmt.Sprintf("%s %s %s %s", w.rel, w.field.fullColName(), w.op, w.expr))
	}
}

type whereList []*where

func (l whereList) toClause() (string, []interface{}) {
	if len(l) == 0 {
		return "", nil
	}
	cl := make([]string, 0, len(l))
	vl := make([]interface{}, 0, len(l))
	for _, where := range l {
		where.toClause(&cl, &vl)
	}
	return fmt.Sprintf(fmt.Sprintf("WHERE %s", strings.Trim(strings.Trim(strings.Trim(strings.Join(cl, " "), " "), "AND"), "OR"))), vl
}
