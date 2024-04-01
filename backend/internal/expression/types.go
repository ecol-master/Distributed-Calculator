package expression

type Expression struct {
	ExpressionID   string
	Expression     string
	PolishNotation string
	Result         int
	Status         string
}

// const (
// 	StageError       = 0
// 	StageCalculating = 1
// 	StageCalculated  = 2
// )

// type Expression struct {
// 	Expression string
// 	UserID     int
// 	Result     int
// 	Stage      int
// }

type Stack[T string | int] struct {
	Array []T
}
