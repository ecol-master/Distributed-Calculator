package main

import (
	"bytes"
	"distributed_calculator/internal/app"
	"distributed_calculator/internal/config"
	"encoding/json"
  "distributed_calculator/internal/logger"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestBase(t *testing.T) {
	client := &http.Client{}
	// creating User
	login, passwors := "admin", "1234"
	url := fmt.Sprintf("http://%s/new_user?login=%s&password=%s", config.ServerAddress, login, passwors)
	response, err := client.Get(url)
	if err != nil {
		t.Error(err)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	var respUser app.CreateUserResponse
	err = json.Unmarshal(data, &respUser)
	if err != nil {
		t.Error(err)
	}
	if respUser.StatusCode != app.StatusSuccessful {
		t.Error("status code is not successful")
		return
	}
	userID := respUser.UserID
	t.Log("created user with ID: ", userID)

	for i := 0; i < 5; i++ {
		expression := "11"
		url = fmt.Sprintf("http://%s/new_expression?value=%s&user_id=%d", config.ServerAddress, expression, userID)
		response, err = client.Get(url)
		if err != nil {
			t.Error(err)
		}
		data, err = io.ReadAll(response.Body)
		if err != nil {
			t.Error(err)
		}
		var respExpression app.CreateExpressionResponse
		err = json.Unmarshal(data, &respExpression)
		if err != nil {
			t.Error(err)
		}
		t.Log("created expression: ", respExpression)

		url = fmt.Sprintf("http://%s/get_expression?expression_id=%d", config.ServerAddress, respExpression.ExpressionID)
		response, err = client.Get(url)
		if err != nil {
			t.Error(err)
		}
		data, err = io.ReadAll(response.Body)
		if err != nil {
			t.Error(err)
		}

		var respGetExpression app.SelectExpressionResponse
		err = json.Unmarshal(data, &respGetExpression)
		if err != nil {
			t.Error(err)
		}
		t.Log("get expression with id: ", respExpression.ExpressionID)
	}

	url = fmt.Sprintf("http://%s/list_of_expressions?user_id=%d", config.ServerAddress, userID)

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("failed to fetch new expression: ", err)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	var respUserExpressions app.SelectUserExpressionsResponse
	err = json.Unmarshal(data, &respUserExpressions)
	if err != nil {
		t.Error(err)
	}
	t.Log("get userId: ", userID, " all expressions: ", respUserExpressions)
}

func TestCreateUser(t *testing.T) {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s/new_user", config.ServerAddress)
	data := []byte(`{"login":"admin", "password":"1234"}`)
	r := bytes.NewReader(data)
	response, err := client.Post(url, "application/json", r)
  if err != nil{
    logger.Error("failed while GET: /new_user", err) 
  }
  fmt.Println(response, err)
}
