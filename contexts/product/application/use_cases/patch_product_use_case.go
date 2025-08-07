package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/contexts/product/application/ports/outbound"
	"github.com/google/uuid"
)

type PatchProductUseCase struct {
	repo outbound.ProductRepositoryPort
}

func NewPatchProductUseCase(repo outbound.ProductRepositoryPort) *PatchProductUseCase {
	return &PatchProductUseCase{
		repo: repo,
	}
}

func (uc *PatchProductUseCase) Execute(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	return uc.repo.Patch(ctx, id, updates)
}
