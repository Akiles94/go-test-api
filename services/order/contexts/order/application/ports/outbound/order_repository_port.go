package outbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/order/contexts/order/domain/models"
	"github.com/google/uuid"
)

type OrderRepositoryPort interface {
	Create(ctx context.Context, order models.Order) error
	CreateOrderItem(ctx context.Context, orderItem models.OrderItem, orderID uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (models.Order, error)
}
