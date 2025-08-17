package outbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models"
	"github.com/google/uuid"
)

type CategoryRepositoryPort interface {
	Create(ctx context.Context, category models.Category) error
	GetByID(ctx context.Context, id uuid.UUID) (models.Category, error)
	GetAll(ctx context.Context, cursor *string, limit int) ([]models.Category, *string, error)
	Update(ctx context.Context, category models.Category) error
	Delete(ctx context.Context, id uuid.UUID) error
	ExistsByName(ctx context.Context, name string, excludeID *uuid.UUID) (bool, error)
}
