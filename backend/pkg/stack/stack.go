package stack

import (
	"errors"
)

var (
	ErrEmptyStack          = errors.New("stack is empty")
	ErrGetElementFromStack = errors.New("can't get new element from stack")
)

type Stack[T string | int] struct {
	Array []T
}

func NewStack[T string | int]() *Stack[T] {
	return &Stack[T]{
		Array: make([]T, 0),
	}
}

func (s *Stack[T]) Pop() (T, error) {
	var retval T
	if len(s.Array) == 0 {
		return retval, ErrEmptyStack
	}
	retval = s.Array[len(s.Array)-1]
	s.Array = s.Array[:len(s.Array)-1]
	return retval, nil
}

func (s *Stack[T]) Push(value T) {
	s.Array = append(s.Array, value)
}

// Arithmetic operations

func GetTwoValues(stack *Stack[int]) (int, int, error) {
	n1, err1 := stack.Pop()
	n2, err2 := stack.Pop()
	if err1 != nil || err2 != nil {
		return 0, 0, ErrGetElementFromStack
	}
	return n1, n2, nil
}

func Sum(stack *Stack[int]) {
	n1, n2, err := GetTwoValues(stack)
	if err == nil {
		stack.Push(n2 + n1)
	}
}

func Diff(stack *Stack[int]) {
	n1, n2, err := GetTwoValues(stack)
	if err == nil {
		stack.Push(n2 - n1)
	}
}

func Multiply(stack *Stack[int]) {
	n1, n2, err := GetTwoValues(stack)
	if err == nil {
		stack.Push(n2 * n1)
	}
}

func Devide(stack *Stack[int]) {
	n1, n2, err := GetTwoValues(stack)
	if err == nil {
		stack.Push(n2 / n1)
	}
}
