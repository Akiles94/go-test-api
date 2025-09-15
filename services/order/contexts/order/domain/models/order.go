package models

import "github.com/google/uuid"

type Status string

const (
	StatusPending   Status = "pending"
	StatusCompleted Status = "completed"
	StatusCancelled Status = "cancelled"
)

type orderItem struct {
	productID uuid.UUID
	quantity  int
}

type OrderItem interface {
	ProductID() uuid.UUID
	Quantity() int
}

type Order interface {
	ID() uuid.UUID
	UserID() uuid.UUID
	Status() Status
	Address() string
	Items() []OrderItem
	Total() float64
}

type order struct {
	id      uuid.UUID
	userID  uuid.UUID
	status  Status
	address string
	items   []OrderItem
	total   float64
}

func NewOrder(id, userID uuid.UUID, status Status, address string, items []OrderItem, total float64) Order {
	return &order{
		id:      id,
		userID:  userID,
		status:  status,
		address: address,
		items:   items,
		total:   total,
	}
}

func NewOrderItem(productID uuid.UUID, quantity int) OrderItem {
	return &orderItem{
		productID: productID,
		quantity:  quantity,
	}
}

func (o *order) ID() uuid.UUID {
	return o.id
}

func (o *order) UserID() uuid.UUID {
	return o.userID
}

func (o *order) Status() Status {
	return o.status
}

func (o *order) Address() string {
	return o.address
}

func (o *order) Items() []OrderItem {
	return o.items
}

func (o *order) Total() float64 {
	return o.total
}

func (oi *orderItem) ProductID() uuid.UUID {
	return oi.productID
}

func (oi *orderItem) Quantity() int {
	return oi.quantity
}
