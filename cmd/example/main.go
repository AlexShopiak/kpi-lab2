package main

import (
	"flag"
	"fmt"
	lab2 "github.com/AlexShopiak/kpi-lab2"
)

var (
	inputFromStdin = flag.String("e", "", "Expression from stdin")
	inputFromFile = flag.String("f", "", "Expression from file")
	outputToFile = flag.String("o", "", "Result to file")
)

func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	//res, _ := lab2.PrefixToPostfix("+ 2 2")
	//fmt.Println(res)
}
