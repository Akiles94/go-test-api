package dto

import "github.com/Akiles94/go-test-api/products/domain/models"

type ProductsResponse struct {
	Products   []models.Product `json:"products"`
	NextCursor *string          `json:"nextCursor,omitempty"`
}
