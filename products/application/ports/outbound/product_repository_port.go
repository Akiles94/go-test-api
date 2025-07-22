package outbound

import (
	"github.com/Akiles94/go-test-api/products/application/dto"
	"github.com/Akiles94/go-test-api/products/domain/models"
	"github.com/google/uuid"
)

type ProductRepositoryPort interface {
	Create(product *models.Product) error
	GetByID(id uuid.UUID) (*models.Product, error)
	GetAll(cursor *string, limit *int) (*dto.ProductsResponse, error)
	Update(id uuid.UUID, product models.Product) error
	Delete(id uuid.UUID) error
	Patch(id uuid.UUID, product dto.ProductPatchBody) error
}
