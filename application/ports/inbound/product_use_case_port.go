package inbound

import (
	"github.com/Akiles94/go-test-api/application/dto"
	"github.com/Akiles94/go-test-api/domain/models"
	"github.com/google/uuid"
)

type ProductUseCasePort interface {
	GetPaginated(cursor *string, limit *int) (*dto.ProductsResponse, error)
	GetByID(id uuid.UUID) (*models.Product, error)
	Create(body *models.Product) error
	Update(id uuid.UUID, body models.Product) error
	Patch(id uuid.UUID, body models.ProductPatch) error
	Delete(id uuid.UUID) error
}
