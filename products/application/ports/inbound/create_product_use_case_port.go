package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/products/domain/models"
)

type CreateProductUseCasePort interface {
	Execute(ctx context.Context, product models.Product) error
}
