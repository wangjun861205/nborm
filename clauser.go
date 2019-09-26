package nborm

import (
	"io"
)

type clauser interface {
	toClause(w io.Writer, vals *[]interface{})
	toSimpleClause(w io.Writer, vals *[]interface{})
}

type referencer interface {
	toRefClause(w io.Writer, vals *[]interface{})
	toSimpleRefClause(w io.Writer, vals *[]interface{})
}
