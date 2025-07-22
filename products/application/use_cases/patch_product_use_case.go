package use_cases

import (
	"github.com/Akiles94/go-test-api/products/application/ports/outbound"
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

func (uc *PatchProductUseCase) Execute(id uuid.UUID, updates map[string]interface{}) error {
	return uc.repo.Patch(id, updates)
}
