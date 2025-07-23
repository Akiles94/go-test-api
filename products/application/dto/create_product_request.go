package dto

import (
	"github.com/Akiles94/go-test-api/products/domain/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateProductRequest struct {
	Sku      string  `json:"sku" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Price    float64 `json:"price" binding:"required,min=0"`
}

func (c *CreateProductRequest) ToDomainModel() (models.Product, error) {
	return models.NewProduct(
		uuid.New(),
		c.Sku,
		c.Name,
		c.Category,
		decimal.NewFromFloat(c.Price),
	)
}
