package main

import (
	"distributed_calculator/internal/app"
	"distributed_calculator/internal/logger"

	"os"
)

func main() {
	application, err := app.New()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	err = application.Run()
	if err != nil {
		logger.Error("failed to read config")
	}
}
