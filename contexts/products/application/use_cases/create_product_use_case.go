package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/contexts/products/application/ports/outbound"
	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
)

type CreateProductUseCase struct {
	repo outbound.ProductRepositoryPort
}

func NewCreateProductUseCase(repo outbound.ProductRepositoryPort) *CreateProductUseCase {
	return &CreateProductUseCase{
		repo: repo,
	}
}

func (uc *CreateProductUseCase) Execute(ctx context.Context, product models.Product) error {
	return uc.repo.Create(ctx, product)
}
