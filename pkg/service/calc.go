package calcapi

import (
	"context"
	"fmt"

	calc "github.com/cham/elk/gen/calc"
	"goa.design/clue/log"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct{}

// Update implements calc.Service.
func (s *calcsrvc) Update(ctx context.Context, req *calc.UpdateAccount) (res *calc.Create, err error) {

	log.Debugf(ctx, "calc.update %#v", req)

	res = &calc.Create{
		Name: &req.Name,
		Href: &req.Email,
	}
	return
}

// Divide implements calc.Service.
func (s *calcsrvc) Divide(ctx context.Context, req *calc.DividePayload) (res int, err error) {

	log.Printf(ctx, "calc.divide %#v", req)

	if req.D == 0 {
		err = calc.MakeDivByZero(fmt.Errorf("division by zero"))
		return
	}
	res = req.C / req.D
	return
}

// NewCalc returns the calc service implementation.
func NewCalc() calc.Service {
	return &calcsrvc{}
}

// Multiply implements multiply.
func (s *calcsrvc) Multiply(ctx context.Context, p *calc.MultiplyPayload) (res int, err error) {
	log.Printf(ctx, "calc.multiply %#v", p)

	res = p.A * p.B
	return
}
