package use_cases

import (
	"github.com/Akiles94/go-test-api/application/dto"
	"github.com/Akiles94/go-test-api/application/ports"
	"github.com/Akiles94/go-test-api/domain/models"
	"github.com/google/uuid"
)

type ProductUseCase struct {
	repo ports.IProductRepository
}

func NewProductUseCase(repo ports.IProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repo: repo,
	}
}

func (uc *ProductUseCase) GetPaginated(cursor *string, limit *int) (dto.ProductsResponse, error) {
	return uc.repo.GetPaginated(cursor, limit)
}
func (uc *ProductUseCase) GetByID(id uuid.UUID) (*models.Product, error) {
	return uc.repo.GetByID(id)
}
func (uc *ProductUseCase) Create(product models.Product) error {
	return uc.repo.Create(product)
}
func (uc *ProductUseCase) Update(id uuid.UUID, product *models.Product) error {
	return uc.repo.Update(id, product)
}
func (uc *ProductUseCase) PatchProduct(id uuid.UUID, product *models.ProductPatch) error {
	return uc.repo.PatchProduct(id, product)
}
func (uc *ProductUseCase) Delete(id uuid.UUID) error {
	return uc.repo.Delete(id)
}
