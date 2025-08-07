package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/product/application/ports/outbound"
	"github.com/google/uuid"
)

type DeleteProductUseCase struct {
	repo outbound.ProductRepositoryPort
}

func NewDeleteProductUseCase(repo outbound.ProductRepositoryPort) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		repo: repo,
	}
}

func (uc *DeleteProductUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}
