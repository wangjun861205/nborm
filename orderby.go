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

func (o *orderBy) toRefClause(w io.Writer, vals *[]interface{}) {
	o.referencer.toRefClause(w, vals)
	w.Write([]byte(string(o.order)))
	w.Write([]byte(" "))
}

func (o *orderBy) toSimpleRefClause(w io.Writer, vals *[]interface{}) {
	o.referencer.toSimpleRefClause(w, vals)
	w.Write([]byte(string(o.order)))
	w.Write([]byte(" "))
}
