package ports

import (
	"github.com/Akiles94/go-test-api/application/dto"
	"github.com/Akiles94/go-test-api/domain/models"
	"github.com/google/uuid"
)

type IProductRepository interface {
	GetPaginated(cursor *string, limit *int) (dto.ProductsResponse, error)
	GetByID(id uuid.UUID) (*models.Product, error)
	Create(product models.Product) error
	Update(id uuid.UUID, product *models.Product) error
	PatchProduct(id uuid.UUID, p *models.ProductPatch) error
	Delete(id uuid.UUID) error
}
