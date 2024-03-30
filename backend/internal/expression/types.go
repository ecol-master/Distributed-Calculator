package expression

type Expression struct {
	ExpressionID   string
	Expression     string
	PolishNotation string
	Result         int
	Status         string
}

type Stack[T string | int] struct {
	Array []T
}
