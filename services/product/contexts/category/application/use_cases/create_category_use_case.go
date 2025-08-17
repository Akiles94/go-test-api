package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/inbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/outbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models"
	"github.com/google/uuid"
)

type CreateCategoryUseCase struct {
	categoryRepository outbound.CategoryRepositoryPort
}

func NewCreateCategoryUseCase(categoryRepository outbound.CategoryRepositoryPort) inbound.CreateCategoryUseCasePort {
	return &CreateCategoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (uc *CreateCategoryUseCase) Execute(ctx context.Context, request dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	// Check if category with same name already exists
	exists, err := uc.categoryRepository.ExistsByName(ctx, request.Name, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, models.ErrCategoryNameEmpty // You might want to create a specific error for this
	}

	// Create new category
	categoryID := uuid.New()
	category, err := models.NewCategory(categoryID, request.Name, request.Description, request.IsActive)
	if err != nil {
		return nil, err
	}

	// Save category
	err = uc.categoryRepository.Create(ctx, category)
	if err != nil {
		return nil, err
	}

	// Return response
	return &dto.CategoryResponse{
		ID:          category.ID(),
		Name:        category.Name(),
		Description: category.Description(),
		IsActive:    category.IsActive(),
	}, nil
}
