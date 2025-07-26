package dto

type UpdateProductRequest struct {
	Sku      string  `json:"sku" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Category string  `json:"category" binding:"required"`
	Price    float64 `json:"price" binding:"required,min=0"`
}
