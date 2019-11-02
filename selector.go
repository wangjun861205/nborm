package nborm

import (
	"database/sql"
	"io"
)

type selectorStatus int

const (
	selectorStatusDistinct  selectorStatus = 1
	selectorStatusFoundRows selectorStatus = 1 << 1
)

type selector interface {
	clauser
	sql.Scanner
	toScan(m Model, selectors *[]interface{})
	getField() Field
}

type selectorList struct {
	status selectorStatus
	list   []selector
}

func (l *selectorList) addStatus(status selectorStatus) {
	l.status |= status
}

func (l *selectorList) checkStatus(status selectorStatus) bool {
	return l.status&status == status
}

func (l selectorList) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l.list) == 0 {
		return
	}
	if *isFirstNode {
		*isFirstNode = false
		w.Write([]byte("SELECT "))
		if l.checkStatus(selectorStatusFoundRows) {
			w.Write([]byte("SQL_CALC_FOUND_ROWS "))
		}
		if l.checkStatus(selectorStatusDistinct) {
			w.Write([]byte("DISTINCT"))
		}
		l.list[0].toClause(w, vals, isFirstGroup, isFirstNode)
	} else {
		w.Write([]byte(", "))
		l.list[0].toClause(w, vals, isFirstGroup, isFirstNode)
	}
	l.list = l.list[1:]
	l.toClause(w, vals, isFirstGroup, isFirstNode)
}

func (l selectorList) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l.list) == 0 {
		return
	}
	if *isFirstNode {
		*isFirstNode = false
		w.Write([]byte("SELECT "))
		l.list[0].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	} else {
		w.Write([]byte(", "))
		l.list[0].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	}
	l.list = l.list[1:]
	l.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}

type exprSelect struct {
	field Field
	expr  *Expr
}

func newExprSelect(f Field, expr *Expr) *exprSelect {
	return &exprSelect{f, expr}
}

func (s *exprSelect) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	s.expr.toClause(w, vals, isFirstGroup, isFirstNode)
}

func (s *exprSelect) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	s.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}

func (s *exprSelect) Scan(v interface{}) error {
	return s.field.Scan(v)
}

func (s *exprSelect) toScan(m Model, selectors *[]interface{}) {
	s.field.toScan(m, selectors)
}

func (s *exprSelect) getField() Field {
	return s.field
}
