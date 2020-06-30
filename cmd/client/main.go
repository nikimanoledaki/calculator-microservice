package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("number of arguments is not valid")
	}

	numbersAsString := os.Args[1:]

	numbers := make([]int, len(numbersAsString))
	for i, arg := range numbersAsString {
		var err error
		numbers[i], err = strconv.Atoi(arg)
		if err != nil {
			fmt.Println("last arguments must be numbers")
		}
	}

	// result, err := calculator.Sum(numbers[0], numbers[1])
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%d\n", result)
}
