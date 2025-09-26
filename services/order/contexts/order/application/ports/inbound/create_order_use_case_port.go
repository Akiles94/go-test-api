package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/order/contexts/order/application/dto"
	"github.com/google/uuid"
)

type CreateOrderUseCasePort interface {
	Execute(ctx context.Context, body dto.CreateOrderRequestDto, userID uuid.UUID) (*dto.CreateOrderResponseDto, error)
}
