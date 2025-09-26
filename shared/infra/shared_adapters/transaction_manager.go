package shared_adapters

import (
	"context"

	"gorm.io/gorm"
)

type contextKey string

const txKey contextKey = "tx"

type GormTransactionManager struct {
	db *gorm.DB
}

func NewGormTransactionManager(db *gorm.DB) *GormTransactionManager {
	return &GormTransactionManager{db: db}
}

func (tm *GormTransactionManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return tm.db.Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, txKey, tx)
		return fn(txCtx)
	})
}

func (tm *GormTransactionManager) GetDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey).(*gorm.DB); ok {
		return tx
	}
	return tm.db
}
