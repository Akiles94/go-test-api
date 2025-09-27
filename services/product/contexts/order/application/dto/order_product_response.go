package dto

import "github.com/Akiles94/go-test-api/services/product/contexts/order/domain/models"

type OrderProductResponse struct {
	Products []models.OrderProduct `json:"products"`
}
