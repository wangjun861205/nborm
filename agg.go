package nborm

// AggResult 统计函数结果container
type AggResult []FieldList

// ToList 将统计结果转化为方便阅读和传递的格式
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

// ToSimpleList 与ToList类似,但map的key只用field的名称，不包含表的alias前缀，故无法保证无冲突，如果有相同的字段名，后者会覆盖前者的值
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
