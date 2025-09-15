package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/product/application/ports/outbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/domain/models"
	"github.com/google/uuid"
)

type GetProductUseCase struct {
	repo outbound.ProductRepositoryPort
}

func NewGetProductUseCase(repo outbound.ProductRepositoryPort) *GetProductUseCase {
	return &GetProductUseCase{
		repo: repo,
	}
}

func (uc *GetProductUseCase) Execute(ctx context.Context, id uuid.UUID) (models.Product, error) {
	return uc.repo.GetByID(ctx, id)
}
