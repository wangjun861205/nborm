package nborm

import "io"

type whereRelation string

const (
	whereAnd whereRelation = "AND"
	whereOr  whereRelation = "OR"
)

type wherer interface {
	clauser
	referencer
	nextNode() wherer
	lastNode() wherer
	append(w wherer)
}

type where struct {
	expr *Expr
	rel  whereRelation
	next wherer
}

func (w *where) toClause(wr io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if w == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		wr.Write([]byte("WHERE "))
	}
	if *isFirstNode {
		*isFirstNode = false
	} else {
		wr.Write([]byte(string(w.rel)))
		wr.Write([]byte(" "))
	}
	w.expr.toClause(wr, vals, nil, nil)
	if w.next != nil {
		w.next.toClause(wr, vals, isFirstGroup, isFirstNode)
	}
}

func (w *where) toSimpleClause(wr io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if w == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		wr.Write([]byte("WHERE "))
	}
	if *isFirstNode {
		*isFirstNode = false
	} else {
		wr.Write([]byte(string(w.rel)))
		wr.Write([]byte(" "))
	}
	w.expr.toSimpleClause(wr, vals, nil, nil)
	if w.next != nil {
		w.next.toSimpleClause(wr, vals, isFirstGroup, isFirstNode)
	}
}

func (w *where) toRefClause(wr io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if w == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
	}
	if *isFirstNode {
		*isFirstNode = false
	} else {
		wr.Write([]byte(string(w.rel)))
		wr.Write([]byte(" "))
	}
	w.expr.toRefClause(wr, vals, nil, nil)
	if w.next != nil {
		w.next.toRefClause(wr, vals, isFirstGroup, isFirstNode)
	}
}

func (w *where) toSimpleRefClause(wr io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if w == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
	}
	if *isFirstNode {
		*isFirstNode = false
	} else {
		wr.Write([]byte(string(w.rel)))
		wr.Write([]byte(" "))
	}
	w.expr.toSimpleRefClause(wr, vals, nil, nil)
	if w.next != nil {
		w.next.toSimpleRefClause(wr, vals, isFirstGroup, isFirstNode)
	}
}

func (w *where) nextNode() wherer {
	return w.next
}

func (w *where) lastNode() wherer {
	if w == nil {
		return nil
	}
	var last wherer = w
	for last.nextNode() != nil {
		last = last.nextNode()
	}
	return last
}

func (w *where) append(wherer wherer) {
	w.next = wherer
}

func newWhere(expr *Expr, rel whereRelation) *where {
	return &where{
		expr: expr,
		rel:  rel,
	}
}

type whereGroup struct {
	first wherer
	rel   whereRelation
	next  wherer
}

func (g *whereGroup) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if g == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("WHERE ("))
	} else {
		w.Write([]byte(string(g.rel)))
		w.Write([]byte(" ("))
	}
	g.first.toClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(") "))
	*isFirstNode = true
	if g.next != nil {
		g.next.toClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (g *whereGroup) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if g == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("WHERE ("))
	} else {
		w.Write([]byte(string(g.rel)))
		w.Write([]byte(" ("))
	}
	g.first.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(") "))
	*isFirstNode = true
	if g.next != nil {
		g.next.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (g *whereGroup) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if g == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("WHERE ("))
	} else {
		w.Write([]byte(string(g.rel)))
		w.Write([]byte(" ("))
	}
	g.first.toRefClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(") "))
	*isFirstNode = true
	if g.next != nil {
		g.next.toRefClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (g *whereGroup) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if g == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("WHERE ("))
	} else {
		w.Write([]byte(string(g.rel)))
		w.Write([]byte(" ("))
	}
	g.first.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(") "))
	*isFirstNode = true
	if g.next != nil {
		g.next.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (g *whereGroup) nextNode() wherer {
	return g.next
}

func (g *whereGroup) lastNode() wherer {
	if g == nil {
		return nil
	}
	var last wherer = g
	for last.nextNode() != nil {
		last = last.nextNode()
	}
	return last
}

func (g *whereGroup) append(w wherer) {
	g.next = w
}

func groupWherers(groupRel whereRelation, wherers ...wherer) *whereGroup {
	if len(wherers) == 0 {
		return nil
	}
	g := new(whereGroup)
	g.first = wherers[0]
	for _, w := range wherers[1:] {
		g.first.lastNode().append(w)
	}
	g.rel = groupRel
	return g
}
