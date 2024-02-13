package calculator

import (
	"strconv"
	"strings"
	"time"

	"distributed_calculator/pkg/cache"
	conf "distributed_calculator/pkg/config"
	ex "distributed_calculator/pkg/expression"
	"distributed_calculator/pkg/stack"
)

type Calculator struct {
	config *conf.Config
}

func NewCalculator(cfg *conf.Config) *Calculator {
	return &Calculator{
		config: cfg,
	}
}

func (c *Calculator) CalculateExpression(ch *cache.Cache, expr *ex.Expression) {

	stackNumbers := stack.NewStack[int]()

	for _, value := range strings.Split(strings.Trim(expr.PolishNotation, " \n"), " ") {
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
	if err != nil {
		expr.Status = ex.StatusError
	} else {
		expr.Result = res
		expr.Status = ex.StatusFinished
	}

	ch.Update(expr)
}

func GetTwoValues(s *stack.Stack[int]) (int, int, error) {
	n1, err1 := s.Pop()
	n2, err2 := s.Pop()
	if err1 != nil || err2 != nil {
		return 0, 0, stack.ErrGetElementFromStack
	}
	return n1, n2, nil
}

func Sum(s *stack.Stack[int]) {
	n1, n2, err := GetTwoValues(s)
	if err == nil {
		s.Push(n2 + n1)
	}
}

func Diff(s *stack.Stack[int]) {
	n1, n2, err := GetTwoValues(s)
	if err == nil {
		s.Push(n2 - n1)
	}
}

func Multiply(s *stack.Stack[int]) {
	n1, n2, err := GetTwoValues(s)
	if err == nil {
		s.Push(n2 * n1)
	}
}

func Devide(s *stack.Stack[int]) {
	n1, n2, err := GetTwoValues(s)
	if err == nil {
		s.Push(n2 / n1)
	}
}
