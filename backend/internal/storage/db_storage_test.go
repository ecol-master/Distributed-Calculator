package storage

import (
	"distributed_calculator/internal/expression"
	"testing"
)

func TestDBStorage(t *testing.T) {
	filename := "../../db/calculator.db"

	storage := NewDBStorage(filename)
	expr, err := expression.NewExpression("1 + 2", "123dfklj")
	if err != nil {
		t.Fatalf("test insert expression failed with error: %v", err)
	}
	storage.insertExpression(expr)
}
