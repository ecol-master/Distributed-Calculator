package main

import (
	"testing"

	"distributed_calculator/internal/config"
	pb "distributed_calculator/internal/proto"

	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestUpdate(t *testing.T) {
	conn, err := grpc.Dial(config.StorageAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Error("failed to connect to storage: ", err)
		return
	}

	grpcStorageClient := pb.NewStorageServiceClient(conn)

	expression, exprectedResult := "1 * 5 * 10", 50
	resp, err := grpcStorageClient.CreateExpression(context.TODO(), &pb.CreateExpressionRequest{
		Expression: expression,
		UserID:     1,
	})

	if err != nil {
		t.Error("failed to create new expression: ", err)
		return
	}
	expressionID := resp.ExpressionID
	resp2, err := grpcStorageClient.SelectExpression(context.TODO(), &pb.SelectExpressionRequest{
		ExpressionID: expressionID,
	})

	if err != nil {
		t.Error("failed to select expressionfrom storage: ", err)
		return
	}

	e := resp2.Expression
	e.Result = int64(exprectedResult)

	_, err = grpcStorageClient.UpdateExpression(context.TODO(), &pb.UpdateExpressionRequest{
		Expression: e,
	})

	if err != nil {
		t.Error("failed to update expression", err)
		return
	}

	resp2, err = grpcStorageClient.SelectExpression(context.TODO(), &pb.SelectExpressionRequest{
		ExpressionID: expressionID,
	})

	if err != nil {
		t.Error("failed to select updated expression")
		return
	}

	if resp2.Expression.Result != int64(exprectedResult) {
		t.Errorf("want result value: %d, but got: %d", exprectedResult, resp2.Expression.Result)
	}
}
