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
	updateExpr
)

type Expr struct {
	exp    string
	values []interface{}
}

func NewExpr(sqlexpr string, values ...interface{}) *Expr {
	return &Expr{
		sqlexpr,
		values,
	}
}

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
				builder.WriteString(e.values[fieldIndex].(Field).fullColName())
				// values = append(values, expandArg(e.values[fieldIndex])...)
				fieldIndex++
			} else {
				builder.WriteRune(r)
			}
		case '?':
			if stat == normal {
				builder.WriteString(valToPlaceholder(e.values[fieldIndex]))
				values = append(values, expandArg(e.values[fieldIndex])...)
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
				builder.WriteString(e.values[fieldIndex].(Field).fullColName())
				values = append(values, expandArg(e.values[fieldIndex])...)
				fieldIndex++
			} else {
				builder.WriteRune(r)
			}
		case '?':
			if stat == normal {
				builder.WriteString(valToPlaceholder(e.values[fieldIndex]))
				values = append(values, expandArg(e.values[fieldIndex])...)
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

type exprList []*Expr

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
		// return fmt.Sprintf("WHERE %s", strings.TrimPrefix(strings.TrimPrefix(strings.Join(cl, " "), " "), "AND")), vl
		return fmt.Sprintf("WHERE %s", trimPreAndOr(strings.Join(cl, " "))), vl
	case updateExpr:
		return strings.Join(cl, ", "), vl
	default:
		panic(fmt.Errorf("unknown expr type(%d)", exprType))
	}
}

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
		// return fmt.Sprintf("WHERE %s", strings.TrimPrefix(strings.TrimPrefix(strings.Join(cl, " "), " "), "AND")), vl
		return fmt.Sprintf("WHERE %s", trimPreAndOr(strings.Join(cl, " "))), vl
	case updateExpr:
		return strings.Join(cl, ", "), vl
	default:
		panic(fmt.Errorf("unknown expr type(%d)", exprType))
	}
}

func (l exprList) andGroup() *Expr {
	cl := make([]string, 0, 16)
	vl := make([]interface{}, 0, 16)
	for _, e := range l {
		cl = append(cl, e.exp)
		vl = append(vl, e.values...)
	}
	return NewExpr(fmt.Sprintf("AND (%s)", trimPreAndOr(strings.Join(cl, " "))), vl...)
}

func (l exprList) orGroup() *Expr {
	cl := make([]string, 0, 16)
	vl := make([]interface{}, 0, 16)
	for _, e := range l {
		cl = append(cl, strings.Trim(strings.Trim(e.exp, " "), "AND"))
		vl = append(vl, e.values...)
	}
	return NewExpr(fmt.Sprintf("OR (%s)", trimPreAndOr(strings.Join(cl, " "))), vl...)
}
