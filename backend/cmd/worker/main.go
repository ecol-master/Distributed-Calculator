package main

import (
	"context"
	"distributed_calculator/internal/config"
	"distributed_calculator/internal/logger"
	pb "distributed_calculator/internal/proto"
	"distributed_calculator/internal/worker"
	"net"
	"os"

	"google.golang.org/grpc"
)

type Server struct {
	pb.WorkerServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Calculate(ctx context.Context, in *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	worker, err := worker.New()
	if err != nil {
		return &pb.CalculateResponse{}, err
	}

	// starting goroutine to calculate new expression
	go worker.CalculateExpression(int(in.ExpressionID))

	return &pb.CalculateResponse{
		Recieved: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", config.WorkerPort)
	if err != nil {
		logger.Error("error starting worker tcp listener, err: ", err)
		os.Exit(1)
	}

	logger.Info("started worker tcp listener")
	grpcServer := grpc.NewServer()
	workerServiceServer := NewServer()
	pb.RegisterWorkerServiceServer(grpcServer, workerServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("error serving worker grpc: ", err)
		os.Exit(1)
	}
}
