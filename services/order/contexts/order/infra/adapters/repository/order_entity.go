package repository

import (
	"github.com/Akiles94/go-test-api/services/order/contexts/order/domain/models"
	"github.com/google/uuid"
)

type OrderEntity struct {
	ID      uuid.UUID         `gorm:"type:uuid;primaryKey"`
	UserID  uuid.UUID         `gorm:"type:uuid;not null"`
	Status  string            `gorm:"type:varchar(20);not null"`
	Address string            `gorm:"type:text;not null"`
	Total   float64           `gorm:"type:numeric(10,2);not null"`
	Items   []OrderItemEntity `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}

func (OrderEntity) TableName() string {
	return "orders"
}

func (oe *OrderEntity) ToDomainModel(orderItems []models.OrderItem) models.Order {
	return models.NewOrder(oe.ID, oe.UserID, models.Status(oe.Status), oe.Address, orderItems, oe.Total)
}
