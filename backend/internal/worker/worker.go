package worker

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"distributed_calculator/internal/config"
	"distributed_calculator/internal/entities"
	"distributed_calculator/internal/logger"

	pb "distributed_calculator/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New() (*Worker, error) {
	connStorage, err := grpc.Dial(config.StorageAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &Worker{}, err
	}

	grpcStorageClient := pb.NewStorageServiceClient(connStorage)
	return &Worker{
		Config:        config.NewConfig(),
		StorageClient: grpcStorageClient,
	}, nil
}

func (w *Worker) updateExpression(expression *pb.Expression) {
	_, err := w.StorageClient.UpdateExpression(context.TODO(), &pb.UpdateExpressionRequest{
		Expression: expression,
	})
	if err != nil {
		logger.Error("failed to update expression, err: ", err.Error())
	}
}

func (w *Worker) calculate(e *pb.Expression) error {
	stackNumbers := NewStack[int]()
	logger.Info("starting converting exprssion: ", e.Expression)
	polishNotation, err := NewConverter(e.Expression).Convert()

	if err != nil {
		logger.Error("failed to convert expression to polish notation, id: ", e.Id, " err: ", err.Error())
		// TODO: update error in expression
		e.Stage = entities.StageError
		w.updateExpression(e)
		return err
	}

	e.Stage = entities.StageCalculating
	w.updateExpression(e)

	for _, value := range strings.Split(strings.Trim(polishNotation, " \n"), " ") {
		number, err := strconv.Atoi(value)
		if err == nil {
			stackNumbers.Push(number)
			continue
		}
		if value == "+" {
			Sum(stackNumbers)
			time.Sleep(w.Config.SumDelay)
		} else if value == "-" {
			Diff(stackNumbers)
			time.Sleep(w.Config.DiffDelay)
		} else if value == "*" {
			Multiply(stackNumbers)
			time.Sleep(w.Config.MultiplyDelay)
		} else if value == "/" {
			Devide(stackNumbers)
			time.Sleep(w.Config.DevideDelay)
		}
	}

	result, err := stackNumbers.Pop()
	if err != nil {
		logger.Error("failed to calculating result")
		e.Stage = entities.StageError
	} else {
		e.Result = int64(result)
		e.Stage = entities.StageCalculated
	}

	w.updateExpression(e)
	log.Println("result of expression: ", result, err)

	return nil
}

func (w *Worker) CalculateExpression(expressionID int) {
	logger.Info("start calculating expreession with ID: ", expressionID)

	res, err := w.StorageClient.SelectExpression(context.TODO(), &pb.SelectExpressionRequest{
		ExpressionID: int32(expressionID),
	})

	if err != nil {
		logger.Error("failed select expression with id: ", expressionID, " err: ", err.Error())
		return
	}

	err = w.calculate(res.Expression)
	if err != nil {
		logger.Error("calculate expression with error: " + err.Error())
	}
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
