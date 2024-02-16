package main

import (
	"distributed_calculator/http"
	"log"
	"os"
)

func main() {
	setupLog()
	log.Printf("App started")
	http.Run()
}

func setupLog() {
	logFile, err := os.OpenFile("../data/info.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		os.Exit(1)
	}

	log.SetOutput(logFile)
	flags := log.Lshortfile | log.Ldate | log.Ltime
	log.SetFlags(flags)
}
