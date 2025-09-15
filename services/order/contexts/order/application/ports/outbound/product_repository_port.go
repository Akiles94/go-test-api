package outbound

import (
	"github.com/Akiles94/go-test-api/services/order/contexts/order/domain/models"
	"github.com/google/uuid"
)

type ProductRepositoryPort interface {
	GetProductsByIds(ids []uuid.UUID) (*[]models.Product, error)
}
