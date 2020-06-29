package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
)

// Calculator is a constructor.
type Calculator struct {
	log hclog.Logger
}

// NewComputation does something.
func NewComputation(l hclog.Logger) *Calculator {
	return &Calculator{l}
}

// GetComputation does something.
func (c *Calculator) GetComputation(ctx context.Context, rr *protos.ComputationRequest) (*protos.ComputationResponse, error) {
	c.log.Info("Handle GetComputation", "operand", rr.GetOperand(), "firstNumber", rr.GetFirstNumber(), "secondNumber", rr.GetSecondNumber())

	return &protos.ComputationResponse{Result: 2}, nil
}
