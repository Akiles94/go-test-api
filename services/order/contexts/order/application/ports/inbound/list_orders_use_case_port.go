package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/order/contexts/order/domain/models"
)

type ListOrdersUseCasePort interface {
	Execute(ctx context.Context) ([]models.Order, error)
}
