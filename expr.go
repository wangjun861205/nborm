package nborm

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

var startAndOrRe = regexp.MustCompile(`^\s*(AND|And|and|OR|Or|or)\s*`)

func trimPreAndOr(s string) string {
	return startAndOrRe.ReplaceAllString(s, "")
}

type expStat int

const (
	normal  expStat = 0
	inQuote expStat = 1
)

type exprType int

const (
	whereExpr exprType = iota
	// updateExpr
	havingExpr
	assignExpr
)

// Expr 自定义sql表达式
type Expr struct {
	exp    string
	values []interface{}
}

// NewExpr 创建新的自定义sql表达式
func NewExpr(sqlexpr string, values ...interface{}) *Expr {
	return &Expr{
		sqlexpr,
		values,
	}
}

// toClause 自定义表达式转化为sql的子句
func (e *Expr) toClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(e.values) == 0 {
		w.Write([]byte(e.exp))
		w.Write([]byte(" "))
		return
	}
	reader := strings.NewReader(e.exp)
	stat := normal
	quoteStack := make([]rune, 0, 8)
	var fieldIndex int
	for r, _, err := reader.ReadRune(); ; r, _, err = reader.ReadRune() {
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		switch r {
		case '"', '`', '\'':
			if stat == normal {
				stat = inQuote
				quoteStack = append(quoteStack, r)
				w.Write([]byte(string(r)))
			} else {
				if quoteStack[len(quoteStack)-1] == r {
					quoteStack = quoteStack[:len(quoteStack)-1]
					if len(quoteStack) == 0 {
						stat = normal
					}

				} else {
					quoteStack = append(quoteStack, r)
					w.Write([]byte(string(r)))
				}
			}
		case '@':
			if stat == normal {
				switch c := e.values[fieldIndex].(type) {
				case clauser:
					c.toClause(w, vals, isFirstGroup, isFirstNode)
				default:
					panic(fmt.Errorf("invalid argument type (%T)", c))
				}
				fieldIndex++
			} else {
				w.Write([]byte(string(r)))
			}
		case '?':
			if stat == normal {
				switch c := e.values[fieldIndex].(type) {
				case *Expr:
					c.toClause(w, vals, isFirstGroup, isFirstNode)
				default:
					w.Write([]byte(valToPlaceholder(c)))
					*vals = append(*vals, expandArg(c)...)
				}
				fieldIndex++
			} else {
				w.Write([]byte(string(r)))
			}
		default:
			w.Write([]byte(string(r)))
		}
	}
	if stat != normal {
		panic(fmt.Errorf("no closed expression (exp: %s)", e.exp))
	}
	if len(e.values) != fieldIndex {
		panic(fmt.Errorf("extra field for express (fields: %v)", e.values))
	}
	w.Write([]byte(" "))
}

// toSimpleClause 自定义表达式转化为sql子句，其中涉及到表名的全部用表的原名而不是用别名
func (e *Expr) toSimpleClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(e.values) == 0 {
		w.Write([]byte(e.exp))
		w.Write([]byte(" "))
		return
	}
	reader := strings.NewReader(e.exp)
	stat := normal
	quoteStack := make([]rune, 0, 8)
	var fieldIndex int
	for r, _, err := reader.ReadRune(); ; r, _, err = reader.ReadRune() {
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		switch r {
		case '"', '`', '\'':
			if stat == normal {
				stat = inQuote
				quoteStack = append(quoteStack, r)
				w.Write([]byte(string(r)))
			} else {
				if quoteStack[len(quoteStack)-1] == r {
					quoteStack = quoteStack[:len(quoteStack)-1]
					if len(quoteStack) == 0 {
						stat = normal
					}
					w.Write([]byte(string(r)))
				} else {
					quoteStack = append(quoteStack, r)
					w.Write([]byte(string(r)))
				}
			}
		case '@':
			if stat == normal {
				switch c := e.values[fieldIndex].(type) {
				case clauser:
					c.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
				default:
					panic(fmt.Errorf("invalid argument type (expr: %s, type: %T, values: %v)", e.exp, c, e.values))
				}
				fieldIndex++
			} else {
				w.Write([]byte(string(r)))
			}
		case '?':
			if stat == normal {
				switch c := e.values[fieldIndex].(type) {
				case *Expr:
					c.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
				default:
					w.Write([]byte(valToPlaceholder(c)))
					*vals = append(*vals, expandArg(c)...)
				}
				fieldIndex++
			} else {
				w.Write([]byte(string(r)))
			}
		default:
			w.Write([]byte(string(r)))
		}
	}
	if stat != normal {
		panic(fmt.Errorf("no closed expression (exp: %s)", e.exp))
	}
	if len(e.values) != fieldIndex {
		panic(fmt.Errorf("extra field for express (fields: %v)", e.values))
	}
	w.Write([]byte(" "))
}

