package nborm

import (
	"io"
)

type orderType string

const (
	asc  orderType = "ASC"
	desc orderType = "DESC"
)

type orderBy struct {
	referencer referencer
	order      orderType
}

func newOrderBy(referencer referencer, order orderType) *orderBy {
	return &orderBy{referencer, order}
}

func (o *orderBy) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		w.Write([]byte("ORDER BY "))
	}
	o.referencer.toRefClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(string(o.order)))
	w.Write([]byte(" "))
}

func (o *orderBy) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		w.Write([]byte("ORDER BY "))
	}
	o.referencer.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(string(o.order)))
	w.Write([]byte(" "))
}

type orderByList []*orderBy

func (l orderByList) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	l[0].toClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toClause(w, vals, isFirstGroup, isFirstNode)
}

func (l orderByList) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	l[0].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}
