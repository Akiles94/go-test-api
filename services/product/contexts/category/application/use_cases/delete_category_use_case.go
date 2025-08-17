package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/inbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/outbound"
	"github.com/google/uuid"
)

type DeleteCategoryUseCase struct {
	categoryRepository outbound.CategoryRepositoryPort
}

func NewDeleteCategoryUseCase(categoryRepository outbound.CategoryRepositoryPort) inbound.DeleteCategoryUseCasePort {
	return &DeleteCategoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (uc *DeleteCategoryUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	// Check if category exists first
	_, err := uc.categoryRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Delete category
	return uc.categoryRepository.Delete(ctx, id)
}
