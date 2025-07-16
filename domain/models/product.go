package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Sku      string    `json:"sku"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Price    int       `json:"price"`
}

type ProductPatch struct {
	Sku      *string `json:"sku,omitempty"`
	Name     *string `json:"name,omitempty"`
	Category *string `json:"category,omitempty"`
	Price    *int    `json:"price,omitempty"`
}
