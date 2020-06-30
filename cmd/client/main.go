package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nikimanoledaki/calculator-microservice/pkg/calculator"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("number of arguments is not valid")
	}

	operation, numbersAsString := os.Args[1], os.Args[2:]

	numbers := make([]int, len(numbersAsString))
	for i, arg := range numbersAsString {
		var err error
		numbers[i], err = strconv.Atoi(arg)
		if err != nil {
			fmt.Println("last arguments must be numbers")
		}
	}

	result, err := calculator.Compute(operation, numbers[0], numbers[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d\n", result)
}
