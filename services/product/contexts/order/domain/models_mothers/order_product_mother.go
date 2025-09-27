package models_mothers

import "github.com/Akiles94/go-test-api/services/product/contexts/order/domain/models"

type OrderProductMother struct {
	ID    string
	Price float64
	Stock int
}

func NewOrderProductMother() *OrderProductMother {
	return &OrderProductMother{
		ID:    "00000000-0000-0000-0000-000000000001",
		Price: 49.99,
		Stock: 100,
	}
}

func (opm *OrderProductMother) WithID(id string) *OrderProductMother {
	opm.ID = id
	return opm
}

func (opm *OrderProductMother) WithPrice(price float64) *OrderProductMother {
	opm.Price = price
	return opm
}

func (opm *OrderProductMother) WithStock(stock int) *OrderProductMother {
	opm.Stock = stock
	return opm
}

func (opm *OrderProductMother) Build() models.OrderProduct {
	return models.NewOrderProduct(opm.ID, opm.Price, opm.Stock)
}
