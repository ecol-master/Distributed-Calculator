package calculator

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	ex "github.com/ecol-master/distributed_calculator/pkg/expression"
)

type CalculatorConcurrency struct{}

func (c *CalculatorConcurrency) CalculateExpression(expr *ex.Expression) int {
	numbers := make(chan int, 1)
	operations := make(chan string, 1)
	defer close(numbers)
	defer close(operations)

	fmt.Println(strings.Split(expr.PolishNotation, " "))
	for _, value := range strings.Split(expr.PolishNotation, " ") {
		fmt.Printf("value: %s\n", value)
		_, err := strconv.Atoi(value)
		if err != nil {
			operations <- value
			fmt.Println("here")
			// switch value {
			// case "+":
			// 	Sum(numbers, operations)
			// case "-":
			// 	Diff(numbers, operations)
			// case "*":
			// 	Mutiply(numbers, operations)
			// case "/":
			// 	Devide(numbers, operations)
			// }
		} else {
			fmt.Println("HERE")
			// numbers <- number
		}
	}
	return 1
	// fmt.Println(numbers)

	// select {
	// case result := <-numbers:
	// 	return result
	// default:
	// 	return -10
	// }
}

func SumConc(numbers chan int, operations <-chan string) {
	go func() {
		fmt.Printf("Numbers::")
		for n := range numbers {
			fmt.Printf("%d ", n)
		}
		n1 := <-numbers
		n2 := <-numbers
		time.Sleep(3 * time.Second)
		numbers <- n1 + n2
	}()
}

func DiffConc(numbers chan int, operations <-chan string) {
	n1 := <-numbers
	n2 := <-numbers
	time.Sleep(3 * time.Second)
	numbers <- n2 - n1

}

func DevideConc(numbers chan int, operations <-chan string) {
	n1 := <-numbers
	n2 := <-numbers
	time.Sleep(3 * time.Second)
	numbers <- n2 / n1
}

func MutiplyConc(numbers chan int, operations <-chan string) {
	n1 := <-numbers
	n2 := <-numbers
	time.Sleep(3 * time.Second)
	numbers <- n1 * n2

}
