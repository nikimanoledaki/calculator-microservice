package calculator

import (
	"errors"
	"fmt"
)

// Compute performs the arithmetic operation on two given numbers based on a given operand.
func Compute(operand string, firstNumber, secondNumber int) (int, error) {
	switch operand {
	case "sum":
		return Sum(firstNumber, secondNumber), nil
	case "subtract":
		return Subtract(firstNumber, secondNumber), nil
	case "multiply":
		return Multiply(firstNumber, secondNumber), nil
	case "divide":
		return Divide(firstNumber, secondNumber), nil
	}
	return 0, errors.New("operand could not be found\n")
}

// Sum returns the sum of two numbers.
func Sum(firstNumber, secondNumber int) int {
	return firstNumber + secondNumber
}

// Subtract  returns the difference between two numbers.
func Subtract(firstNumber, secondNumber int) int {
	return firstNumber - secondNumber
}

// Multiply  returns the product of two numbers.
func Multiply(firstNumber, secondNumber int) int {
	return firstNumber * secondNumber
}

// Divide  returns the quotient of two numbers.
func Divide(firstNumber, secondNumber int) int {
	if secondNumber == 0 {
		fmt.Println("Number cannot be divided by zero.")
	}
	return firstNumber / secondNumber
}
