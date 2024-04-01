package worker

import (
	conf "distributed_calculator/internal/config"
	"errors"
)

var ErrOperationIsNotValid = errors.New("operation is not allowed")

type Worker struct {
	config *conf.Config
}

type Stack[T string | int] struct {
	Array []T
}

// Polish Notation converters
type PNConverter struct {
	Expression string
	stackOps   *Stack[string]
}
