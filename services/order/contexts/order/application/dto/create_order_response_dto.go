package dto

import "github.com/google/uuid"

type CreateOrderResponseDto struct {
	ID uuid.UUID `json:"id"`
}
