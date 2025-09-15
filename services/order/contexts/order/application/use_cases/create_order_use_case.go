package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/order/contexts/order/application/dto"
	"github.com/Akiles94/go-test-api/services/order/contexts/order/application/ports/outbound"
)

type CreateOrderUseCase struct {
	orderRepo   outbound.OrderRepositoryPort
	productRepo outbound.ProductRepositoryPort
}

func NewCreateOrderUseCase(orderRepo outbound.OrderRepositoryPort, productRepo outbound.ProductRepositoryPort) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, payload dto.CreateOrderRequestDto) (*dto.CreateOrderResponseDto, error) {
	// Implementation of the use case logic goes here
	return nil, nil
}