func (e *Expr) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(e.values) == 0 {
		w.Write([]byte(e.exp))
		w.Write([]byte(" "))
		return
	}
	reader := strings.NewReader(e.exp)
	stat := normal
	quoteStack := make([]rune, 0, 8)
	var fieldIndex int
	for r, _, err := reader.ReadRune(); ; r, _, err = reader.ReadRune() {
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		switch r {
		case '"', '`', '\'':
			if stat == normal {
				stat = inQuote
				quoteStack = append(quoteStack, r)
				w.Write([]byte(string(r)))
			} else {
				if quoteStack[len(quoteStack)-1] == r {
					quoteStack = quoteStack[:len(quoteStack)-1]
					if len(quoteStack) == 0 {
						stat = normal
					}

				} else {
					quoteStack = append(quoteStack, r)
					w.Write([]byte(string(r)))
				}
			}
		case '@':
			if stat == normal {
				switch c := e.values[fieldIndex].(type) {
				case referencer:
					c.toRefClause(w, vals, isFirstGroup, isFirstNode)
				default:
					panic(fmt.Errorf("invalid argument type (%T)", c))
				}
				fieldIndex++
			} else {
				w.Write([]byte(string(r)))
			}
		case '?':
			if stat == normal {
				switch c := e.values[fieldIndex].(type) {
				case *Expr:
					c.toClause(w, vals, isFirstGroup, isFirstNode)
				default:
					w.Write([]byte(valToPlaceholder(c)))
					*vals = append(*vals, expandArg(c)...)
				}
				fieldIndex++
			} else {
				w.Write([]byte(string(r)))
			}
		default:
			w.Write([]byte(string(r)))
		}
	}
	if stat != normal {
		panic(fmt.Errorf("no closed expression (exp: %s)", e.exp))
	}
	if len(e.values) != fieldIndex {
		panic(fmt.Errorf("extra field for express (fields: %v)", e.values))
	}
	w.Write([]byte(" "))
}

func (e *Expr) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(e.values) == 0 {
		w.Write([]byte(e.exp))
		w.Write([]byte(" "))
		return
	}
	reader := strings.NewReader(e.exp)
	stat := normal
	quoteStack := make([]rune, 0, 8)
	var fieldIndex int
	for r, _, err := reader.ReadRune(); ; r, _, err = reader.ReadRune() {
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		switch r {
		case '"', '`', '\'':
			if stat == normal {
				stat = inQuote
				quoteStack = append(quoteStack, r)
				w.Write([]byte(string(r)))
			} else {
				if quoteStack[len(quoteStack)-1] == r {
					quoteStack = quoteStack[:len(quoteStack)-1]
					if len(quoteStack) == 0 {
						stat = normal
					}
					w.Write([]byte(string(r)))
				} else {
					quoteStack = append(quoteStack, r)
					w.Write([]byte(string(r)))
				}
			}
		case '@':
			if stat == normal {
				switch c := e.values[fieldIndex].(type) {
				case referencer:
					c.toSimpleRefClause(w, vals, isFirstGroup, isFirstNode)
				default:
					panic(fmt.Errorf("invalid argument type (expr: %s, type: %T, values: %v)", e.exp, c, e.values))
				}
				fieldIndex++
			} else {
				w.Write([]byte(string(r)))
			}
		case '?':
			if stat == normal {
				switch c := e.values[fieldIndex].(type) {
				case *Expr:
					c.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
				default:
					w.Write([]byte(valToPlaceholder(c)))
					*vals = append(*vals, expandArg(c)...)
				}
				fieldIndex++
			} else {
				w.Write([]byte(string(r)))
			}
		default:
			w.Write([]byte(string(r)))
		}
	}
	if stat != normal {
		panic(fmt.Errorf("no closed expression (exp: %s)", e.exp))
	}
	if len(e.values) != fieldIndex {
		panic(fmt.Errorf("extra field for express (fields: %v)", e.values))
	}
	w.Write([]byte(" "))
}

