package main

import (
	"context"
	cfg "distributed_calculator/internal/config"
	pb "distributed_calculator/internal/proto"
	wk "distributed_calculator/internal/worker"
	"fmt"
	"log"
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
	config := cfg.NewConfig()
	worker := wk.NewWorker(config)
	worker.CalculateExpression(int(in.ExpressionID))

	return &pb.CalculateResponse{
		Recieved: true,
	}, nil
}

func main() {
	host := "localhost"
	port := "8000"
	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)
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
