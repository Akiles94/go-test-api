package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/inbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/outbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models"
	"github.com/google/uuid"
)

type UpdateCategoryUseCase struct {
	categoryRepository outbound.CategoryRepositoryPort
}

func NewUpdateCategoryUseCase(categoryRepository outbound.CategoryRepositoryPort) inbound.UpdateCategoryUseCasePort {
	return &UpdateCategoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (uc *UpdateCategoryUseCase) Execute(ctx context.Context, id uuid.UUID, request dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	// Check if category exists
	_, err := uc.categoryRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Check if another category with same name already exists
	exists, err := uc.categoryRepository.ExistsByName(ctx, request.Name, &id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, models.ErrCategoryNameEmpty // You might want to create a specific error for this
	}

	// Create updated category
	updatedCategory, err := models.NewCategory(id, request.Name, request.Description, request.IsActive)
	if err != nil {
		return nil, err
	}

	// Update category
	err = uc.categoryRepository.Update(ctx, updatedCategory)
	if err != nil {
		return nil, err
	}

	// Return response
	return &dto.CategoryResponse{
		ID:          updatedCategory.ID(),
		Name:        updatedCategory.Name(),
		Description: updatedCategory.Description(),
		IsActive:    updatedCategory.IsActive(),
	}, nil
}
