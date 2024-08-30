package calcapi

import (
	"context"

	calc "github.com/cham/elk/gen/calc"
	"goa.design/clue/log"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct{}

// NewCalc returns the calc service implementation.
func NewCalc() calc.Service {
	return &calcsrvc{}
}

// Multiply implements multiply.
func (s *calcsrvc) Multiply(ctx context.Context, p *calc.MultiplyPayload) (res int, err error) {
	log.Printf(ctx, "calc.multiply")
	return
}
