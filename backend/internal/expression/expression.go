package expression

import (
	"strings"
)

const (
	StageCreated     = 0
	StageError       = 1
	StageCalculating = 2
	StageCalculated  = 3
)

type Expression struct {
	ID         int
	Expression string
	UserID     int
	Result     int
	Stage      int
}

func parseExpressionValue(exprValue string) string {
	exprValue = strings.Replace(exprValue, "PP", "+", -1)
	exprValue = strings.Replace(exprValue, "BO", "(", -1)
	exprValue = strings.Replace(exprValue, "BC", ")", -1)
	return exprValue
}

func NewExpression(exprValue string, userID int) *Expression {
	expression := parseExpressionValue(exprValue)
	return &Expression{
		Expression: parseExpressionValue(expression),
		UserID:     userID,
	}
}
