package main

import (
	"distributed_calculator/internal/server"
	"fmt"
	"log"
	"os"
)

// initialize application

func main() {
	setupLog()
	log.Printf("App started")
	err := server.Run()
	fmt.Println(err)
}

func setupLog() {
	logFile, err := os.OpenFile("info.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile)
	flags := log.Lshortfile | log.Ldate | log.Ltime
	log.SetFlags(flags)
}