package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	calc "github.com/ecol-master/distributed_calculator/pkg/calculator"
)

var calculator calc.Calculator

type Response struct {
	StatusCode   int    // ответ 200, 400 или 500
	Result       int    // результат ответа (если сгенерировано)
	Status       string // текущий статус обработки выражения
	ExpressionID string // ID запроса, который сейчас обрабатываем
}

// func NewResponse(exp *calculator.Calculator) ([]byte, error) {
// 	var response Response
// 	err := exp.ParseExpression()
// 	fmt.Println(exp.PolishNotation)
// 	if err != nil {
// 		response.StatusCode = 400
// 		return json.Marshal(response)
// 	}
// 	response.StatusCode = 200
// 	response.Result = exp.CalculateExpression()
// 	return json.Marshal(response)
// }

// handler listen "http://localhost:8000/new_expression?value="
func HandlerNewExpression(w http.ResponseWriter, r *http.Request) {

	var response Response

	// w.Header().Set("Content-Type", "application/json")
	exprValue := r.URL.Query().Get("value")
	exprValue = strings.Replace(exprValue, "PP", "+", -1)

	expressionID := r.URL.Query().Get("id")
	expression, err := calc.NewExpression(exprValue, expressionID)
	if err != nil {
		response.StatusCode = 400
		response.ExpressionID = expressionID
	} else {
		response.StatusCode = 200
		result := calculator.CalculateExpression(expression)
		response.Result = result
	}
	// response, err := NewResponse(expression)
	responseBytes, err := json.Marshal(response)
	if err == nil {
		fmt.Fprint(w, string(responseBytes))
	}
}
