package dto

import (
	"github.com/Akiles94/go-test-api/services/product/contexts/product/domain/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductResponse struct {
	ID         uuid.UUID         `json:"id"`
	Sku        string            `json:"sku"`
	Name       string            `json:"name"`
	CategoryId uuid.UUID         `json:"category_id"`
	Category   *CategoryResponse `json:"category"`
	Price      decimal.Decimal   `json:"price"`
}

type CategoryResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
}

func NewProductResponseFromDomainModel(product models.Product) ProductResponse {
	var categoryResponse *CategoryResponse
	if product.Category() != nil {
		productCategory := *product.Category()
		categoryResponse = &CategoryResponse{
			ID:          productCategory.ID(),
			Name:        productCategory.Name(),
			Slug:        productCategory.Slug(),
			Description: productCategory.Description(),
		}
	}
	return ProductResponse{
		ID:       product.ID(),
		Sku:      product.Sku(),
		Name:     product.Name(),
		Category: categoryResponse,
		Price:    product.Price(),
	}
}
