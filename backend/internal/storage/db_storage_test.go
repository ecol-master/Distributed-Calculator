package storage

import (
	"distributed_calculator/internal/expression"
	"testing"
)

func TestDBStorage(t *testing.T) {
	filename := "../../db/calculator.db"

	storage := NewDBStorage(filename)
	expr := expression.NewExpression("1 + 2", 10)
	storage.insertExpression(expr)
}
