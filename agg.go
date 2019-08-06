package nborm

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type agg struct {
	expr  *Expr
	name  string
	field Field
}

func (a *agg) MarshalJSON() ([]byte, error) {
	if !a.field.IsValid() || a.field.IsNull() {
		return []byte("null"), nil
	}
	buffer := bytes.NewBuffer(make([]byte, 0, 64))
	fmt.Fprintf(buffer, "{ \"%s\": ", a.field.colName())
	b, err := json.Marshal(a.field)
	if err != nil {
		return nil, err
	}
	buffer.Write(b)
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

func newStrAgg(expr *Expr, name string) *agg {
	f := new(String)
	f.Init(nil, name, "", -1)
	return &agg{expr, name, f}
}

func newIntAgg(expr *Expr, name string) *agg {
	f := new(Int)
	f.Init(nil, name, "", -1)
	return &agg{expr, name, f}
}

func newDateAgg(expr *Expr, name string) *agg {
	f := new(Date)
	f.Init(nil, name, "", -1)
	return &agg{expr, name, f}
}

func newDatetimeAgg(expr *Expr, name string) *agg {
	f := new(Datetime)
	f.Init(nil, name, "", -1)
	return &agg{expr, name, f}
}

func newDecAgg(expr *Expr, name string) *agg {
	f := new(Decimal)
	f.Init(nil, name, "", -1)
	return &agg{expr, name, f}
}

func (a *agg) toRefClause() string {
	return a.name
}

func (a *agg) toSimpleRefClause() string {
	return a.name
}

type aggList []*agg

func (l aggList) copy() aggList {
	nl := make(aggList, 0, 8)
	for _, agg := range l {
		switch agg.field.(type) {
		case *String:
			nl = append(nl, newStrAgg(agg.expr, agg.field.colName()))
		case *Int:
			nl = append(nl, newIntAgg(agg.expr, agg.field.colName()))
		case *Date:
			nl = append(nl, newDateAgg(agg.expr, agg.field.colName()))
		case *Datetime:
			nl = append(nl, newDatetimeAgg(agg.expr, agg.field.colName()))
		case *Decimal:
			nl = append(nl, newDecAgg(agg.expr, agg.field.colName()))
		default:
			panic(fmt.Errorf("unsupported field type (%T)", agg.field))
		}

	}
	return nl
}
