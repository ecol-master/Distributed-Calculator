package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	calc "distributed_calculator/pkg/calculator"
	conf "distributed_calculator/pkg/config"
	ex "distributed_calculator/pkg/expression"
	storage "distributed_calculator/pkg/storage"
)

var (
	config     = conf.NewConfig()
	calculator = calc.NewCalculator(config)

	syncStorageInterval = time.Second * 5
	appStorage          = storage.NewStorage("../data/data.json", syncStorageInterval)
)

type Response struct {
	StatusCode   int    // ответ 200, 400 или 500
	Result       int    // результат ответа (если сгенерировано)
	Status       string // текущий статус обработки выражения
	ExpressionID string // ID запроса
}

// handler listen "http://localhost:8000/new_expression?value={}&id={}"
func HandlerNewExpression(w http.ResponseWriter, r *http.Request) {
	var response Response

	exprValue := r.URL.Query().Get("value")
	expressionID := r.URL.Query().Get("id")

	expression, err := ex.NewExpression(exprValue, expressionID)
	appStorage.AddExpression(*expression)

	go calculator.CalculateExpression(appStorage, expression)
	if err != nil {
		response.StatusCode = 400
		response.ExpressionID = expressionID
	} else {
		response.StatusCode = 200
	}
	response.ExpressionID = expressionID

	responseBytes, err := json.Marshal(response)
	if err == nil {
		fmt.Fprint(w, string(responseBytes))
	}
}

func HandlerListExpressions(w http.ResponseWriter, r *http.Request) {
	result := []ex.Expression{}
	for _, v := range appStorage.Data() {
		result = append(result, v)
	}
	fmt.Println(result)
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Fprint(w, "[]")
	}
	fmt.Println(string(data))
	fmt.Fprint(w, string(data))
}

func HandlerGetOneExpression(w http.ResponseWriter, r *http.Request) {
	expressionID := r.URL.Query().Get("id")
	expression, found := appStorage.GetExpressionByID(expressionID)

	if !found {
		fmt.Fprint(w, "{Error: true}")
		return
	}

	response, err := json.Marshal(expression)
	if err != nil {
		fmt.Fprint(w, "{Error: true}")
		return
	}
	fmt.Fprint(w, string(response))
}

func HandlerGetConfig(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(config)
	if err != nil {
		fmt.Fprint(w, "{\"Error\": \"true\"}")
	}
	fmt.Println(string(response))
	fmt.Fprint(w, string(response))
}

func HandlePostConfig(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	// var newConfig conf.Config
	decoder.Decode(config)

	// fmt.Println(newConfig)
}
