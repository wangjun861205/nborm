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

func (o *orderBy) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	o.referencer.toRefClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(string(o.order)))
	w.Write([]byte(" "))
}

func (o *orderBy) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode bool) {
	o.referencer.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(string(o.order)))
	w.Write([]byte(" "))
}
