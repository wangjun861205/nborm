package nborm

import (
	"database/sql"
	"io"
)

type selector interface {
	clauser
	sql.Scanner
	findOrCopy(m Model)
}

type selectorList []selector

func (l selectorList) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	if len(l) == 0 {
		return
	}
	if isFirstNode {
		isFirstNode = false
		w.Write([]byte("SELECT "))
		l[0].toClause(w, vals, isFirstGroup, isFirstNode)
	} else {
		w.Write([]byte(", "))
		l[0].toClause(w, vals, isFirstGroup, isFirstNode)
	}
	selectorList(l[1:]).toClause(w, vals, isFirstGroup, isFirstNode)
}

func (l selectorList) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	if len(l) == 0 {
		return
	}
	if isFirstNode {
		isFirstNode = false
		w.Write([]byte("SELECT "))
		l[0].toClause(w, vals, isFirstGroup, isFirstNode)
	} else {
		w.Write([]byte(", "))
		l[0].toClause(w, vals, isFirstGroup, isFirstNode)
	}
	selectorList(l[1:]).toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}
