package nborm

import (
	"io"
)

type clauser interface {
	toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool)
	toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool)
}

type referencer interface {
	toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool)
	toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool)
}
