package expression

import (
	"strings"
)

const (
	StageError       = 0
	StageCalculating = 1
	StageCalculated  = 2
)

type Expression struct {
	Expression string
	UserID     int
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
