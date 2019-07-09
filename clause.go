package nborm

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type clauseStatus int

// const (
// 	byValue clauseStatus = iota
// 	byExpr
// )

type update struct {
	expr   *Expr
	values []interface{}
}

func newUpdate(expr *Expr, values ...interface{}) *update {
	return &update{expr, values}
}

func (u *update) toClause() (string, []interface{}) {
	return u.expr.String(), u.values
}

func (u *update) toSimpleClause() (string, []interface{}) {
	return u.expr.SimpleString(), u.values
}

type updateList []*update

func (ul updateList) toClause() (string, []interface{}) {
	if len(ul) == 0 {
		return "", nil
	}
	cl := make([]string, 0, len(ul))
	vl := make([]interface{}, 0, len(ul))
	for _, u := range ul {
		c, v := u.toClause()
		cl = append(cl, c)
		vl = append(vl, v...)
	}
	return strings.Join(cl, ", "), vl
}

func (ul updateList) toSimpleClause() (string, []interface{}) {
	if len(ul) == 0 {
		return "", nil
	}
	cl := make([]string, 0, len(ul))
	vl := make([]interface{}, 0, len(ul))
	for _, u := range ul {
		c, v := u.toSimpleClause()
		cl = append(cl, c)
		vl = append(vl, v...)
	}
	return strings.Join(cl, ", "), vl
}

// type updateSet struct {
// 	field  Field
// 	value  interface{}
// 	expr   *Expr
// 	status clauseStatus
// }

// func newUpdateSet(field Field, value interface{}) *updateSet {
// 	set := &updateSet{field: field}
// 	if v, ok := value.(*Expr); ok {
// 		set.expr = v
// 		set.status = byExpr
// 	} else {
// 		set.value = value
// 		set.status = byValue
// 	}
// 	return set
// }

// func (u *updateSet) toClause(cl *[]string, vl *[]interface{}) {
// 	if u.status == byValue {
// 		*cl = append(*cl, fmt.Sprintf("%s = ?", u.field.fullColName()))
// 		*vl = append(*vl, u.value)
// 		return
// 	}
// 	*cl = append(*cl, fmt.Sprintf("%s = %s", u.field.fullColName(), u.expr.String()))
// }

// type updateSetList []*updateSet

// func (l updateSetList) toClause() (string, []interface{}) {
// 	cl := make([]string, 0, len(l))
// 	vl := make([]interface{}, 0, len(l))
// 	for _, u := range l {
// 		u.toClause(&cl, &vl)
// 	}
// 	return fmt.Sprintf("SET %s", strings.Join(cl, ",")), vl
// }

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
}

func newWhere(rel rel, expr *Expr, val ...interface{}) *where {
	return &where{
		rel,
		expr,
		val,
	}
}

type whereList []*where

func (l whereList) toClause() (string, []interface{}) {
	if len(l) == 0 {
		return "", nil
	}
	var builder strings.Builder
	vals := make([]interface{}, 0, 8)
	for _, w := range l {
		builder.WriteString(fmt.Sprintf("%s %s ", w.rel, w.expr.String()))
		for i, val := range w.val {
			switch v := val.(type) {
			case []string:
				for _, s := range v {
					vals = append(vals, s)
				}
			case []int:
				for _, i := range v {
					vals = append(vals, i)
				}
			case []float32:
				for _, f := range v {
					vals = append(vals, f)
				}
			case []float64:
				for _, f := range v {
					vals = append(vals, f)
				}
			case []time.Time:
				switch w.expr.fields[i].(type) {
				case *Date:
					for _, t := range v {
						vals = append(vals, t.Format("2006-01-02"))
					}
				case *Datetime:
					for _, t := range v {
						vals = append(vals, t.Format("2006-01-02 15:04:05"))
					}
				default:
					panic(errors.New("invalid field type for []time.Time where value"))
				}
			default:
				vals = append(vals, w.val...)
			}
		}
	}
	return fmt.Sprintf("WHERE %s", strings.TrimPrefix(strings.TrimPrefix(builder.String(), "AND "), "OR ")), vals
}

