package outbound

import (
	"context"

	"github.com/Akiles94/go-test-api/contexts/product/domain/models"
	"github.com/google/uuid"
)

type ProductRepositoryPort interface {
	Create(ctx context.Context, product models.Product) error
	GetByID(ctx context.Context, id uuid.UUID) (models.Product, error)
	GetAll(ctx context.Context, cursor *string, limit *int) ([]models.Product, *string, error)
	Update(ctx context.Context, id uuid.UUID, product models.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
	Patch(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
}
