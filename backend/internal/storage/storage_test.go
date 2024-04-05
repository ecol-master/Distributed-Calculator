package storage

import (
	"context"
	"fmt"
	"testing"
)

func TestStorage(t *testing.T) {
	filename := "../../db/calculator.db"

	storage := NewStorage(filename)
	userID, err := storage.InsertUser(context.TODO(), "dima", "1234")
	if err == nil {
		u, err := storage.SelectUserByID(context.TODO(), int(userID))
		fmt.Println("Selected user: ", u, err)
	}

	id, err := storage.InsertExpression(context.TODO(), "1 + 2 + 3", int(userID))
	if err == nil {
		e, err := storage.SelectExpressionByID(context.TODO(), int(id))
		fmt.Println("Selected expression: ", e, err)
	}
	_, err = storage.InsertExpression(context.TODO(), "1 + 2 + 3 + 4", int(userID))
	if err == nil {
		exps, err := storage.SelectExpressionsByUserID(context.TODO(), int(userID))
		fmt.Println("User expressions: ", exps, err)
	}
}
