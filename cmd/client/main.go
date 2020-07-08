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

	err = client.NewRequest(operation, clientService, os.Args[2:])
	if err != nil {
		log.Fatalf("Failed to call %v on %v: %s", operation, clientService, err)
		os.Exit(1)
	}

}
