package nborm

import "io"

type insertFielder interface {
	referencer
	sqlLiteral() string
}

type insertValuer interface {
	clauser
	sqlLiteral() string
}

type insert struct {
	// field referencer
	// value clauser
	field insertFielder
	value insertValuer
}

// func newInsert(field referencer, value clauser) *insert {
// 	return &insert{field, value}
// }

func newInsert(field insertFielder, value insertValuer) *insert {
	return &insert{field, value}
}

func (i *insert) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		i.field.toRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" = "))
		i.value.toClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	} else {
		w.Write([]byte(", "))
		i.field.toRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" = "))
		i.value.toClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (i *insert) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		i.field.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" = "))
		i.value.toClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	} else {
		w.Write([]byte(", "))
		i.field.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" = "))
		i.value.toClause(w, vals, isFirstGroup, isFirstNode)
	}
}

func (i *insert) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		i.field.toRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	} else {
		w.Write([]byte(", "))
		i.field.toRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	}
}

func (i *insert) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		i.field.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	} else {
		w.Write([]byte(", "))
		i.field.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	}
}

func (i *insert) toValueListClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		i.value.toClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	} else {
		w.Write([]byte(", "))
		i.value.toClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	}
}

func (i *insert) toCSVColName(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		i.field.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	} else {
		w.Write([]byte(", "))
		i.field.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
		w.Write([]byte(" "))
	}
}

func (i *insert) toCSVVal(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		exp := i.field.(ValueField).getExpr()
		if exp == nil {
			w.Write([]byte(i.field.sqlLiteral()))
		}
	} else {
		exp := i.field.(ValueField).getExpr()
		if exp == nil {
			w.Write([]byte(", "))
			w.Write([]byte(i.field.sqlLiteral()))
		}
	}
}

func (i *insert) toLoadDataExpr(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		exp := i.field.(ValueField).getExpr()
		if exp != nil {
			w.Write([]byte("SET "))
			i.field.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
			w.Write([]byte(" = "))
			w.Write([]byte(exp.exp))
		}
	} else {
		exp := i.field.(ValueField).getExpr()
		if exp != nil {
			w.Write([]byte(", "))
			i.field.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
			w.Write([]byte(" = "))
			w.Write([]byte(exp.exp))
		}
	}
}

type insertList []*insert

func (l insertList) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("SET "))
	}
	l[0].toClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toClause(w, vals, isFirstGroup, isFirstNode)
}

func (l insertList) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("SET "))
	}
	l[0].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}

func (l insertList) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("( "))
	}
	l[0].toRefClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toRefClause(w, vals, isFirstGroup, isFirstNode)
}

func (l insertList) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("( "))
	}
	l[0].toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
}

func (l insertList) toValueListClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	if *isFirstGroup {
		*isFirstGroup = false
		w.Write([]byte("VALUES ("))
	}
	l[0].toValueListClause(w, vals, isFirstGroup, isFirstNode)
	l[1:].toValueListClause(w, vals, isFirstGroup, isFirstNode)
}
