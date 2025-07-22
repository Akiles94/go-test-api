package use_cases

import (
	"github.com/Akiles94/go-test-api/products/application/ports/outbound"
	"github.com/Akiles94/go-test-api/products/domain/models"
	"github.com/google/uuid"
)

type UpdateProductUseCase struct {
	repo outbound.ProductRepositoryPort
}

func NewUpdateProductUseCase(repo outbound.ProductRepositoryPort) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		repo: repo,
	}
}

func (uc *UpdateProductUseCase) Execute(id uuid.UUID, body models.Product) error {
	return uc.repo.Update(id, body)
}
