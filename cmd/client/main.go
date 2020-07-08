package main

import (
	"fmt"
	"log"
	"os"

	client "github.com/nikimanoledaki/calculator-microservice/pkg/client"
	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
	"google.golang.org/grpc"
)

func main() {
	operation, err := client.ParseArguments(os.Args)
	if err != nil {
		log.Fatalf("Error when parsing arguments: %s", err)
	}

	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to port 9092: %s", err)
	}
	defer conn.Close()

	clientService := protos.NewCalculatorClient(conn)

	var response *protos.SumResponse
	if operation == "sum" {
		response, err = client.PrintSum(clientService, os.Args[2:])
	}
	// else {
	// 	fmt.Println(client.PrintAverage(clientService, os.Args[2:]))
	// }
	if err != nil {
		log.Fatalf("%v.GetAverage() = _, %v: ", clientService, err)
		os.Exit(1)
	}

	fmt.Println(response)
}
