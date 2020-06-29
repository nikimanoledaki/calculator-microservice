package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
	"github.com/nikimanoledaki/calculator-microservice/server"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewComputation(log)

	protos.RegisterCalculatorServer(gs, cs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