func (l whereList) toSimpleClause() (string, []interface{}) {
	if len(l) == 0 {
		return "", nil
	}
	var builder strings.Builder
	vals := make([]interface{}, 0, 8)
	for _, w := range l {
		builder.WriteString(fmt.Sprintf("%s %s ", w.rel, w.expr.SimpleString()))
		for i, val := range w.val {
			switch v := val.(type) {
			case []string:
				for _, s := range v {
					vals = append(vals, s)
				}
			case []int:
				for _, i := range v {
					vals = append(vals, i)
				}
			case []float32:
				for _, f := range v {
					vals = append(vals, f)
				}
			case []float64:
				for _, f := range v {
					vals = append(vals, f)
				}
			case []time.Time:
				switch w.expr.fields[i].(type) {
				case *Date:
					for _, t := range v {
						vals = append(vals, t.Format("2006-01-02"))
					}
				case *Datetime:
					for _, t := range v {
						vals = append(vals, t.Format("2006-01-02 15:04:05"))
					}
				default:
					panic(errors.New("invalid field type for []time.Time where value"))
				}
			default:
				vals = append(vals, w.val...)
			}
		}
	}
	return fmt.Sprintf("WHERE %s", strings.TrimPrefix(strings.TrimPrefix(builder.String(), "AND "), "OR ")), vals
}

type havingList []*where

func (l havingList) toClause() (string, []interface{}) {
	if len(l) == 0 {
		return "", nil
	}
	var builder strings.Builder
	vals := make([]interface{}, 0, 8)
	for _, w := range l {
		builder.WriteString(fmt.Sprintf("%s %s ", w.rel, w.expr.String()))
		for i, val := range w.val {
			switch v := val.(type) {
			case []string:
				for _, s := range v {
					vals = append(vals, s)
				}
			case []int:
				for _, i := range v {
					vals = append(vals, i)
				}
			case []float32:
				for _, f := range v {
					vals = append(vals, f)
				}
			case []float64:
				for _, f := range v {
					vals = append(vals, f)
				}
			case []time.Time:
				switch w.expr.fields[i].(type) {
				case *Date:
					for _, t := range v {
						vals = append(vals, t.Format("2006-01-02"))
					}
				case *Datetime:
					for _, t := range v {
						vals = append(vals, t.Format("2006-01-02 15:04:05"))
					}
				default:
					panic(errors.New("invalid field type for []time.Time where value"))
				}
			default:
				vals = append(vals, w.val...)
			}
		}
	}
	return fmt.Sprintf("HAVING %s", strings.TrimPrefix(strings.TrimPrefix(builder.String(), "AND "), "OR ")), vals
}

func (l havingList) toSimpleClause() (string, []interface{}) {
	if len(l) == 0 {
		return "", nil
	}
	var builder strings.Builder
	vals := make([]interface{}, 0, 8)
	for _, w := range l {
		builder.WriteString(fmt.Sprintf("%s %s ", w.rel, w.expr.SimpleString()))
		for i, val := range w.val {
			switch v := val.(type) {
			case []string:
				for _, s := range v {
					vals = append(vals, s)
				}
			case []int:
				for _, i := range v {
					vals = append(vals, i)
				}
			case []float32:
				for _, f := range v {
					vals = append(vals, f)
				}
			case []float64:
				for _, f := range v {
					vals = append(vals, f)
				}
			case []time.Time:
				switch w.expr.fields[i].(type) {
				case *Date:
					for _, t := range v {
						vals = append(vals, t.Format("2006-01-02"))
					}
				case *Datetime:
					for _, t := range v {
						vals = append(vals, t.Format("2006-01-02 15:04:05"))
					}
				default:
					panic(errors.New("invalid field type for []time.Time where value"))
				}
			default:
				vals = append(vals, w.val...)
			}
		}
	}
	return fmt.Sprintf("HAVING %s", strings.TrimPrefix(strings.TrimPrefix(builder.String(), "AND "), "OR ")), vals
}
