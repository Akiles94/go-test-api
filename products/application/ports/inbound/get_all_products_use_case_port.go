package inbound

import (
	"github.com/Akiles94/go-test-api/products/domain/models"
)

type GetAllProductsUseCasePort interface {
	Execute(cursor *string, limit *int) ([]models.Product, *string, error)
}
