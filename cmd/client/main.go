package main

import (
	"context"
	"log"
	"os"

	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9092")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := protos.NewCalculatorClient(conn)

	sum, err := client.GetSum(context.Background(), &protos.SumRequest)
	if err != nil {
		log.Fatalf("%v.GetSum(_) = _, %v: ", client, err)
		os.Exit(1)
	}

	log.Println(sum)

}
