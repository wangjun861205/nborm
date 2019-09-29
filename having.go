package nborm

import "io"

type havinger interface {
	clauser
	referencer
	nextNode() havinger
	lastNode() havinger
	append(w havinger)
}

type having struct {
	expr *Expr
	rel  whereRelation
	next havinger
}

func (h *having) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if h == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("HAVING "))
	}
	if *isFirstNode {
		*isFirstNode = false
	} else {
		w.Write([]byte(string(h.rel)))
		w.Write([]byte(" "))
	}
	h.expr.toClause(w, vals, isFirstGroup, isFirstNode)
	if h.next != nil {
		h.next.toClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (h *having) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if h == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("HAVING "))
	}
	if *isFirstNode {
		*isFirstNode = false
	} else {
		w.Write([]byte(string(h.rel)))
		w.Write([]byte(" "))
	}
	h.expr.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	if h.next != nil {
		h.next.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (h *having) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if h == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
	}
	if *isFirstNode {
		*isFirstNode = false
	} else {
		w.Write([]byte(string(h.rel)))
		w.Write([]byte(" "))
	}
	h.expr.toRefClause(w, vals, isFirstGroup, isFirstNode)
	if h.next != nil {
		h.next.toRefClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (h *having) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if h == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
	}
	if *isFirstNode {
		*isFirstNode = false
	} else {
		w.Write([]byte(string(h.rel)))
		w.Write([]byte(" "))
	}
	h.expr.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
	if h.next != nil {
		h.next.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (h *having) nextNode() havinger {
	return h.next
}

func (h *having) lastNode() havinger {
	if h == nil {
		return nil
	}
	var last havinger = h
	for last.nextNode() != nil {
		last = last.nextNode()
	}
	return last
}

func (h *having) append(havinger havinger) {
	h.next = havinger
}

func newHaving(expr *Expr, rel whereRelation) *having {
	return &having{
		expr: expr,
		rel:  rel,
	}
}

type havingGroup struct {
	first havinger
	rel   whereRelation
	next  havinger
}

func (g *havingGroup) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if g == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("HAVING ("))
	}
	if *isFirstNode {
		*isFirstNode = false
	} else {
		w.Write([]byte(string(g.rel)))
		w.Write([]byte(" "))
	}
	g.first.toClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(")"))
	if g.next != nil {
		g.next.toClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (g *havingGroup) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if g == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("HAVING ("))
	}
	if *isFirstNode {
		*isFirstNode = false
	} else {
		w.Write([]byte(string(g.rel)))
		w.Write([]byte(" "))
	}
	g.first.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(")"))
	if g.next != nil {
		g.next.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (g *havingGroup) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if g == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("("))
	}
	if *isFirstNode {
		*isFirstNode = false
	} else {
		w.Write([]byte(string(g.rel)))
		w.Write([]byte(" "))
	}
	g.first.toRefClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(")"))
	if g.next != nil {
		g.next.toRefClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (g *havingGroup) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if g == nil {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("("))
	}
	if *isFirstNode {
		*isFirstNode = false
	} else {
		w.Write([]byte(string(g.rel)))
		w.Write([]byte(" "))
	}
	g.first.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(")"))
	if g.next != nil {
		g.next.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (g *havingGroup) nextNode() havinger {
	return g.next
}

func (g *havingGroup) lastNode() havinger {
	if g == nil {
		return nil
	}
	var last havinger = g
	for last.nextNode() != nil {
		last = last.nextNode()
	}
	return last
}

func (g *havingGroup) append(w havinger) {
	g.next = w
}

func groupHavings(groupRel whereRelation, havingers ...havinger) *havingGroup {
	if len(havingers) == 0 {
		return nil
	}
	g := new(havingGroup)
	g.first = havingers[0]
	for _, w := range havingers[1:] {
		g.first.lastNode().append(w)
	}
	g.rel = groupRel
	return g
}
