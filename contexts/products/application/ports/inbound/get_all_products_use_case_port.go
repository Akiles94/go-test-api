package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
)

type GetAllProductsUseCasePort interface {
	Execute(ctx context.Context, cursor *string, limit *int) ([]models.Product, *string, error)
}
