package models_mothers

import (
	"github.com/Akiles94/go-test-api/services/order/contexts/order/domain/models"
	"github.com/google/uuid"
)

type OrderMother struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	Status  string
	Address string
	Items   []models.OrderItem
	Total   float64
}

func NewOrderMother() *OrderMother {
	return &OrderMother{
		ID:      uuid.MustParse("00000000-0000-0000-0000-000000000123"),
		UserID:  uuid.MustParse("00000000-0000-0000-0000-000000000456"),
		Status:  "pending",
		Address: "123 Main St, Anytown, USA",
		Items: []models.OrderItem{
			models.NewOrderItem(uuid.MustParse("00000000-0000-0000-0000-000000000001"), 1),
			models.NewOrderItem(uuid.MustParse("00000000-0000-0000-0000-000000000002"), 2),
		},
		Total: 99.99,
	}
}

func (om *OrderMother) WithID(id string) *OrderMother {
	om.ID = uuid.MustParse(id)
	return om
}

func (om *OrderMother) WithUserID(userID string) *OrderMother {
	om.UserID = uuid.MustParse(userID)
	return om
}

func (om *OrderMother) WithStatus(status string) *OrderMother {
	om.Status = status
	return om
}

func (om *OrderMother) WithAddress(address string) *OrderMother {
	om.Address = address
	return om
}

func (om *OrderMother) WithItems(items []models.OrderItem) *OrderMother {
	om.Items = items
	return om
}

func (om *OrderMother) WithTotal(total float64) *OrderMother {
	om.Total = total
	return om
}

func (om *OrderMother) Build() *models.Order {
	model := models.NewOrder(
		om.ID,
		om.UserID,
		models.Status(om.Status),
		om.Address,
		om.Items,
		om.Total,
	)
	return &model
}
