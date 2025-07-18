package use_cases

import (
	"github.com/Akiles94/go-test-api/application/dto"
	"github.com/Akiles94/go-test-api/application/ports/outbound"
	"github.com/Akiles94/go-test-api/domain/models"
	"github.com/google/uuid"
)

type ProductUseCase struct {
	repo outbound.IProductRepository
}

func NewProductUseCase(repo outbound.IProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repo: repo,
	}
}

func (uc *ProductUseCase) GetPaginated(cursor *string, limit *int) (*dto.ProductsResponse, error) {
	return uc.repo.GetPaginated(cursor, limit)
}
func (uc *ProductUseCase) GetByID(id uuid.UUID) (*models.Product, error) {
	return uc.repo.GetByID(id)
}
func (uc *ProductUseCase) Create(body *models.Product) error {
	return uc.repo.Create(body)
}
func (uc *ProductUseCase) Update(id uuid.UUID, body models.Product) error {
	return uc.repo.Update(id, body)
}
func (uc *ProductUseCase) Patch(id uuid.UUID, body models.ProductPatch) error {
	return uc.repo.Patch(id, body)
}
func (uc *ProductUseCase) Delete(id uuid.UUID) error {
	return uc.repo.Delete(id)
}
