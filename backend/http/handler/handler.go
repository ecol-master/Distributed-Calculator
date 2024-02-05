package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ecol-master/distributed_calculator/pkg/parser"
)

type Response struct {
	StatusCode   int    // ответ 200, 400 или 500
	Result       int    // результат ответа (если сгенерировано)
	Status       string // текущий статус обработки выражения
	ExpressionID int    // ID запроса, который сейчас обрабатываем
}

func NewResponse(exp *parser.Expression) ([]byte, error) {
	var response Response
	err := exp.ParseExpression()
	if err != nil {
		response.StatusCode = 400
		return json.Marshal(response)
	}
	response.StatusCode = 200
	response.Result = exp.CalculateExpression()
	return json.Marshal(response)
}

func HandlerNewExpression(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	exprValue := r.URL.Query().Get("value")
	exprValue = strings.Replace(exprValue, "PP", "+", -1)
	expression := parser.NewExpression(exprValue)
	response, err := NewResponse(expression)
	if err == nil {
		fmt.Fprint(w, string(response))
	}
}
