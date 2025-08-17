package use_cases_tests

import (
	"context"
	"testing"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/use_cases"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/use_cases/use_cases_mocks"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models/models_mothers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetOneCategoryUseCase_ValidId_ShouldReturnCategory(t *testing.T) {
	// Arrange
	mockRepo := &use_cases_mocks.MockCategoryRepository{}
	useCase := use_cases.NewGetOneCategoryUseCase(mockRepo)
	mother := models_mothers.NewCategoryMother()

	categoryID := uuid.New()
	category := mother.ValidCategoryWithId(categoryID)

	mockRepo.On("GetByID", mock.Anything, categoryID).Return(category, nil)

	// Act
	response, err := useCase.Execute(context.Background(), categoryID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, categoryID, response.ID)
	assert.Equal(t, category.Name(), response.Name)
	assert.Equal(t, category.Description(), response.Description)
	assert.Equal(t, category.IsActive(), response.IsActive)
	mockRepo.AssertExpectations(t)
}

func TestGetOneCategoryUseCase_InvalidId_ShouldReturnError(t *testing.T) {
	// Arrange
	mockRepo := &use_cases_mocks.MockCategoryRepository{}
	useCase := use_cases.NewGetOneCategoryUseCase(mockRepo)

	categoryID := uuid.New()
	expectedError := assert.AnError

	mockRepo.On("GetByID", mock.Anything, categoryID).Return(nil, expectedError)

	// Act
	response, err := useCase.Execute(context.Background(), categoryID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
