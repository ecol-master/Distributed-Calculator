package handler

import (
	"context"
	"distributed_calculator/internal/config"
	pb "distributed_calculator/internal/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
)

// var (
//
//	config     = conf.NewConfig()
//	calculator = calc.NewCalculator(config)
//
//	syncStorageInterval = time.Second * 30
//	appStorage          = storage.NewStorage("../data/data.json", syncStorageInterval)
//
// )

var (
	grpcStorageClient pb.StorageServiceClient
	grpcWorkerClient  pb.WorkerServiceClient
)

func init() {
	log.Printf("initialize connections")
	// grpc connection to storage
	connStorage, err := grpc.Dial(config.StorageAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("failed connect to storage")
		os.Exit(1)
	}
	grpcStorageClient = pb.NewStorageServiceClient(connStorage)

	//resp, err := grpcStorageClient.CreateUser(context.TODO(), &pb.CreateUserRequest{
	//	Login:    "u",
	//	Password: "1",
	//})
	//fmt.Println(resp, err)

	// grpc connection to worker
	connWorker, err := grpc.Dial(config.WorkerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("failed connect to worker")
		os.Exit(1)
	}
	grpcWorkerClient = pb.NewWorkerServiceClient(connWorker)
}

type Response struct {
	StatusCode   int    // ответ 200, 400 или 500
	ExpressionID string // ID запроса
}

// handler listen "http://localhost:8000/new_expression?value={}&id={}"
// v2 handler listen "http://localhost:8000/new_expression?value={}"
func HandlerNewExpression(w http.ResponseWriter, r *http.Request) {
	exprValue := r.URL.Query().Get("value")

	log.Printf("handler /new_expression with value=\"%s\"", exprValue)
	result, err := grpcStorageClient.CreateUser(context.TODO(), &pb.CreateUserRequest{
		Login:    "u",
		Password: "1",
	})
	fmt.Println(result, err)

	res, err := grpcStorageClient.CreateExpression(context.TODO(), &pb.CreateExpressionRequest{
		Expression: exprValue,
		UserID:     10,
	})
	if err != nil {
		log.Printf("failed to create expression: %v", err)
	}
	log.Println("result of creating new expression: ", res)
}
