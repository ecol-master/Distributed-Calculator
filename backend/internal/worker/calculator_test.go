package worker

import (
	"distributed_calculator/internal/config"
	"testing"
)

func TestCalculator(t *testing.T) {
	config := config.NewConfig()
	calculator := NewWorker(config)

	calculator.CalculateExpression(11)
}
