package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
)

// CalculatorServer is a struct used to create methods that implement
// the gRPC interface.
type CalculatorServer struct {
	log hclog.Logger
}

// NewComputation is a constructor that implements the Calculator service interface.
func NewComputation(l hclog.Logger) *CalculatorServer {
	return &CalculatorServer{l}
}

// GetSum implements the gRPC method for GetSum by handling the SumRequest and returning the SumResponse.
func (cs *CalculatorServer) GetSum(ctx context.Context, sum *protos.SumRequest) (*protos.SumResponse, error) {
	cs.log.Info("Handle GetSum", "firstNumber", sum.GetFirstNumber(), "secondNumber", sum.GetSecondNumber())

	return &protos.SumResponse{Result: 2}, nil
}

// GetAverage implements the gRPC method for GetAverage by handling the AverageRequest and returning the AverageResponse.
func (cs *CalculatorServer) GetAverage(ctx context.Context, average *protos.AverageRequest) (*protos.AverageResponse, error) {
	cs.log.Info("Handle GetAverage", "firstNumber", average.GetFirstNumber(), "secondNumber", average.GetSecondNumber())

	return &protos.AverageResponse{Result: 2}, nil
}
