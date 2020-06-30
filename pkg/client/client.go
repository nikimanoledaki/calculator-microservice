package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
)

// ParseArguments manages the error handling for the unhappy paths of the client.
func ParseArguments(args []string) (string, error) {
	operation := args[1]
	if operation != "sum" && operation != "average" {
		return "", fmt.Errorf("operation not recognized")
	}

	numbersAsStrings := args[2:]
	if len(numbersAsStrings) != 2 {
		return "", fmt.Errorf("expected 2 arguments")
	}

	return operation, nil
}

// PrintSum does something.
func PrintSum(client protos.CalculatorClient, args []string) {

	numbers := make([]int32, 2)
	for i, arg := range args {
		number, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("failed to convert %d to integer", err)
		}
		numbers[i] = int32(number)
	}

	sumReq := &protos.SumRequest{
		FirstNumber:  numbers[0],
		SecondNumber: numbers[1],
	}

	response, err := client.GetSum(context.Background(), sumReq)
	if err != nil {
		log.Fatalf("%v.GetSum() = _, %v: ", client, err)
		os.Exit(1)
	}

	log.Println(response)
}

// PrintAverage does something.
func PrintAverage(client protos.CalculatorClient, args []string) {
	numbers := make([]float32, 2)
	for i, arg := range args {
		number, err := strconv.ParseFloat(arg, 32)
		if err != nil {
			log.Fatalf("failed to convert %d to float", err)
		}
		numbers[i] = float32(number)
	}

	avgReq := &protos.AverageRequest{
		FirstNumber:  numbers[0],
		SecondNumber: numbers[1],
	}

	response, err := client.GetAverage(context.Background(), avgReq)
	if err != nil {
		log.Fatalf("%v.GetAverage() = _, %v: ", client, err)
		os.Exit(1)
	}

	log.Println(response)
}
