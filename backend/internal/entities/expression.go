package entities

import (
	pb "distributed_calculator/internal/proto"
	"strings"
)

const (
	StageCreated     = 0
	StageError       = 1
	StageCalculating = 2
	StageCalculated  = 3
)

type ExpressionID = int

type Expression struct {
	ID         ExpressionID `json:"id"`
	Expression string       `json:"expression"`
	UserID     int          `json:"user_id"`
	Result     int          `json:"result"`
	Stage      int          `json:"stage"`
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

func ConvertToTransport(expr Expression) *pb.Expression {
	return &pb.Expression{
		Id:         int32(expr.ID),
		Expression: expr.Expression,
		UserId:     int64(expr.UserID),
		Result:     int64(expr.Result),
		Stage:      int32(expr.Stage),
	}
}

func ConvertFromTransport(expr *pb.Expression) Expression {
	return Expression{
		ID:         int(expr.Id),
		Expression: expr.Expression,
		UserID:     int(expr.UserId),
		Result:     int(expr.Result),
		Stage:      int(expr.Stage),
	}
}
