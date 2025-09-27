package dto

import "github.com/google/uuid"

type OrderProductRequest struct {
	ProductsIds []uuid.UUID `json:"products_ids"`
}