// exprList 自定义sql表达式列表类型
type exprList []*Expr

// toClause 将sql自定义表达式列表转化为sql子句
func (l exprList) toClause(exprType exprType, w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	var clauseBuilder strings.Builder
	for _, exp := range l {
		exp.toClause(&clauseBuilder, vals, isFirstGroup, isFirstNode)
	}
	switch exprType {
	case whereExpr:
		w.Write([]byte("WHERE "))
		var builder strings.Builder
		for _, exp := range l {
			exp.toClause(&builder, vals, isFirstGroup, isFirstNode)
		}
		w.Write([]byte(trimPreAndOr(builder.String())))
		w.Write([]byte(" "))
	case assignExpr:
		for _, exp := range l {
			exp.toClause(w, vals, isFirstGroup, isFirstNode)
			w.Write([]byte(", "))
		}
	case havingExpr:
		w.Write([]byte("HAVING "))
		var builder strings.Builder
		for _, exp := range l {
			exp.toClause(&builder, vals, isFirstGroup, isFirstNode)
		}
		w.Write([]byte(trimPreAndOr(builder.String())))
		w.Write([]byte(" "))
	default:
		panic(fmt.Errorf("unknown expr type(%d)", exprType))
	}
}

// toSimpleClause 将sql自定义表达式转化为sql子句, 其中表名全部用原表名而不是别名
func (l exprList) toSimpleClause(exprType exprType, w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	if len(l) == 0 {
		return
	}
	var clauseBuilder strings.Builder
	for _, exp := range l {
		exp.toClause(&clauseBuilder, vals, isFirstGroup, isFirstNode)
	}
	switch exprType {
	case whereExpr:
		w.Write([]byte("WHERE "))
		var builder strings.Builder
		for _, exp := range l {
			exp.toSimpleClause(&builder, vals, isFirstGroup, isFirstNode)
		}
		w.Write([]byte(trimPreAndOr(builder.String())))
		w.Write([]byte(" "))
	case assignExpr:
		for _, exp := range l {
			exp.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
			w.Write([]byte(", "))
		}
	case havingExpr:
		w.Write([]byte("HAVING "))
		var builder strings.Builder
		for _, exp := range l {
			exp.toSimpleClause(&builder, vals, isFirstGroup, isFirstNode)
		}
		w.Write([]byte(trimPreAndOr(builder.String())))
		w.Write([]byte(" "))
	default:
		panic(fmt.Errorf("unknown expr type(%d)", exprType))
	}
}

// andGroup 将表达式列表包装成一个and group(例如: AND (tab1.field1 = 1 or tab.field2 = 2)), 此方法主要用来包装joinWheres, 以免其中的逻辑关系与其他wheres产生混淆
func (l exprList) andGroup() *Expr {
	cl := make([]string, 0, 16)
	vl := make([]interface{}, 0, 16)
	for _, e := range l {
		cl = append(cl, e.exp)
		vl = append(vl, e.values...)
	}
	return NewExpr(fmt.Sprintf("AND (%s)", trimPreAndOr(strings.Join(cl, " "))), vl...)
}

// orGroup 同andGroup, 仅将AND前缀替换成OR
func (l exprList) orGroup() *Expr {
	cl := make([]string, 0, 16)
	vl := make([]interface{}, 0, 16)
	for _, e := range l {
		cl = append(cl, strings.Trim(strings.Trim(e.exp, " "), "AND"))
		vl = append(vl, e.values...)
	}
	return NewExpr(fmt.Sprintf("OR (%s)", trimPreAndOr(strings.Join(cl, " "))), vl...)
}
