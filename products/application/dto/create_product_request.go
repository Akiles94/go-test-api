package dto

import "github.com/shopspring/decimal"

type CreateProductRequest struct {
	Sku      string          `json:"sku" binding:"required"`
	Name     string          `json:"name" binding:"required"`
	Category string          `json:"category" binding:"required"`
	Price    decimal.Decimal `json:"price" binding:"required,min=0"`
}
