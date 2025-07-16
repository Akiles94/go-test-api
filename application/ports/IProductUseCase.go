package ports

import (
	"github.com/Akiles94/go-test-api/application/dto"
	"github.com/Akiles94/go-test-api/domain/models"
	"github.com/google/uuid"
)

type IProductUseCase interface {
	GetPaginated(cursor *string, limit *int) (dto.ProductsResponse, error)
	GetByID(id uint) (*models.Product, error)
	Create(product models.Product) error
	Update(id uuid.UUID, product *models.Product) error
	PatchProduct(id uuid.UUID, p *models.ProductPatch) error
	Delete(id uuid.UUID) error
}
