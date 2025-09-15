package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/inbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/outbound"
	"github.com/Akiles94/go-test-api/shared/application/shared_dto"
)

type ListCategoriesUseCase struct {
	categoryRepository outbound.CategoryRepositoryPort
}

func NewListCategoriesUseCase(categoryRepository outbound.CategoryRepositoryPort) inbound.ListCategoriesUseCasePort {
	return &ListCategoriesUseCase{
		categoryRepository: categoryRepository,
	}
}

func (uc *ListCategoriesUseCase) Execute(ctx context.Context, cursor *string, limit int) (*dto.PaginatedCategoryResponse, error) {
	categories, nextCursor, err := uc.categoryRepository.GetAll(ctx, cursor, limit)
	if err != nil {
		return nil, err
	}

	categoryResponses := make([]dto.CategoryResponse, len(categories))
	for i, category := range categories {
		categoryResponses[i] = dto.CategoryResponse{
			ID:          category.ID(),
			Name:        category.Name(),
			Description: category.Description(),
			IsActive:    category.IsActive(),
		}
	}

	result := shared_dto.NewPaginatedResult(categoryResponses, nextCursor)
	return &result, nil
}
