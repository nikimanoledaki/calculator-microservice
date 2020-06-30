package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
)

// Calculator is a struct used to create methods that implement
// the gRPC interface.
type Calculator struct {
	log hclog.Logger
}

// NewComputation is a constructor that implements the Calculator service interface.
func NewComputation(l hclog.Logger) *Calculator {
	return &Calculator{l}
}

// GetComputation implements the gRPC interface to handle the ComputationRequest
// and return the ComputationResponse methods.
func (c *Calculator) GetComputation(ctx context.Context, rr *protos.ComputationRequest) (*protos.ComputationResponse, error) {
	c.log.Info("Handle GetComputation", "Operation", rr.GetOperation(), "firstNumber", rr.GetFirstNumber(), "secondNumber", rr.GetSecondNumber())

	return &protos.ComputationResponse{Result: 2}, nil
}
