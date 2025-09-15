package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/order/contexts/order/domain/models"
	"github.com/google/uuid"
)

type GetOrderUseCasePort interface {
	Execute(ctx context.Context, id uuid.UUID) (models.Order, error)
}
