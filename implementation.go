package lab2

import (
	"fmt"
	"math"
	"strings"
	"strconv"
)

type Stack []string

func (s *Stack) Pop() (string, bool) { //element, err
	if len(*s) == 0 {
		return "", true
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, false
	}
}

func evaluate(operator string, a float64, b float64) float64 {
	if operator == "+" {
		return a + b
	} else if operator == "-" {
		return a - b
	} else if operator == "*" {
		return a * b
	} else if operator == "/" {
		return a / b
	} else  {
		return math.Pow(a, b) // "^"
	}
}

func CalculatePrefix(input string) (string, error) {
	operators := "+-*/^"

	array := strings.Split(input, " ")
	var stack Stack

	for i := len(array) - 1; i >= 0; i-- {
		if !strings.Contains(operators, array[i]) {
			stack = append(stack, array[i])
		} else {
			stringElem1, err1 := stack.Pop()
			stringElem2, err2 := stack.Pop()
			if err1 || err2 {
				return "", fmt.Errorf("Not a prefix notation order")
			}
			floatElem1, _ := strconv.ParseFloat(stringElem1, 64)
			floatElem2, _ := strconv.ParseFloat(stringElem2, 64)

			floatResult := evaluate(array[i], floatElem1, floatElem2)
			stringResult := strconv.FormatFloat(floatResult, 'f', -1, 64)
			stack = append(stack, stringResult)
		}
	}

	finalResult, _ := stack.Pop()
	return finalResult, nil
}
