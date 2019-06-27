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
		*cl = append(*cl, fmt.Sprintf("%s = ?", u.field.fullColName()))
		*vl = append(*vl, u.value)
		return
	}
	*cl = append(*cl, fmt.Sprintf("%s = %s", u.field.fullColName()))
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
	eq    operator = "="
	neq   operator = "<>"
	lt    operator = "<"
	lte   operator = "<="
	gt    operator = ">"
	gte   operator = ">="
	is    operator = "IS NULL"
	nis   operator = "IS NOT NULL"
	like  operator = "LIKE"
	nlike operator = "NOT LIKE"
	in    operator = "IN"
	nin   operator = "NOT IN"
)

type rel string

const (
	and rel = "AND"
	or  rel = "OR"
)

func checkOp(op string) {
	switch op {
	case "=", "<>", "<", "<=", ">", ">=", "IS NULL", "IS NOT NULL", "LIKE", "NOT LIKE", "IN", "NOT IN":
		return
	default:
		panic(fmt.Sprintf("unsupported operator (%s)", op))
	}
}

type where struct {
	rel  rel
	expr *Expr
	val  []interface{}
	next *where
}

func newWhere(rel rel, expr *Expr, val ...interface{}) *where {
	return &where{
		rel,
		expr,
		val,
		nil,
	}
}

func (w *where) append(nw *where) *where {
	if w == nil {
		return nw
	}
	if nw == nil {
		return w
	}
	lastWhere := w
	for lastWhere.next != nil {
		lastWhere = lastWhere.next
	}
	lastWhere.next = nw
	return w
}

func (w *where) toClause(cl *[]string, vl *[]interface{}) {
	if w == nil {
		return
	}
	*cl = append(*cl, fmt.Sprintf("%s %s", w.rel, w.expr.String()))
	for _, value := range w.val {
		switch v := value.(type) {
		case []int:
			for _, i := range v {
				*vl = append(*vl, i)
			}
		case []float32:
			for _, f := range v {
				*vl = append(*vl, f)
			}
		case []float64:
			for _, f := range v {
				*vl = append(*vl, f)
			}
		case []string:
			for _, s := range v {
				*vl = append(*vl, s)
			}
		case []time.Time:
			switch w.expr.fields[0].(type) {
			case *Date:
				for _, t := range v {
					*vl = append(*vl, t.Format("2006-01-02"))
				}
			default:
				for _, t := range v {
					*vl = append(*vl, t.Format("2006-01-02"))
				}
			}
		case time.Time:
			switch w.expr.fields[0].(type) {
			case *Date:
				*vl = append(*vl, v.Format("2006-01-02"))
			default:
				*vl = append(*vl, v.Format("2006-01-02"))
			}
		default:
			*vl = append(*vl, v)
		}
	}
	if w.next != nil {
		w.next.toClause(cl, vl)
	}
}

func (w *where) toSimpleClause(cl *[]string, vl *[]interface{}) {
	if w == nil {
		return
	}
	*cl = append(*cl, fmt.Sprintf("%s %s", w.rel, w.expr.SimpleString()))
	if len(w.val) > 0 {
		switch v := w.val[0].(type) {
		case []int:
			for _, i := range v {
				*vl = append(*vl, i)
			}
		case []float32:
			for _, f := range v {
				*vl = append(*vl, f)
			}
		case []float64:
			for _, f := range v {
				*vl = append(*vl, f)
			}
		case []string:
			for _, s := range v {
				*vl = append(*vl, s)
			}
		case []time.Time:
			switch w.expr.fields[0].(type) {
			case *Date:
				for _, t := range v {
					*vl = append(*vl, t.Format("2006-01-02"))
				}
			default:
				for _, t := range v {
					*vl = append(*vl, t.Format("2006-01-02"))
				}
			}
		case time.Time:
			switch w.expr.fields[0].(type) {
			case *Date:
				*vl = append(*vl, v.Format("2006-01-02"))
			default:
				*vl = append(*vl, v.Format("2006-01-02"))
			}
		default:
			*vl = append(*vl, w.val[0])
		}
	}
	if w.next != nil {
		w.next.toClause(cl, vl)
	}
}
