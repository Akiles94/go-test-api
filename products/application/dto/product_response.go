package dto

import (
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
