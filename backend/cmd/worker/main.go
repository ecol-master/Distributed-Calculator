package main

import (
	"context"
	"distributed_calculator/internal/config"
	cfg "distributed_calculator/internal/config"
	pb "distributed_calculator/internal/proto"
	wk "distributed_calculator/internal/worker"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Server struct {
	pb.WorkerServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Calculate(ctx context.Context, in *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	conf := cfg.NewConfig()
	worker := wk.NewWorker(conf)
	worker.CalculateExpression(int(in.ExpressionID))

	return &pb.CalculateResponse{
		Recieved: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", config.WorkerAddress)
	if err != nil {
		log.Println("error starting worker tcp listener")
		os.Exit(1)
	}

	log.Println("started worker tcp listener")
	grpcServer := grpc.NewServer()
	workerServiceServer := NewServer()
	pb.RegisterWorkerServiceServer(grpcServer, workerServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving worker grpc: ", err)
		os.Exit(1)
	}
}
