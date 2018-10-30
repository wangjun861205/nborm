package nborm

import (
	"fmt"
	"strings"
)

//Where is used for query
type Where struct {
	db       string
	table    string
	column   string
	operator string
	value    interface{}
	relation string
	next     *Where
	sub      *Where
}

func newWhere(db, table, column, operator string, value interface{}) *Where {
	return &Where{db: db, table: table, column: column, operator: operator, value: value}
}

func (w *Where) toClause() (string, []interface{}) {
	if w == nil {
		return "", nil
	}
	colStr, valList := w.toSQL()
	return "WHERE " + colStr, valList
}

func (w *Where) toSQL() (string, []interface{}) {
	builder := strings.Builder{}
	valueList := make([]interface{}, 0, 8)
	if w.sub != nil {
		subStr, l := w.sub.toSQL()
		builder.WriteString(fmt.Sprintf("(%s)", subStr))
		valueList = append(valueList, l...)
	} else {
		if w.value == "NULL" {
			builder.WriteString(fmt.Sprintf("%s.%s.%s %s NULL", w.db, w.table, w.column, w.operator))
		} else {
			builder.WriteString(fmt.Sprintf("%s.%s.%s %s ?", w.db, w.table, w.column, w.operator))
			valueList = append(valueList, w.value)
		}
	}
	if w.next != nil {
		nextStr, l := w.next.toSQL()
		builder.WriteString(fmt.Sprintf(" %s %s", w.relation, nextStr))
		valueList = append(valueList, l...)
	}
	return builder.String(), valueList
}

func (w *Where) toAndSQL() (string, []interface{}) {
	s, l := w.toSQL()
	return " AND " + s, l

}

//And and operation
func (w *Where) And(other *Where) *Where {
	end := w
	for end.next != nil {
		end = end.next
	}
	end.relation = "AND"
	end.next = other
	return w
}

//Or or operation
func (w *Where) Or(other *Where) *Where {
	end := w
	for end.next != nil {
		end = end.next
	}
	end.relation = "OR"
	end.next = other
	return w
}

//Group group operation
func Group(w *Where) *Where {
	return &Where{sub: w}
}
