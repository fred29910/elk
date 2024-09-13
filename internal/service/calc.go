package service

import (
	"context"

	calc "github.com/cham/elk/internal/gen/calc"
)

// 声明 svc 实现了 calc.Service 接口
var _ calc.Service = (*svc)(nil)

type svc struct{}

func (s *svc) Multiply(ctx context.Context, p *calc.MultiplyPayload) (int, error) {
	return p.A + p.B, nil
}
