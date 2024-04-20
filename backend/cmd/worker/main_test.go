package main

import (
	"context"
	"distributed_calculator/internal/config"
	pb "distributed_calculator/internal/proto"
	"fmt"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestWorker(t *testing.T) {
	conn, err := grpc.Dial(config.StorageAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Error("failed connect to storage", err)
		os.Exit(1)
	}
	defer conn.Close()
	grpcClient := pb.NewStorageServiceClient(conn)
	res, err := grpcClient.CreateUser(context.TODO(), &pb.CreateUserRequest{
		Login: "user", Password: "1234",
	})
	fmt.Println(res, err)
}
