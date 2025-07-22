package dto

import "github.com/shopspring/decimal"

type PatchProductRequest struct {
	Sku      *string          `json:"sku,omitempty"`
	Name     *string          `json:"name,omitempty"`
	Category *string          `json:"category,omitempty"`
	Price    *decimal.Decimal `json:"price,omitempty"`
}
