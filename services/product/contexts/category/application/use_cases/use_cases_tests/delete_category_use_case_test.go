package use_cases_tests

import (
	"context"
	"testing"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/use_cases"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/use_cases/use_cases_mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteCategoryUseCase_ValidId_ShouldDeleteCategory(t *testing.T) {
	// Arrange
	mockRepo := &use_cases_mocks.MockCategoryRepository{}
	useCase := use_cases.NewDeleteCategoryUseCase(mockRepo)

	categoryID := uuid.New()

	mockRepo.On("GetByID", mock.Anything, categoryID).Return(mock.Anything, nil)
	mockRepo.On("Delete", mock.Anything, categoryID).Return(nil)

	// Act
	err := useCase.Execute(context.Background(), categoryID)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCategoryUseCase_InvalidId_ShouldReturnError(t *testing.T) {
	// Arrange
	mockRepo := &use_cases_mocks.MockCategoryRepository{}
	useCase := use_cases.NewDeleteCategoryUseCase(mockRepo)

	categoryID := uuid.New()
	expectedError := assert.AnError

	mockRepo.On("GetByID", mock.Anything, categoryID).Return(nil, expectedError)

	// Act
	err := useCase.Execute(context.Background(), categoryID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
