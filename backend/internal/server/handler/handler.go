package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	conf "distributed_calculator/internal/config"

	calc "distributed_calculator/internal/calculator"
	ex "distributed_calculator/internal/expression"
	storage "distributed_calculator/internal/storage"
)

var (
	config     = conf.NewConfig()
	calculator = calc.NewCalculator(config)

	syncStorageInterval = time.Second * 30
	appStorage          = storage.NewStorage("../data/data.json", syncStorageInterval)
)

type Response struct {
	StatusCode   int    // ответ 200, 400 или 500
	ExpressionID string // ID запроса
}

// handler listen "http://localhost:8000/new_expression?value={}&id={}"
// v2 handler listen "http://localhost:8000/new_expression?value={}"
func HandlerNewExpression(w http.ResponseWriter, r *http.Request) {
	exprValue := r.URL.Query().Get("value")
	log.Printf("handler /new_expression with value=\"%s\"", exprValue)

	expression, err := ex.NewExpression(exprValue, expressionID)
	appStorage.AddExpression(*expression)

	go calculator.CalculateExpression(appStorage, expression)

	response := Response{ExpressionID: expressionID}
	if err != nil {
		response.StatusCode = 400
	} else {
		response.StatusCode = 200
	}

	responseBytes, err := json.Marshal(response)

	if err == nil {
		fmt.Fprint(w, string(responseBytes))
	}
	log.Printf("processed /new_expression with value=\"%s\" and status=%d", expression.Expression, response.StatusCode)
}

func HandlerListExpressions(w http.ResponseWriter, r *http.Request) {
	log.Printf("handler /list_of_expressions")
	result := []ex.Expression{}
	for _, v := range appStorage.Data() {
		result = append(result, v)
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Fprint(w, "[]")
		log.Printf("processed /list_of_expressions with error=%s", err)
		return
	}

	fmt.Fprint(w, string(data))
	log.Printf("processed /list_of_expressions with %d expression", len(result))
}

func HandlerGetOneExpression(w http.ResponseWriter, r *http.Request) {
	expressionID := r.URL.Query().Get("id")
	log.Printf("handler /get_expression expressionID=\"%s\"", expressionID)

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
	log.Printf("processed /get_expression expressionID=\"%s\"", expressionID)
}

func HandlerGetConfig(w http.ResponseWriter, r *http.Request) {
	log.Printf("handler /get_config ")

	response, err := json.Marshal(config)
	if err != nil {
		fmt.Fprint(w, "{\"Error\": \"true\"}")
	}

	fmt.Fprint(w, string(response))
	log.Printf("processed /get_config send config settings=\"%s\"", config.AsString())
}

func HandlePostConfig(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	// create newConfig to check correct new values for config settings
	var newConfig conf.Config
	decoder.Decode(&newConfig)
	config.CopySettings(newConfig)
	log.Printf("processes post /set_config, config_value=\"%s\"", config.AsString())
}
