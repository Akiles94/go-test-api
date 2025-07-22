package dto

import "github.com/Akiles94/go-test-api/products/domain/models"

type ProductsResult struct {
	Products   []models.Product `json:"products"`
	NextCursor *string          `json:"next_cursor,omitempty"`
}
