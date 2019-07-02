package nborm

type AggResult []FieldList

func (r AggResult) ToList() []map[string]interface{} {
	l := make([]map[string]interface{}, 0, 64)
	for _, rec := range r {
		m := make(map[string]interface{})
		for _, f := range rec {
			m[f.fullColName()] = f.Value()
		}
		l = append(l, m)
	}
	return l
}

func (r AggResult) ToSimpleList() []map[string]interface{} {
	l := make([]map[string]interface{}, 0, 64)
	for _, rec := range r {
		m := make(map[string]interface{})
		for _, f := range rec {
			m[f.colName()] = f.Value()
		}
		l = append(l, m)
	}
	return l
}

type aggExp struct {
	expr  *Expr
	field Field
}

func newStrAgg(expr *Expr, name string) *aggExp {
	f := new(String)
	f.col = name
	return &aggExp{expr, f}
}

func newIntAgg(expr *Expr, name string) *aggExp {
	f := new(Int)
	f.col = name
	return &aggExp{expr, f}
}

func newDateAgg(expr *Expr, name string) *aggExp {
	f := new(Date)
	f.col = name
	return &aggExp{expr, f}
}

func newDatetimeAgg(expr *Expr, name string) *aggExp {
	f := new(Datetime)
	f.col = name
	return &aggExp{expr, f}
}

func newDecAgg(expr *Expr, name string) *aggExp {
	f := new(Datetime)
	f.col = name
	return &aggExp{expr, f}
}
