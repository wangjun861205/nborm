package nborm

import "fmt"

type orderType string

const (
	asc  orderType = "ASC"
	desc orderType = "DESC"
)

type orderBy struct {
	refClauser refClauser
	order      orderType
}

func newOrderBy(refClauser refClauser, order orderType) *orderBy {
	return &orderBy{refClauser, order}
}

func (o *orderBy) toRefClause() string {
	return fmt.Sprintf("%s %s", o.refClauser.toRefClause(), o.order)
}

func (o *orderBy) toSimpleRefClause() string {
	return fmt.Sprintf("%s %s", o.refClauser.toSimpleRefClause(), o.order)
}
