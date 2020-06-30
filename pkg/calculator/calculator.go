package calculator

import (
	"errors"
	"fmt"
)

// Compute performs the arithmetic operation on two given numbers based on a given type of artithmetic operation.
func Compute(operation string, firstNumber, secondNumber int) (int, error) {
	switch operation {
	case "sum":
		return Sum(firstNumber, secondNumber), nil
	case "subtract":
		return Subtract(firstNumber, secondNumber), nil
	case "multiply":
		return Multiply(firstNumber, secondNumber), nil
	case "divide":
		return Divide(firstNumber, secondNumber), nil
	}
	return 0, errors.New("operation could not be found")
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
		fmt.Println("number cannot be divided by zero")
	}
	return firstNumber / secondNumber
}
