package nborm

import "io"

type insert struct {
	field referencer
	value clauser
}

func newInsert(field referencer, value clauser) *insert {
	return &insert{field, value}
}

func (i *insert) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		w.Write([]byte("SET "))
	} else {
		w.Write([]byte(", "))
	}
	i.field.toRefClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(" = "))
	i.value.toClause(w, vals, isFirstGroup, isFirstNode)
}

func (i *insert) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		w.Write([]byte("SET "))
	} else {
		w.Write([]byte(", "))
	}
	i.field.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(" = "))
	i.value.toClause(w, vals, isFirstGroup, isFirstNode)
}

type insertList []*insert

func (l insertList) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	l[0].toClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toClause(w, vals, isFirstGroup, isFirstNode)
}

func (l insertList) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	l[0].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}
