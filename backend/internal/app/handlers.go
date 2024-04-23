package app

import (
	"context"
	"distributed_calculator/internal/config"
	"distributed_calculator/internal/entities"
	"distributed_calculator/internal/logger"
	pb "distributed_calculator/internal/proto"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
		os.Exit(1)
	}
	grpcStorageClient = pb.NewStorageServiceClient(connStorage)

	// grpc connection to worker
	connWorker, err := grpc.Dial(config.WorkerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("failed connect to worker", err)
		os.Exit(1)
	}
	grpcWorkerClient = pb.NewWorkerServiceClient(connWorker)

	logger.Info("connections to storage and worker successful initialize")
}

// HandlerNewUser "http://localhost:8080/new_user?login={}&password={}"
func HandlerNewUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("invoke /new_user handler")

	var user struct {
		Password string `json:"password"`
		Login    string `json:"login"`
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error whlie parsing form: ", err)
	}
	password, login := user.Password, user.Login

	logger.Info("new user password: ", password, " login: ", login)

	var bytes []byte
	if (strings.TrimSpace(login) == "") || (strings.TrimSpace(password) == "") {
		bytes = marshalJSONResponse(NewCreateUserResponse(0, StatusClientError, "login and password can not be empty"))
		fmt.Fprint(w, string(bytes))
		return
	}

	res, err := grpcStorageClient.CreateUser(context.TODO(), &pb.CreateUserRequest{
		Login: login, Password: password,
	})
	if err != nil {
		logger.Error("server error while creating user" + err.Error())
		bytes = marshalJSONResponse(NewCreateUserResponse(0, StatusServerError, "server error while creating user"))
	} else {
		bytes = marshalJSONResponse(NewCreateUserResponse(int(res.UserID), StatusSuccessful, ""))
	}
	fmt.Fprint(w, string(bytes))
}

// HandlerNewExpression "http://localhost:8080/new_expression?value={}&id={}"
// v2 "http://localhost:8080/new_expression?value={}&user_id={}"
func HandlerNewExpression(w http.ResponseWriter, r *http.Request) {
	logger.Info("invoke /new_expression handler")

	var expression struct {
		Value  string `json:"value"`
		UserID int    `json:"user_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&expression)
	if err != nil {
		http.Error(w, "error while parsing expression form: "+err.Error(), 500)
		return
	}

	res, err := grpcStorageClient.CreateExpression(context.TODO(), &pb.CreateExpressionRequest{
		Expression: expression.Value,
		UserID:     int32(expression.UserID),
	})

	var bytes []byte
	if err != nil {
		bytes = marshalJSONResponse(NewCreateExpressionResponse(0, StatusServerError, "server error while creating expression"))
		logger.Error("failed to create expression: ", err.Error())
		fmt.Fprint(w, string(bytes))
		return
	}
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
		bytes = marshalJSONResponse(NewSelectExpressionResponse(entities.Expression{}, StatusClientError, "value expression id should be integer"))
		fmt.Fprint(w, string(bytes))
		return
	}

	res, err := grpcStorageClient.SelectExpression(context.TODO(), &pb.SelectExpressionRequest{ExpressionID: int32(expressionID)})
	if err != nil {
		logger.Error("error while selecting expression with ID: ", expressionID, " err: ", err.Error())
		msg := "server error"
		bytes = marshalJSONResponse(NewSelectExpressionResponse(entities.Expression{}, StatusServerError, msg))
	} else {
		e := entities.ConvertFromTransport(res.Expression)
		bytes = marshalJSONResponse(NewSelectExpressionResponse(e, StatusSuccessful, ""))
	}
	fmt.Fprint(w, string(bytes))
}

// HandlerSelectUserExpressions http://localhost:8080/list_of_expressions?user_id={}
func HandlerSelectUserExpressions(w http.ResponseWriter, r *http.Request) {
	logger.Info("invoke /list_of_expressions")

	var bytes []byte
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		bytes = marshalJSONResponse(NewSelectUserExpressionsResponse([]entities.Expression{}, StatusClientError, "user_id mush be an interger"))
		fmt.Fprint(w, string(bytes))
		return
	}

	res, err := grpcStorageClient.SelectUserExpressions(context.TODO(), &pb.SelectUserExpressionsRequest{
		UserID: int32(userID),
	})

	if err != nil {
		logger.Error("can not select user expressions: ", err.Error())
		bytes = marshalJSONResponse(NewSelectUserExpressionsResponse([]entities.Expression{}, StatusServerError, "can not select user expressions"))
	} else {
		var exs []entities.Expression
		for _, e := range res.Expressions {
			exs = append(exs, entities.ConvertFromTransport(e))
		}
		bytes = marshalJSONResponse(NewSelectUserExpressionsResponse(exs, StatusSuccessful, ""))
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
