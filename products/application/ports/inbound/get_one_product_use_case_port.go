package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/products/domain/models"
	"github.com/google/uuid"
)

type GetOneProductUseCasePort interface {
	Execute(ctx context.Context, id uuid.UUID) (models.Product, error)
}
