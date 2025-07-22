package inbound

import (
	"github.com/Akiles94/go-test-api/products/domain/models"
	"github.com/google/uuid"
)

type UpdateProductUseCasePort interface {
	Execute(id uuid.UUID, product models.Product) error
}
