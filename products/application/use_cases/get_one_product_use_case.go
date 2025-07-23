package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/products/application/ports/outbound"
	"github.com/Akiles94/go-test-api/products/domain/models"
	"github.com/google/uuid"
)

type GetOneProductUseCase struct {
	repo outbound.ProductRepositoryPort
}

func NewGetOneProductUseCase(repo outbound.ProductRepositoryPort) *GetOneProductUseCase {
	return &GetOneProductUseCase{
		repo: repo,
	}
}

func (uc *GetOneProductUseCase) Execute(ctx context.Context, id uuid.UUID) (models.Product, error) {
	return uc.repo.GetByID(ctx, id)
}
