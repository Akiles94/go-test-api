package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Sku      string    `json:"sku" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Category string    `json:"category" binding:"required"`
	Price    int       `json:"price" binding:"required"`
}

type ProductPatch struct {
	Sku      *string `json:"sku,omitempty"`
	Name     *string `json:"name,omitempty"`
	Category *string `json:"category,omitempty"`
	Price    *int    `json:"price,omitempty"`
}
