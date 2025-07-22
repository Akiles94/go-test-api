package dto

type ProductsListResponse struct {
	Products   []ProductResponse `json:"products"`
	NextCursor *string           `json:"nextCursor,omitempty"`
}
