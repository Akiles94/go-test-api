package use_cases_tests

import (
	"context"
	"testing"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/use_cases"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/use_cases/use_cases_mocks"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models/models_mothers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCategoryUseCase_ValidId_ShouldDeleteCategory(t *testing.T) {
	// Arrange
	mockRepo := use_cases_mocks.NewMockCategoryRepository()
	categoryID := uuid.New()
	category := models_mothers.NewCategoryMother().MustBuild()
	mockRepo.SetupGetByIDSuccess(category)
	mockRepo.SetupDeleteSuccess()
	useCase := use_cases.NewDeleteCategoryUseCase(mockRepo)

	// Act
	err := useCase.Execute(context.Background(), categoryID)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCategoryUseCase_InvalidId_ShouldReturnError(t *testing.T) {
	// Arrange
	mockRepo := use_cases_mocks.NewMockCategoryRepository()
	categoryID := uuid.New()
	category := models_mothers.NewCategoryMother().MustBuild()
	mockRepo.SetupGetByIDSuccess(category)
	mockRepo.SetupDeleteError()
	useCase := use_cases.NewDeleteCategoryUseCase(mockRepo)

	expectedError := assert.AnError

	// Act
	err := useCase.Execute(context.Background(), categoryID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
