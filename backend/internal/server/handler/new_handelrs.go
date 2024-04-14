package handler

import (
	"context"
	"distributed_calculator/internal/config"
	"distributed_calculator/internal/expression"
	"distributed_calculator/internal/logger"
	pb "distributed_calculator/internal/proto"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcStorageClient pb.StorageServiceClient
	grpcWorkerClient  pb.WorkerServiceClient
)

// funcion initialize connections to worker and storage
func init() {
	connStorage, err := grpc.Dial(config.StorageAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("failed connect to storage", err)
		panic("failed connect to storage")
	}
	grpcStorageClient = pb.NewStorageServiceClient(connStorage)

	// grpc connection to worker
	connWorker, err := grpc.Dial(config.WorkerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("failed connect to worker", err)
		panic("failed connect to worker")
	}
	grpcWorkerClient = pb.NewWorkerServiceClient(connWorker)

	logger.Info("connections to storage and worker successful initialize")
}

// HandlerNewUser "http://localhost:8080/new_user?login={},password={}"
func HandlerNewUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("invoke /new_user handler")

	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")

	var bytes []byte
	if strings.TrimSpace(login) == "" || strings.TrimSpace(password) == "" {
		bytes = marshalJSONResponse(NewCreateUserResponse(0, StatusClientError, "login and password can not be empty"))
		fmt.Fprint(w, string(bytes))
		return
	}

	res, err := grpcStorageClient.CreateUser(context.TODO(), &pb.CreateUserRequest{
		Login: login, Password: password,
	})
	if err != nil {
		bytes = marshalJSONResponse(NewCreateUserResponse(0, StatusServerError, "server error while creating user"))
	} else {
		bytes = marshalJSONResponse(NewCreateUserResponse(int(res.UserID), StatusSuccessful, ""))
	}
	fmt.Fprint(w, string(bytes))
}

// HandlerNewExpression "http://localhost:8080/new_expression?value={}&id={}"
// v2 "http://localhost:8080/new_expression?value={},user_id={}"
func HandlerNewExpression(w http.ResponseWriter, r *http.Request) {
	logger.Info("invoke /new_expression handler")

	exprValue := r.URL.Query().Get("value")
	userIDValue := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDValue)

	var bytes []byte
	if err != nil {
		bytes = marshalJSONResponse(NewCreateExpressionResponse(0, StatusClientError, "userId should be an integer"))
		fmt.Fprint(w, string(bytes))
	}

	res, err := grpcStorageClient.CreateExpression(context.TODO(), &pb.CreateExpressionRequest{
		Expression: exprValue,
		UserID:     int32(userID),
	})

	if err != nil {
		bytes = marshalJSONResponse(NewCreateExpressionResponse(0, StatusServerError, "server error while creating expression"))
		logger.Info("failed to create expression: ", err.Error())
		fmt.Fprint(w, string(bytes))
		return
	}

	logger.Info("result of creating new expression: ", res)

	resp, err := grpcWorkerClient.Calculate(context.TODO(), &pb.CalculateRequest{
		ExpressionID: res.ExpressionID,
	})
	if err != nil || !resp.Recieved {
		bytes = marshalJSONResponse(NewCreateExpressionResponse(0, StatusClientError, "error while send expression to worker"))
	} else {
		bytes = marshalJSONResponse(NewCreateExpressionResponse(int(res.ExpressionID), StatusSuccessful, ""))
	}
	fmt.Fprint(w, string(bytes))
}

// HnalderGetOneExpression "http://localhost:8080/get_expression?expression_id={}"
func HandlerSelectExpression(w http.ResponseWriter, r *http.Request) {
	logger.Info("invoke /get_expression handler")

	var bytes []byte
	expressionID, err := strconv.Atoi(r.URL.Query().Get("expression_id"))
	if err != nil {
		bytes = marshalJSONResponse(NewSelectExpressionResponse(expression.Expression{}, StatusClientError, "value expression id should be integer"))
		fmt.Fprint(w, string(bytes))
		return
	}

	res, err := grpcStorageClient.SelectExpression(context.TODO(), &pb.SelectExpressionRequest{ExpressionID: int32(expressionID)})
	if err != nil {
		msg := "error while selecting expression with ID: " + string(expressionID)
		bytes = marshalJSONResponse(NewSelectExpressionResponse(expression.Expression{}, StatusServerError, msg))
	} else {
		e := expression.ConvertFromTransport(res.Expression)
		bytes = marshalJSONResponse(NewSelectExpressionResponse(e, StatusSuccessful, ""))
	}
	fmt.Fprint(w, string(bytes))
}

func marshalJSONResponse(response interface{}) []byte {
	bytes, err := json.Marshal(response)
	if err != nil {
		panic("JSON marshal error: " + err.Error())
	}
	return bytes
}
