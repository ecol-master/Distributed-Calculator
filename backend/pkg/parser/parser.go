package parser

import (
	"strconv"
	"strings"

	"github.com/ecol-master/distributed_calculator/pkg/stack"
)

func GetOperationRank(operation string) int {
	switch operation {
	case "+":
		return 1
	case "-":
		return 1
	case "*":
		return 2
	case "/":
		return 2
	default:
		return 3
	}
}

type Expression struct {
	expression     string
	polishNotation string
	stackOpeations *stack.Stack[string]
	stackNumbers   *stack.Stack[int]
}

func NewExpression(expression string) *Expression {
	return &Expression{
		expression:     expression,
		polishNotation: "",
		stackOpeations: stack.NewStack[string](),
		stackNumbers:   stack.NewStack[int](),
	}
}

func (e *Expression) AddHigherOpsToExp(operation string) {
	if operation == ")" {
		for len(e.stackOpeations.Array) > 0 {
			if e.stackOpeations.Array[len(e.stackOpeations.Array)-1] == "(" {
				e.stackOpeations.Pop()
				break
			}
			value, err := e.stackOpeations.Pop()
			if err == nil {
				e.stackOpeations.Push(value)
			}
		}
		return
	} else {
		size := len(e.stackOpeations.Array)
		operaionRank := GetOperationRank(operation)
		for i := 0; i < size; i++ {
			lastOperaion, err := e.stackOpeations.Pop()
			if err != nil {
				break
			}
			lastOpRank := GetOperationRank(lastOperaion)
			if lastOpRank == 3 {
				e.stackOpeations.Push(lastOperaion)
				break
			}

			if lastOpRank >= operaionRank {
				e.polishNotation += lastOperaion + " "
			} else {
				e.stackOpeations.Push(lastOperaion)
				break
			}
		}
	}
}

func (e *Expression) ParseExpression() string {
	for _, value := range strings.Split(strings.Trim(e.expression, "\n ,.!?"), " ") {
		_, err := strconv.Atoi(value)
		if err == nil {
			// то есть числоL
			e.polishNotation += value + " "
			continue
		}
		// теперь рассматриваем строку
		e.AddHigherOpsToExp(value)
		if value != ")" {
			e.stackOpeations.Push(value)
		}
	}
	for _, op := range e.stackOpeations.Array {
		e.polishNotation += op + " "
	}
	return e.polishNotation
}

func (e *Expression) CalculateExpression() int {

	for _, value := range strings.Split(strings.Trim(e.polishNotation, " \n"), " ") {
		number, err := strconv.Atoi(value)
		if err == nil {
			e.stackNumbers.Push(number)
			continue
		}
		if value == "+" {
			stack.Sum(e.stackNumbers)
		} else if value == "-" {
			stack.Diff(e.stackNumbers)
		} else if value == "*" {
			stack.Multiply(e.stackNumbers)
		} else if value == "/" {
			stack.Devide(e.stackNumbers)
		}
	}

	res, err := e.stackNumbers.Pop()
	if err != nil {
		return -1
	}
	return res
}
