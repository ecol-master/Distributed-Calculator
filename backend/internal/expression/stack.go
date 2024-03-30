package expression

import (
	"errors"
)

var (
	ErrEmptyStack          = errors.New("stack is empty")
	ErrGetElementFromStack = errors.New("can't get new element from stack")
)

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
