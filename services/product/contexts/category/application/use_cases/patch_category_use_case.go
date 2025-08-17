package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/inbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/outbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models"
	"github.com/google/uuid"
)

type PatchCategoryUseCase struct {
	categoryRepository outbound.CategoryRepositoryPort
}

func NewPatchCategoryUseCase(categoryRepository outbound.CategoryRepositoryPort) inbound.PatchCategoryUseCasePort {
	return &PatchCategoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (uc *PatchCategoryUseCase) Execute(ctx context.Context, id uuid.UUID, request dto.PatchCategoryRequest) (*dto.CategoryResponse, error) {
	// Get existing category
	existingCategory, err := uc.categoryRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Prepare updated values
	name := existingCategory.Name()
	description := existingCategory.Description()
	isActive := existingCategory.IsActive()

	if request.Name != nil {
		name = *request.Name
		// Check if another category with same name already exists
		exists, err := uc.categoryRepository.ExistsByName(ctx, name, &id)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, models.ErrCategoryNameEmpty // You might want to create a specific error for this
		}
	}

	if request.Description != nil {
		description = *request.Description
	}

	if request.IsActive != nil {
		isActive = *request.IsActive
	}

	// Create updated category
	updatedCategory, err := models.NewCategory(id, name, description, isActive)
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
