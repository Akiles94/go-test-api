package shared_ports

import (
	"context"

	"gorm.io/gorm"
)

type TransactionManagerPort interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
	GetDB(ctx context.Context) *gorm.DB
}
