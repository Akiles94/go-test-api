package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/order/contexts/order/application/dto"
)

type CreateOrderUseCasePort interface {
	Execute(ctx context.Context, body dto.CreateOrderRequestDto) (*dto.CreateOrderResponseDto, error)
}
