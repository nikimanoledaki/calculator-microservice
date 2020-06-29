package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nikimanoledaki/calculator-microservice/pkg/calculator"
)

func main() {
	operand, firstNumberString, secondNumberString := os.Args[1], os.Args[2], os.Args[3]

	firstNumber, _ := strconv.Atoi(firstNumberString)
	secondNumber, _ := strconv.Atoi(secondNumberString)

	result, err := calculator.Compute(operand, firstNumber, secondNumber)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d\n", result)
}
