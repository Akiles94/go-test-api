package use_cases

import (
	"github.com/Akiles94/go-test-api/products/application/ports/outbound"
	"github.com/Akiles94/go-test-api/products/domain/models"
)

type CreateProductUseCase struct {
	repo outbound.ProductRepositoryPort
}

func NewCreateProductUseCase(repo outbound.ProductRepositoryPort) *CreateProductUseCase {
	return &CreateProductUseCase{
		repo: repo,
	}
}

func (uc *CreateProductUseCase) Execute(product *models.Product) error {
	return uc.repo.Create(product)
}
