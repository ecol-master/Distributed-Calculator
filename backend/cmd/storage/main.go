package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"context"
	"distributed_calculator/internal/config"
	"distributed_calculator/internal/expression"
	pb "distributed_calculator/internal/proto"
	"distributed_calculator/internal/storage"

	"google.golang.org/grpc"
)

type StorageServer struct {
	storage *storage.Storage
	pb.StorageServiceServer
}

func NewStorageServer() *StorageServer {
	return &StorageServer{
		storage: storage.NewStorage("../../db/calculator.db"),
	}
}

func (s *StorageServer) CreateExpression(ctx context.Context, in *pb.CreateExpressionRequest) (*pb.CreateExpressionResponse, error) {
	log.Println("invoking Create Expression")
	res, err := s.storage.InsertExpression(ctx, in.Expression, int(in.UserID))
	return &pb.CreateExpressionResponse{
		ExpressionID: int32(res),
	}, err
}

func (s *StorageServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Println("invoking Create User")
	res, err := s.storage.InsertUser(ctx, in.Login, in.Password)
	return &pb.CreateUserResponse{
		UserID: res,
	}, err
}

func (s *StorageServer) UpdateExpression(ctx context.Context, in *pb.UpdateExpressionRequest) (*pb.Empty, error) {
	e := expression.ConvertFromTransport(in.Expression)
	res := s.storage.UpdateExpression(ctx, e)
	return &pb.Empty{}, res
}

func (s *StorageServer) SelectUserExpressions(ctx context.Context, in *pb.SelectUserExpressionsRequest) (*pb.SelectUserExpressionsResponse, error) {
	res, err := s.storage.SelectExpressionsByUserID(ctx, int(in.UserID))
	if err != nil {
		return &pb.SelectUserExpressionsResponse{
			Expressions: nil,
		}, err
	}

	var expressions []*pb.Expression
	for _, e := range res {
		expressions = append(expressions, expression.ConvertToTransport(e))
	}
	return &pb.SelectUserExpressionsResponse{
		Expressions: expressions,
	}, nil
}

func (s *StorageServer) SelectExpression(ctx context.Context, in *pb.SelectExpressionRequest) (*pb.SelectExpressionResponse, error) {
	res, err := s.storage.SelectExpressionByID(ctx, int(in.ExpressionID))
	return &pb.SelectExpressionResponse{
		Expression: expression.ConvertToTransport(res),
	}, err
}

func main() {
	lis, err := net.Listen("tcp", config.StorageAddress)
	if err != nil {
		log.Println("error starting storage listener")
		os.Exit(1)
	}
	log.Println("started storage listener")
	grpcServer := grpc.NewServer()
	storageServiceServer := NewStorageServer()
	pb.RegisterStorageServiceServer(grpcServer, storageServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("error serving storage grpc: ", err)
		os.Exit(1)
	}
}
