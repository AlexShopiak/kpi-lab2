package lab2

import (
	"fmt"
	"math"
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

func PrefixToPostfix(input string) (string, error) {
	return "TODO", fmt.Errorf("TODO")
}
