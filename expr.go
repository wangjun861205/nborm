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
func (e *Expr) toClause() (string, []interface{}) {
	if len(e.values) == 0 {
		return e.exp, nil
	}
	reader := strings.NewReader(e.exp)
	stat := normal
	quoteStack := make([]rune, 0, 8)
	values := make([]interface{}, 0, len(e.values))
	var fieldIndex int
	var builder strings.Builder
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
				builder.WriteRune(r)
			} else {
				if quoteStack[len(quoteStack)-1] == r {
					quoteStack = quoteStack[:len(quoteStack)-1]
					if len(quoteStack) == 0 {
						stat = normal
					}
					builder.WriteRune(r)
				} else {
					quoteStack = append(quoteStack, r)
					builder.WriteRune(r)
				}
			}
		case '@':
			if stat == normal {
				switch c := e.values[fieldIndex].(type) {
				case clauser:
					cs, cv := c.toClause()
					builder.WriteString(cs)
					values = append(values, cv...)
				case refClauser:
					cs := c.toRefClause()
					builder.WriteString(cs)
				default:
					panic(fmt.Errorf("invalid argument type (%T)", c))
				}
				// builder.WriteString(e.values[fieldIndex].(ValueField).fullColName())
				fieldIndex++
			} else {
				builder.WriteRune(r)
			}
		case '?':
			if stat == normal {
				switch c := e.values[fieldIndex].(type) {
				case *Expr:
					cs, vs := c.toClause()
					builder.WriteString(cs)
					values = append(values, vs...)
				default:
					builder.WriteString(valToPlaceholder(c))
					values = append(values, expandArg(c)...)
				}
				// builder.WriteString(valToPlaceholder(e.values[fieldIndex]))
				// values = append(values, expandArg(e.values[fieldIndex])...)
				fieldIndex++
			} else {
				builder.WriteRune(r)
			}
		default:
			builder.WriteRune(r)
		}
	}
	if stat != normal {
		panic(fmt.Errorf("no closed expression (exp: %s)", e.exp))
	}
	if len(e.values) != fieldIndex {
		panic(fmt.Errorf("extra field for express (fields: %v)", e.values))
	}
	return builder.String(), values
}

// toSimpleClause 自定义表达式转化为sql子句，其中涉及到表名的全部用表的原名而不是用别名
func (e *Expr) toSimpleClause() (string, []interface{}) {
	if len(e.values) == 0 {
		return e.exp, nil
	}
	reader := strings.NewReader(e.exp)
	stat := normal
	quoteStack := make([]rune, 0, 8)
	values := make([]interface{}, 0, len(e.values))
	var fieldIndex int
	var builder strings.Builder
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
				builder.WriteRune(r)
			} else {
				if quoteStack[len(quoteStack)-1] == r {
					quoteStack = quoteStack[:len(quoteStack)-1]
					if len(quoteStack) == 0 {
						stat = normal
					}
					builder.WriteRune(r)
				} else {
					quoteStack = append(quoteStack, r)
					builder.WriteRune(r)
				}
			}
		case '@':
			if stat == normal {
				switch c := e.values[fieldIndex].(type) {
				case clauser:
					cs, cv := c.toSimpleClause()
					builder.WriteString(cs)
					values = append(values, cv...)
				case refClauser:
					cs := c.toSimpleRefClause()
					builder.WriteString(cs)
				default:
					panic(fmt.Errorf("invalid argument type (expr: %s, type: %T, values: %v)", e.exp, c, e.values))
				}
				// builder.WriteString(e.values[fieldIndex].(Field).rawFullColName())
				// values = append(values, expandArg(e.values[fieldIndex])...)
				fieldIndex++
			} else {
				builder.WriteRune(r)
			}
		case '?':
			if stat == normal {
				switch c := e.values[fieldIndex].(type) {
				case *Expr:
					cs, vs := c.toSimpleClause()
					builder.WriteString(cs)
					values = append(values, vs...)
				default:
					builder.WriteString(valToPlaceholder(c))
					values = append(values, expandArg(c)...)
				}
				// builder.WriteString(valToPlaceholder(e.values[fieldIndex]))
				// values = append(values, expandArg(e.values[fieldIndex])...)
				fieldIndex++
			} else {
				builder.WriteRune(r)
			}
		default:
			builder.WriteRune(r)
		}
	}
	if stat != normal {
		panic(fmt.Errorf("no closed expression (exp: %s)", e.exp))
	}
	if len(e.values) != fieldIndex {
		panic(fmt.Errorf("extra field for express (fields: %v)", e.values))
	}
	return builder.String(), values
}

func (e *Expr) toRefClause() string {
	c, _ := e.toClause()
	return c
}

func (e *Expr) toSimpleRefClause() string {
	c, _ := e.toSimpleClause()
	return c
}

// exprList 自定义sql表达式列表类型
type exprList []*Expr

// toClause 将sql自定义表达式列表转化为sql子句
func (l exprList) toClause(exprType exprType) (string, []interface{}) {
	if len(l) == 0 {
		return "", nil
	}
	cl := make([]string, 0, len(l))
	vl := make([]interface{}, 0, len(l)*2)
	for _, exp := range l {
		c, vs := exp.toClause()
		cl = append(cl, c)
		vl = append(vl, vs...)
	}
	switch exprType {
	case whereExpr:
		return fmt.Sprintf("WHERE %s", trimPreAndOr(strings.Join(cl, " "))), vl
	case assignExpr:
		return strings.Join(cl, ", "), vl
	case havingExpr:
		return fmt.Sprintf("HAVING %s", trimPreAndOr(strings.Join(cl, " "))), vl
	default:
		panic(fmt.Errorf("unknown expr type(%d)", exprType))
	}
}

// toSimpleClause 将sql自定义表达式转化为sql子句, 其中表名全部用原表名而不是别名
func (l exprList) toSimpleClause(exprType exprType) (string, []interface{}) {
	if len(l) == 0 {
		return "", nil
	}
	cl := make([]string, 0, len(l))
	vl := make([]interface{}, 0, len(l)*2)
	for _, exp := range l {
		c, vs := exp.toSimpleClause()
		cl = append(cl, c)
		vl = append(vl, vs...)
	}
	switch exprType {
	case whereExpr:
		return fmt.Sprintf("WHERE %s", trimPreAndOr(strings.Join(cl, " "))), vl
	case assignExpr:
		return strings.Join(cl, ", "), vl
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
