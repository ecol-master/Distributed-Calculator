package worker

import (
	"log"
	"strconv"
	"strings"
	"time"

	conf "distributed_calculator/internal/config"
	"distributed_calculator/internal/expression"
)

func NewWorker(cfg *conf.Config) *Worker {
	return &Worker{
		config: cfg,
	}
}

func fetchExpression(expressionID int) (expression.Expression, error) {
	testExpression := expression.NewExpression("1 + 2 * 5", 11)
	return *testExpression, nil
}

func (c *Worker) CalculateExpression(expressionID int) {
	// log.Printf("start calculating expression with with value=\"%s\" and id=\"%s\"", expr.Expression, expr.ExpressionID)
	expr, _ := fetchExpression(expressionID)

	stackNumbers := NewStack[int]()
	converter := NewConverter(expr.Expression)
	polishNotation, err := converter.Convert()

	if err != nil {
		// adding--
	}

	for _, value := range strings.Split(strings.Trim(polishNotation, " \n"), " ") {
		number, err := strconv.Atoi(value)
		if err == nil {
			stackNumbers.Push(number)
			continue
		}
		if value == "+" {
			Sum(stackNumbers)
			time.Sleep(c.config.SumDelay)
		} else if value == "-" {
			Diff(stackNumbers)
			time.Sleep(c.config.DiffDelay)
		} else if value == "*" {
			Multiply(stackNumbers)
			time.Sleep(c.config.MultiplyDelay)
		} else if value == "/" {
			Devide(stackNumbers)
			time.Sleep(c.config.DevideDelay)
		}
	}

	res, err := stackNumbers.Pop()
	log.Printf("result of expression: %d", res)

	// appStorage.Update(*expr)
	// log.Printf("finished calculating expression with with value=\"%s\" and id=\"%s\"", expr.Expression, expr.ExpressionID)
}

func GetTwoValues(s *Stack[int]) (int, int, error) {
	n1, err1 := s.Pop()
	n2, err2 := s.Pop()
	if err1 != nil || err2 != nil {
		return 0, 0, ErrGetElementFromStack
	}
	return n1, n2, nil
}

func Sum(s *Stack[int]) {
	n1, n2, err := GetTwoValues(s)
	if err == nil {
		s.Push(n2 + n1)
	}
}

func Diff(s *Stack[int]) {
	n1, n2, err := GetTwoValues(s)
	if err == nil {
		s.Push(n2 - n1)
	}
}

func Multiply(s *Stack[int]) {
	n1, n2, err := GetTwoValues(s)
	if err == nil {
		s.Push(n2 * n1)
	}
}

func Devide(s *Stack[int]) {
	n1, n2, err := GetTwoValues(s)
	if err == nil {
		s.Push(n2 / n1)
	}
}
