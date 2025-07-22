package inbound

import (
	"github.com/Akiles94/go-test-api/products/domain/models"
)

type CreateProductUseCasePort interface {
	Execute(product *models.Product) error
}
