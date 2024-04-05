package worker

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "distributed_calculator/internal/proto"
)

func TestWorker(t *testing.T) {
	workerHost := "localhost"
	workerPort := "8000"
	addr := fmt.Sprintf("%s:%s", workerHost, workerPort)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("could not connect to grpc server: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	workerClient := pb.NewWorkerServiceClient(conn)
	recieved, err := workerClient.Calculate(context.TODO(), &pb.CalculateRequest{
		ExpressionID: 13,
	})
	if err != nil {
		log.Println("failed invoking Calculate: ", err)
	}
	log.Println("grpc server result: ", recieved)
}
