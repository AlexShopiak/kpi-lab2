package lab2

import "fmt"


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

func PrefixToPostfix(input string) (string, error) {
	return "TODO", fmt.Errorf("TODO")
}
