package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/order/contexts/order/application/dto"
	"github.com/Akiles94/go-test-api/services/order/contexts/order/application/ports/outbound"
	"github.com/Akiles94/go-test-api/services/order/contexts/order/domain/models"
	"github.com/google/uuid"
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

func (uc *CreateOrderUseCase) Execute(ctx context.Context, payload dto.CreateOrderRequestDto, userID uuid.UUID) (*dto.CreateOrderResponseDto, error) {
	var productIds []uuid.UUID
	for _, productItem := range payload.ProductItems {
		productIds = append(productIds, uuid.MustParse(productItem.ProductID))
	}
	products, err := uc.productRepo.GetProductsByIds(ctx, productIds)
	if err != nil {
		return nil, err
	}
	var orderItems []models.OrderItem
	var productItemsQuantitiesDtoMap = make(map[uuid.UUID]int)
	for _, productItem := range payload.ProductItems {
		productItemsQuantitiesDtoMap[uuid.MustParse(productItem.ProductID)] = productItem.Quantity
	}
	var total float64
	for _, product := range *products {
		orderItems = append(orderItems, models.NewOrderItem(uuid.New(), product.ID(), productItemsQuantitiesDtoMap[product.ID()], product.Price()))
		total += product.Price()
	}
	order := models.NewOrder(
		uuid.New(),
		userID,
		models.StatusPending,
		payload.Address,
		orderItems,
		total,
	)
	err = uc.orderRepo.Create(ctx, order)
	if err != nil {
		return nil, err
	}
	return &dto.CreateOrderResponseDto{
		ID: order.ID(),
	}, nil
}
