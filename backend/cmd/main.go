package main

import (
	"fmt"

	"github.com/ecol-master/distributed_calculator/pkg/parser"
)

func main() {
	expression := parser.NewExpression("2 + 2 * 2 + 2")
	plolish := expression.ParseExpression()
	fmt.Println("polish: ", plolish)
	fmt.Println("result: ", expression.CalculateExpression())
}
