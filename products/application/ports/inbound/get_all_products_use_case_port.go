package inbound

import (
	"github.com/Akiles94/go-test-api/products/application/dto"
)

type GetAllProductsUseCasePort interface {
	Execute(cursor *string, limit *int) (*dto.ProductsResponse, error)
}
