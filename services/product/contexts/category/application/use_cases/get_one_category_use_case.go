package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/inbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/outbound"
	"github.com/google/uuid"
)

type GetOneCategoryUseCase struct {
	categoryRepository outbound.CategoryRepositoryPort
}

func NewGetOneCategoryUseCase(categoryRepository outbound.CategoryRepositoryPort) inbound.GetOneCategoryUseCasePort {
	return &GetOneCategoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (uc *GetOneCategoryUseCase) Execute(ctx context.Context, id uuid.UUID) (*dto.CategoryResponse, error) {
	category, err := uc.categoryRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:          category.ID(),
		Name:        category.Name(),
		Description: category.Description(),
		IsActive:    category.IsActive(),
	}, nil
}
