package nborm

type Expr string

func NewExpr(sqlexpr string) *Expr {
	expr := Expr(sqlexpr)
	return &expr
}

func (e *Expr) String() string {
	return string(*e)
}
