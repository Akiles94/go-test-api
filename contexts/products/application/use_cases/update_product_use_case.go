package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/contexts/products/application/ports/outbound"
	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
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

func (uc *UpdateProductUseCase) Execute(ctx context.Context, id uuid.UUID, body models.Product) error {
	return uc.repo.Update(ctx, id, body)
}
