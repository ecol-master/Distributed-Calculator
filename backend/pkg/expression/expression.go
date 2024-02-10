package expression

import (
	"errors"
	"fmt"
	"github.com/ecol-master/distributed_calculator/pkg/stack"
	"strconv"
	"strings"
)

var (
	ErrOperationIsNotValid = errors.New("opeation is not valid")
)

// Expression Statuses
const (
	StatusError       = "error"
	StatusCalculating = "calculating"
	StatusFinished    = "finished"
)

type Expression struct {
	ExpressionID   string
	Expression     string
	PolishNotation string
	Result         int
	Status         string
}

func GetOperationRank(operation string) (int, error) {
	if operation == "+" || operation == "-" {
		return 1, nil
	}
	if operation == "*" || operation == "/" {
		return 2, nil
	}
	if operation == "(" || operation == ")" {
		return 3, nil
	}
	return -1, ErrOperationIsNotValid
}

func NewExpression(expression string, expressionID string) (*Expression, error) {
	expr := &Expression{
		Expression:   expression,
		ExpressionID: expressionID,
	}

	polishNotation, err := convertExprToPN(expression)
	if err != nil {
		expr.Status = StatusError
		return expr, err
	}

	expr.PolishNotation = polishNotation
	expr.Status = StatusCalculating
	return expr, nil
}

func addHigherOpsToPN(stackOperations *stack.Stack[string], polishNotation, operation string) (string, *stack.Stack[string], error) {
	if operation == ")" {
		for {
			fmt.Println(stackOperations.Array)
			if len(stackOperations.Array) == 0 {
				break
			}
			if stackOperations.Array[len(stackOperations.Array)-1] == "(" {
				stackOperations.Pop()
				break
			}
			value, err := stackOperations.Pop()
			if err != nil {
				return polishNotation, stackOperations, err
			}
			polishNotation += value + " "
		}
		return polishNotation, stackOperations, nil
	}
	size := len(stackOperations.Array)
	operaionRank, err := GetOperationRank(operation)
	if err != nil {
		return polishNotation, stackOperations, err
	}
	for i := 0; i < size; i++ {
		lastOperaion, err := stackOperations.Pop()
		if err != nil {
			return polishNotation, stackOperations, err
		}
		lastOpRank, _ := GetOperationRank(lastOperaion)
		if err != nil {
			return polishNotation, stackOperations, err
		}

		if lastOpRank == 3 {
			stackOperations.Push(lastOperaion)
			break
		}

		if lastOpRank < operaionRank {
			stackOperations.Push(lastOperaion)
			break
		}
		polishNotation += lastOperaion + " "
	}
	return polishNotation, stackOperations, nil
}

// PN - polish notation
func convertExprToPN(expr string) (string, error) {
	fmt.Println(expr)
	stackOperations := stack.NewStack[string]()
	polishNotation := ""
	for _, value := range strings.Split(strings.Trim(expr, "\n ,.!?"), " ") {
		_, err := strconv.Atoi(value)
		if err == nil {
			polishNotation += value + " "
			continue
		}
		polishNotation, stackOperations, err = addHigherOpsToPN(stackOperations, polishNotation, value)
		if err != nil {
			return polishNotation, err
		}

		if value != ")" {
			stackOperations.Push(value)
		}
	}
	for i := len(stackOperations.Array) - 1; i >= 0; i-- {
		polishNotation += stackOperations.Array[i] + " "
	}
	return polishNotation, nil
}
