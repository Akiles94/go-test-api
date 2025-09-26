package outbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/order/contexts/order/domain/models"
	"github.com/google/uuid"
)

type ProductRepositoryPort interface {
	GetProductsByIds(ctx context.Context, ids []uuid.UUID) (*[]models.Product, error)
}
