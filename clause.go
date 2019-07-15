package nborm

import "fmt"

type operator string

const (
	eq    operator = "="
	neq   operator = "<>"
	lt    operator = "<"
	lte   operator = "<="
	gt    operator = ">"
	gte   operator = ">="
	is    operator = "IS NULL"
	nis   operator = "IS NOT NULL"
	like  operator = "LIKE"
	nlike operator = "NOT LIKE"
	in    operator = "IN"
	nin   operator = "NOT IN"
)

func checkOp(op string) {
	switch op {
	case "=", "<>", "<", "<=", ">", ">=", "IS NULL", "IS NOT NULL", "LIKE", "NOT LIKE", "IN", "NOT IN":
		return
	default:
		panic(fmt.Sprintf("unsupported operator (%s)", op))
	}
}
