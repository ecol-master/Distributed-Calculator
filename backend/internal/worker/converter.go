package worker

import (
	"strconv"
	"strings"
)

func NewConverter(expression string) *PNConverter {

	return &PNConverter{
		Expression: expression,
		stackOps:   NewStack[string](),
	}
}

func (c *PNConverter) addHigherOpsToPN(polishNotation, operation string) (string, error) {
	if operation == ")" {
		for {
			if len(c.stackOps.Array) == 0 {
				break
			}
			if c.stackOps.Array[len(c.stackOps.Array)-1] == "(" {
				c.stackOps.Pop()
				break
			}
			value, err := c.stackOps.Pop()
			if err != nil {
				return polishNotation, err
			}
			polishNotation += value + " "
		}
		return polishNotation, nil
	}
	size := len(c.stackOps.Array)
	operaionRank, err := c.getOperationRank(operation)
	if err != nil {
		return polishNotation, err
	}
	for i := 0; i < size; i++ {
		lastOperaion, err := c.stackOps.Pop()
		if err != nil {
			return polishNotation, err
		}
		lastOpRank, err := c.getOperationRank(lastOperaion)
		if err != nil {
			return polishNotation, err
		}

		if lastOpRank == 3 {
			c.stackOps.Push(lastOperaion)
			break
		}

		if lastOpRank < operaionRank {
			c.stackOps.Push(lastOperaion)
			break
		}
		polishNotation += lastOperaion + " "
	}
	return polishNotation, nil
}

// PN - polish notation
func (c *PNConverter) Convert() (string, error) {
	stackOperations := NewStack[string]()
	polishNotation := ""

	for _, value := range strings.Split(strings.Trim(c.Expression, "\n ,.!?"), " ") {
		_, err := strconv.Atoi(value)
		if err == nil {
			polishNotation += value + " "
			continue
		}
		polishNotation, err = c.addHigherOpsToPN(polishNotation, value)
		if err != nil {
			return polishNotation, err
		}

		if value != ")" {
			stackOperations.Push(value)
		}
	}
	for i := len(stackOperations.Array) - 1; i >= 0; i-- {
		polishNotation += stackOperations.Array[i] + " "
	}
	return polishNotation, nil
}

// function return arithmetic operation rank
// "+  -" - 1
// "*  /" - 2
// "(  )" - 3
func (c *PNConverter) getOperationRank(operation string) (int, error) {
	if operation == "+" || operation == "-" {
		return 1, nil
	}
	if operation == "*" || operation == "/" {
		return 2, nil
	}
	if operation == "(" || operation == ")" {
		return 3, nil
	}
	return -1, ErrOperationIsNotValid
}
