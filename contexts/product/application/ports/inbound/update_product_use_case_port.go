package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/contexts/product/domain/models"
	"github.com/google/uuid"
)

type UpdateProductUseCasePort interface {
	Execute(ctx context.Context, id uuid.UUID, product models.Product) error
}
