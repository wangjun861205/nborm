package nborm

import "io"

type groupBy struct {
	referencer
}

func (g groupBy) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		w.Write([]byte("GROUP BY"))
	}
	g.referencer.toRefClause(w, vals, isFirstGroup, isFirstNode)
}

func (g groupBy) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		w.Write([]byte("GROUP BY"))
	}
	g.referencer.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
}

type groupByList []*groupBy

func (l groupByList) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	l[0].toClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toClause(w, vals, isFirstGroup, isFirstNode)
}

func (l groupByList) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	l[0].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}
