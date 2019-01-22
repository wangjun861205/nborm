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
	// value    interface{}
	value    []interface{}
	relation string
	next     *Where
	sub      *Where
}

// func newWhere(db, table, column, operator string, value interface{}) *Where {
// 	return &Where{db: db, table: table, column: column, operator: operator, value: value}
// }

func newWhere(db, table, column, operator string, value ...interface{}) *Where {
	return &Where{db: db, table: table, column: column, operator: operator, value: value}
}

func (w *Where) fullColumnName() string {
	return fmt.Sprintf("%s.%s.%s", wrap(w.db), wrap(w.table), wrap(w.column))
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
		switch {
		case w.value[0] == "NULL":
			// builder.WriteString(fmt.Sprintf("%s.%s.%s %s NULL", wrap(w.db), wrap(w.table), wrap(w.column), w.operator))
			builder.WriteString(fmt.Sprintf("%s %s NULL", w.fullColumnName(), w.operator))
		case w.operator == "IN":
			// builder.WriteString(fmt.Sprintf("%s.%s.%s IN %s", wrap(w.db), wrap(w.table), wrap(w.column), w.value))
			builder.WriteString(fmt.Sprintf("%s IN (%s)", w.fullColumnName(), genPlaceHolder(len(w.value))))
			valueList = append(valueList, w.value...)
		case w.operator == "NOT IN":
			// builder.WriteString(fmt.Sprintf("%s.%s.%s NOT IN %s", wrap(w.db), wrap(w.table), wrap(w.column), w.value))
			builder.WriteString(fmt.Sprintf("%s NOT IN (%s)", w.fullColumnName(), genPlaceHolder(len(w.value))))
			valueList = append(valueList, w.value...)
		default:
			builder.WriteString(fmt.Sprintf("%s.%s.%s %s ?", wrap(w.db), wrap(w.table), wrap(w.column), w.operator))
			valueList = append(valueList, w.value[0])
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
	autoIncrementType whereType = iota
	primaryKeyType
	uniqueKeyType
	otherType
)

func genWhere(addr uintptr, tabInfo *TableInfo) (*Where, whereType) {
	inc := getIncWithTableInfo(addr, tabInfo)
	if inc != nil && inc.IsValid() && !inc.IsNull() {
		return inc.where(), autoIncrementType
	}
	var where *Where
	pk := getPrimaryKeyFieldsWithTableInfo(addr, tabInfo)
	if len(pk) > 0 {
		for _, col := range pk {
			where = where.And(col.where())
		}

	}
	if where != nil {
		return where, primaryKeyType
	}
	uni := getUniqueFieldsWithTableInfo(addr, tabInfo)
	if len(uni) > 0 {
		for _, col := range uni {
			where = where.And(col.where())
		}
	}
	if where != nil {
		return where, uniqueKeyType
	}
	for _, col := range tabInfo.Columns {
		field := getFieldByColumnInfo(addr, col)
		if field.IsValid() {
			where = where.And(field.where())
		}
	}
	return where, otherType
}
