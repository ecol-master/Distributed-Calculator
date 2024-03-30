package main

import (
	"distributed_calculator/internal/server"
	"log"
	"os"
)

func main() {
	setupLog()
	log.Printf("App started")

	server.Run()
}

func setupLog() {
	logFile, err := os.OpenFile("../data/info.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile)
	flags := log.Lshortfile | log.Ldate | log.Ltime
	log.SetFlags(flags)
}
