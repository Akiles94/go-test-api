package dto

import (
	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductResponse struct {
	ID       uuid.UUID       `json:"id"`
	Sku      string          `json:"sku"`
	Name     string          `json:"name"`
	Category string          `json:"category"`
	Price    decimal.Decimal `json:"price"`
}

func NewProductResponseFromDomainModel(product models.Product) ProductResponse {
	return ProductResponse{
		ID:       product.ID(),
		Sku:      product.Sku(),
		Name:     product.Name(),
		Category: product.Category(),
		Price:    product.Price(),
	}
}
