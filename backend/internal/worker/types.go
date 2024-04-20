package worker

import (
	conf "distributed_calculator/internal/config"
	pb "distributed_calculator/internal/proto"
	"errors"
)

var ErrOperationIsNotValid = errors.New("operation is not allowed")

type Worker struct {
	Config        *conf.Config
	StorageClient pb.StorageServiceClient
}

type Stack[T string | int] struct {
	Array []T
}

// Polish Notation converters
type PNConverter struct {
	Expression string
	stackOps   *Stack[string]
}
