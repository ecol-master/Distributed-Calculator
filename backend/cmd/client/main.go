package main

import (
	"distributed_calculator/internal/config"
	"fmt"
	"net/http"
)

func main() {
	var expression string
	fmt.Printf("Type your exression to calculate: ")
	fmt.Scanf("%s", &expression)
	client := &http.Client{}
	url := fmt.Sprintf("http://%s/new_expression?value=%s", config.ServerAddress, expression)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("failed to fetch new expression: ", err)
	}
	fmt.Println(resp)
}
