package nborm

import (
	"fmt"
	"io"
	"strings"
)

type expStat int

const (
	normal  expStat = 0
	inQuote expStat = 1
)

type Expr struct {
	exp    string
	fields []Field
}

func NewExpr(sqlexpr string, fields ...Field) *Expr {
	for _, f := range fields {
		f.addModelStatus(forModelRef)
	}
	return &Expr{
		sqlexpr,
		fields,
	}
}

func (e *Expr) String() string {
	if len(e.fields) == 0 {
		return e.exp
	}
	reader := strings.NewReader(e.exp)
	stat := normal
	quoteStack := make([]rune, 0, 8)
	var fieldIndex int
	var builder strings.Builder
	for r, _, err := reader.ReadRune(); ; r, _, err = reader.ReadRune() {
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if r == '"' || r == '`' || r == '\'' {
			if stat == normal {
				stat = inQuote
				quoteStack = append(quoteStack, r)
			} else {
				if quoteStack[len(quoteStack)-1] == r {
					quoteStack = quoteStack[:len(quoteStack)-1]
					if len(quoteStack) == 0 {
						stat = normal
					}
				} else {
					quoteStack = append(quoteStack, r)
				}
			}
		}
		if r == '@' {
			if stat == normal {
				builder.WriteString(e.fields[fieldIndex].fullColName())
				fieldIndex++
			} else {
				builder.WriteRune(r)
			}
		} else {
			builder.WriteRune(r)
		}
	}
	if stat != normal {
		panic(fmt.Errorf("no closed expression (exp: %s)", e.exp))
	}
	if len(e.fields) != 0 {
		panic(fmt.Errorf("extra field for express (fields: %v)", e.fields))
	}
	return builder.String()
}

func (e *Expr) SimpleString() string {
	if len(e.fields) == 0 {
		return e.exp
	}
	reader := strings.NewReader(e.exp)
	stat := normal
	quoteStack := make([]rune, 0, 8)
	var builder strings.Builder
	for r, _, err := reader.ReadRune(); ; r, _, err = reader.ReadRune() {
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if r == '"' || r == '`' || r == '\'' {
			if stat == normal {
				stat = inQuote
			} else {
				if quoteStack[len(quoteStack)-1] == r {
					quoteStack = quoteStack[:len(quoteStack)-1]
					if len(quoteStack) == 0 {
						stat = normal
					}
				} else {
					quoteStack = append(quoteStack, r)
				}
			}
		}
		if r == '@' {
			if stat == normal {
				builder.WriteString(e.fields[0].rawFullColName())
				e.fields = e.fields[1:]
			} else {
				builder.WriteRune(r)
			}
		} else {
			builder.WriteRune(r)
		}
	}
	if stat != normal {
		panic(fmt.Errorf("no closed expression (exp: %s)", e.exp))
	}
	if len(e.fields) != 0 {
		panic(fmt.Errorf("extra field for express (fields: %v)", e.fields))
	}
	return builder.String()
}
