package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/product/domain/models"
)

type ListProductsUseCasePort interface {
	Execute(ctx context.Context, cursor *string, limit *int) (products []models.Product, nextCursor *string, err error)
}
