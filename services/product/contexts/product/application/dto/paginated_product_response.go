package dto

// PaginatedProductResponse represents a paginated response for products
type PaginatedProductResponse struct {
	Items      []ProductResponse `json:"items"`
	NextCursor *string           `json:"next_cursor,omitempty"`
}

// NewPaginatedProductResponse creates a new paginated product response
func NewPaginatedProductResponse(products []ProductResponse, nextCursor *string) PaginatedProductResponse {
	return PaginatedProductResponse{
		Items:      products,
		NextCursor: nextCursor,
	}
}
