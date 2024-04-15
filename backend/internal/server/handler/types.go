package handler

import "distributed_calculator/internal/expression"

type StatusCode = int

const (
	StatusSuccessful  = StatusCode(200)
	StatusClientError = StatusCode(400)
	StatusServerError = StatusCode(500)
)

type createUserResponse struct {
	UserID       int        `json:"id"`
	StatusCode   StatusCode `json:"status_code"`
	ErrorMessage string     `json:"error_message"`
}

func NewCreateUserResponse(userID, statusCode int, message string) *createUserResponse {
	return &createUserResponse{
		UserID:       userID,
		StatusCode:   statusCode,
		ErrorMessage: message,
	}
}

type createExpressionResponse struct {
	ExpressionID int        `json:"id"`
	StatusCode   StatusCode `json:"status_code"`
	ErrorMessage string     `json:"error_message"`
}

func NewCreateExpressionResponse(expressionID, statusCode StatusCode, errorMessage string) *createExpressionResponse {
	return &createExpressionResponse{
		ExpressionID: expressionID,
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}
}

type selectExpressionResponse struct {
	expression.Expression
	StatusCode   StatusCode `json:"status_code"`
	ErrorMessage string     `json:"error_message"`
}

func NewSelectExpressionResponse(e expression.Expression, statusCode StatusCode, message string) *selectExpressionResponse {
	return &selectExpressionResponse{
		Expression:   e,
		StatusCode:   statusCode,
		ErrorMessage: message,
	}
}