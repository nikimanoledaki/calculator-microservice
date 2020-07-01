package main

import (
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

	// TODO: Move this logic to pkg/client and test it
	if operation == "sum" {
		client.PrintSum(clientService, os.Args[2:])
	} else {
		client.PrintAverage(clientService, os.Args[2:])
	}

}
