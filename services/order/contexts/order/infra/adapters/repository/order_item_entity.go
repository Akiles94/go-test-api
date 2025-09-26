package repository

import (
	"github.com/Akiles94/go-test-api/services/order/contexts/order/domain/models"
	"github.com/google/uuid"
)

type OrderItemEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	Quantity  int       `gorm:"not null"`
	Price     float64   `gorm:"type:numeric(10,2);not null"`
}

func (OrderItemEntity) TableName() string {
	return "order_item"
}

func (oie *OrderItemEntity) ToDomainModel() models.OrderItem {
	return models.NewOrderItem(oie.ID, oie.ProductID, oie.Quantity, oie.Price)
}
