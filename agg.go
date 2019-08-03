package nborm

import (
	"encoding/json"
	"fmt"
)

type agg struct {
	expr  *Expr
	field Field
}

func (a *agg) MarshalJSON() ([]byte, error) {
	if !a.field.IsValid() || a.field.IsNull() {
		return []byte("null"), nil
	}
	b, err := json.Marshal(a.field)
	if err != nil {
		return nil, err
	}
	return append([]byte(fmt.Sprintf("%s: ", a.field.colName())), b...), nil
}

func newStrAgg(expr *Expr, name string) *agg {
	f := new(String)
	f.Init(nil, name, "")
	return &agg{expr, f}
}

func newIntAgg(expr *Expr, name string) *agg {
	f := new(Int)
	f.Init(nil, name, "")
	return &agg{expr, f}
}

func newDateAgg(expr *Expr, name string) *agg {
	f := new(Date)
	f.Init(nil, name, "")
	return &agg{expr, f}
}

func newDatetimeAgg(expr *Expr, name string) *agg {
	f := new(Datetime)
	f.Init(nil, name, "")
	return &agg{expr, f}
}

func newDecAgg(expr *Expr, name string) *agg {
	f := new(Decimal)
	f.Init(nil, name, "")
	return &agg{expr, f}
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
