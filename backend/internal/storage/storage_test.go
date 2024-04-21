package storage

import (
	"context"
	"testing"
)

func TestStorage(t *testing.T) {
	filename := "../../db/calculator.db"

	storage := NewStorage(filename)
	userID, err := storage.InsertUser(context.TODO(), "dima", "1234")
	if err == nil {
		_, err := storage.SelectUserByID(context.TODO(), int(userID))
		if err != nil {
			t.Errorf("failed to select user with ID: %d, err: %s", int(userID), err)
		}
	}

	id, err := storage.InsertExpression(context.TODO(), "1 + 2 + 3", int(userID))
	if err == nil {
		_, err := storage.SelectExpressionByID(context.TODO(), int(id))
		if err != nil {
			t.Errorf("failed to select expression with ID: %d, err: %s", int(id), err)
		}
	}
	_, err = storage.InsertExpression(context.TODO(), "1 + 2 + 3 + 4", int(userID))
	if err == nil {
		_, err := storage.SelectExpressionsByUserID(context.TODO(), int(userID))
		if err != nil {
			t.Errorf("faled to select expressions by user ID: %d, err: %s", int(userID), err)
		}
	}
}
