package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	ch "github.com/ecol-master/distributed_calculator/pkg/cache"
	calc "github.com/ecol-master/distributed_calculator/pkg/calculator"
	conf "github.com/ecol-master/distributed_calculator/pkg/config"
	ex "github.com/ecol-master/distributed_calculator/pkg/expression"
)

var (
	config     = conf.NewConfig()
	calculator = calc.NewCalculator(config)
	cache      = ch.NewCache()
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

	exprValue := parseExpressionValue(r)
	expressionID := r.URL.Query().Get("id")

	expression, err := ex.NewExpression(exprValue, expressionID)
	cache.AddExpression(expression)

	go calculator.CalculateExpression(cache, expression)
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

// функция форматирует данные из запроса в арифметическое выражение
func parseExpressionValue(r *http.Request) string {
	exprValue := r.URL.Query().Get("value")
	exprValue = strings.Replace(exprValue, "PP", "+", -1)
	exprValue = strings.Replace(exprValue, "BO", "(", -1)
	exprValue = strings.Replace(exprValue, "BC", ")", -1)
	return exprValue
}

func HandlerListExpressions(w http.ResponseWriter, r *http.Request) {
	result := []ex.Expression{}
	for _, v := range cache.GetAllData() {
		result = append(result, *v)
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
	expression, found := cache.GetExpressionByID(expressionID)

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
