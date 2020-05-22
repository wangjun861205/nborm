package nborm

import "io"

type groupBy struct {
	referencer
}

func (g groupBy) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		g.referencer.toRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	} else {
		w.Write([]byte(", "))
		g.referencer.toRefClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (g groupBy) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		g.referencer.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	} else {
		w.Write([]byte(", "))
		g.referencer.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
	}
}

type groupByList []*groupBy

func (l groupByList) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("GROUP BY "))
	}
	l[0].toClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toClause(w, vals, isFirstGroup, isFirstNode)
}

func (l groupByList) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("GROUP BY "))
	}
	l[0].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}
