package models

import "github.com/google/uuid"

type orderItem struct {
	id        uuid.UUID
	productID uuid.UUID
	quantity  int
	price     float64
}

type OrderItem interface {
	ID() uuid.UUID
	ProductID() uuid.UUID
	Quantity() int
	Price() float64
}

func NewOrderItem(id, productID uuid.UUID, quantity int, price float64) OrderItem {
	return &orderItem{
		id:        id,
		productID: productID,
		quantity:  quantity,
		price:     price,
	}
}

func (oi *orderItem) ID() uuid.UUID {
	return oi.id
}

func (oi *orderItem) ProductID() uuid.UUID {
	return oi.productID
}

func (oi *orderItem) Quantity() int {
	return oi.quantity
}

func (oi *orderItem) Price() float64 {
	return oi.price
}
