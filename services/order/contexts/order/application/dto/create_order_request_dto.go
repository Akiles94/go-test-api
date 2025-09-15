package dto

type ProductItem struct {
	ProductID string `json:"productId" binding:"required,uuid"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

type CreateOrderRequestDto struct {
	UserID       string        `json:"userId" binding:"required,uuid"`
	Status       string        `json:"status" binding:"required,oneof=pending completed cancelled"`
	Address      string        `json:"address" binding:"required"`
	ProductItems []ProductItem `json:"productItems" binding:"required,dive"`
}
