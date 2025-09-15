package outbound

import (
	"github.com/Akiles94/go-test-api/services/order/contexts/order/domain/models"
	"github.com/google/uuid"
)

type OrderRepositoryPort interface {
	Create(order models.Order) error
	GetByID(id uuid.UUID) (models.Order, error)
}
