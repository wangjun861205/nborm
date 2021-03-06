package nborm

import "io"

type update struct {
	field referencer
	value clauser
}

func newUpdate(field referencer, value clauser) *update {
	return &update{field, value}
}

func (u *update) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		u.field.toRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte("= "))
		u.value.toClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	} else {
		w.Write([]byte(", "))
		u.field.toRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte("= "))
		u.value.toClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (u *update) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		u.field.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte("= "))
		u.value.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	} else {
		w.Write([]byte(", "))
		u.field.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte("= "))
		u.value.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	}
}

type updateList []*update

func (l updateList) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("SET "))
	}
	l[0].toClause(w, vals, isFirstGroup, isFirstNode)
	l = l[1:]
	l.toClause(w, vals, isFirstGroup, isFirstNode)
}

func (l updateList) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("SET "))
	}
	l[0].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	l = l[1:]
	l.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}
