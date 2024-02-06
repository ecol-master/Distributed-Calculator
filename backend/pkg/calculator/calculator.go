package calculator

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ecol-master/distributed_calculator/pkg/stack"
)

var (
	ErrOperationIsNotValid = errors.New("opeation is not valid")
)

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

type Calculator struct {
	// PolishNotation string
	// stackNumbers   *stack.Stack[int]
}

type Expression struct {
	ExpressionID   string
	PolishNotation string
	// stackOpeations *stack.Stack[string]
}

func NewCalculator() *Calculator {

	return &Calculator{
		// expression:     expression,
		// PolishNotation: "",
		// stackOpeations: stack.NewStack[string](),
		// stackNumbers: stack.NewStack[int](),
	}
}

func NewExpression(expression string, expressionID string) (*Expression, error) {
	polishNotation, err := convertExprToPN(expression)
	if err != nil {
		return &Expression{}, err
	}

	return &Expression{
		PolishNotation: polishNotation,
		ExpressionID:   expressionID,
	}, nil
}

func addHigherOpsToPN(stackOperations *stack.Stack[string], polishNotation, operation string) (string, *stack.Stack[string], error) {
	if operation == ")" {
		for len(stackOperations.Array) > 0 {
			if stackOperations.Array[len(stackOperations.Array)-1] == "(" {
				stackOperations.Pop()
				break
			}
			value, err := stackOperations.Pop()
			if err == nil {
				stackOperations.Push(value)
			}
		}
		return polishNotation, stackOperations, nil
	}
	size := len(stackOperations.Array)
	operaionRank, _ := GetOperationRank(operation)
	for i := 0; i < size; i++ {
		lastOperaion, err := stackOperations.Pop()
		if err != nil {
			break
		}
		lastOpRank, _ := GetOperationRank(lastOperaion)
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
	stackOperations := stack.NewStack[string]()
	polishNotation := ""
	for _, value := range strings.Split(strings.Trim(expr, "\n ,.!?"), " ") {
		_, err := strconv.Atoi(value)
		if err == nil {
			polishNotation += value + " "
			continue
		}
		polishNotation, stackOperations, _ = addHigherOpsToPN(stackOperations, polishNotation, value)

		if value != ")" {
			stackOperations.Push(value)
		}
	}
	for i := len(stackOperations.Array) - 1; i >= 0; i-- {
		polishNotation += stackOperations.Array[i] + " "
	}
	return polishNotation, nil
}

func (c *Calculator) CalculateExpression(expr *Expression) int {
	stackNumbers := stack.NewStack[int]()

	for _, value := range strings.Split(strings.Trim(expr.PolishNotation, " \n"), " ") {
		number, err := strconv.Atoi(value)
		if err == nil {
			stackNumbers.Push(number)
			continue
		}
		if value == "+" {
			stack.Sum(stackNumbers)
		} else if value == "-" {
			stack.Diff(stackNumbers)
		} else if value == "*" {
			stack.Multiply(stackNumbers)
		} else if value == "/" {
			stack.Devide(stackNumbers)
		}
	}

	res, err := stackNumbers.Pop()
	if err != nil {
		return -1
	}
	return res
}
