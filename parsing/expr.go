package parsing

type Tokens int

const (
	IDEN = iota
	FUNC
	POW
	OPERATORS
)

type Funcs int

var ExprString string

func AcceptExpr(expr string) {

}
