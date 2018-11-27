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

func (w *Where) toAndClause() (string, []interface{}) {
	if w == nil {
		return "", nil
	}
	colStr, valList := w.toSQL()
	return "AND " + colStr, valList
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

//And and operation
func (w *Where) And(other *Where) *Where {
	if w == nil {
		return other
	}
	if other == nil {
		return w
	}
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

type whereType int

const (
	autoIncrement whereType = iota
	primaryKey
	uniqueKeys
	others
)

func genWhere(addr uintptr, tabInfo *tableInfo) (*Where, whereType) {
	inc := getIncWithTableInfo(addr, tabInfo)
	if inc != nil && inc.IsValid() {
		return inc.where(), autoIncrement
	}
	pks := getPksWithTableInfo(addr, tabInfo)
	var where *Where
	for _, pk := range pks {
		if !pk.IsValid() {
			where = nil
			break
		}
		where = where.And(pk.where())
	}
	if where != nil {
		return where, primaryKey
	}
	unis := getUinsWithTableInfo(addr, tabInfo)
	if len(unis) > 0 {
		for _, uni := range unis {
			if uni.IsValid() {
				return uni.where(), uniqueKeys
			}
		}
	}
	for _, col := range tabInfo.columns {
		field := getFieldByColumnInfo(addr, col)
		if field.IsValid() {
			where = where.And(field.where())
		}
	}
	return where, others
}
