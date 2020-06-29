package calculator

import "fmt"

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
