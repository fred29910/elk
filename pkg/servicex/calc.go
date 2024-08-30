package calcapi

import (
	"context"

	calc "github.com/cham/elk/pkg/servicex/gen/calc"
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

// Divide returns the integral division of two integers.
func (s *calcsrvc) Divide(ctx context.Context, p *calc.DividePayload) (res int, err error) {
	log.Printf(ctx, "calc.divide")
	return
}

// Change account name
func (s *calcsrvc) Update(ctx context.Context, p *calc.UpdateAccount) (res *calc.Create, err error) {
	res = &calc.Create{}
	log.Printf(ctx, "calc.update")
	return
}
