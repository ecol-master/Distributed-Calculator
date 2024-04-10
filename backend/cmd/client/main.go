package main

import (
	"context"
	pb "distributed_calculator/internal/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {
	host := "localhost"
	port := "8000"

	addr := fmt.Sprintf("%s:%s", host, port) // используем адрес сервера
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("could not connect to grpc server: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	grpcClient := pb.NewWorkerServiceClient(conn)
	result, err := grpcClient.Calculate(context.TODO(), &pb.CalculateRequest{ExpressionID: 1})
	fmt.Println(result, err)
}
