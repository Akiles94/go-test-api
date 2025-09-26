package repository

import (
	"context"

	"github.com/Akiles94/go-test-api/services/order/contexts/order/domain/models"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository struct {
	txManager shared_ports.TransactionManagerPort
}

func NewOrderRepository(txManager shared_ports.TransactionManagerPort) *OrderRepository {
	return &OrderRepository{
		txManager: txManager,
	}
}

func (or *OrderRepository) Create(ctx context.Context, order models.Order) error {
	db := or.txManager.GetDB(ctx)

	orderEntity := &OrderEntity{
		ID:      order.ID(),
		UserID:  order.UserID(),
		Status:  string(order.Status()),
		Address: order.Address(),
		Total:   order.Total(),
	}

	if err := db.Create(orderEntity).Error; err != nil {
		return err
	}

	for _, item := range order.Items() {
		orderItemEntity := &OrderItemEntity{
			ID:        uuid.New(),
			OrderID:   order.ID(),
			ProductID: item.ProductID(),
			Quantity:  item.Quantity(),
			Price:     item.Price(),
		}
		if err := db.Create(orderItemEntity).Error; err != nil {
			return err
		}
	}

	return nil
}

func (or *OrderRepository) GetByID(ctx context.Context, id uuid.UUID) (models.Order, error) {
	var entity OrderEntity
	if err := or.txManager.GetDB(ctx).Preload("Items").First(&entity, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	var orderItems []models.OrderItem
	for _, itemEntity := range entity.Items {
		orderItem := itemEntity.ToDomainModel()
		orderItems = append(orderItems, orderItem)
	}

	return entity.ToDomainModel(orderItems), nil
}
