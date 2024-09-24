package service

import (
	"context"

	"gorm.io/gorm"
)

type Transation interface {
	InTransaction(ctx context.Context, fn func(ctx context.Context) error) error
	DB() *gorm.DB
}
