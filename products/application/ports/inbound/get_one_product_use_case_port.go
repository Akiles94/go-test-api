package inbound

import (
	"github.com/Akiles94/go-test-api/products/domain/models"
	"github.com/google/uuid"
)

type GetOneProductUseCasePort interface {
	Execute(id uuid.UUID) (models.Product, error)
}
