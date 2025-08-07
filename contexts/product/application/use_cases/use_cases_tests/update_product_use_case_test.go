package use_cases_tests

import (
	"context"
	"errors"
	"testing"

	"github.com/Akiles94/go-test-api/contexts/product/application/use_cases"
	"github.com/Akiles94/go-test-api/contexts/product/application/use_cases/use_cases_mocks"
	"github.com/Akiles94/go-test-api/contexts/product/domain/models/models_mothers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func UpdateProductUseCase_Execute(t *testing.T) {
	t.Run("should update product successfully", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := use_cases_mocks.NewMockProductRepository()
		useCase := use_cases.NewUpdateProductUseCase(mockRepo)

		productID := uuid.New()
		product := models_mothers.NewProductMother().MustBuild()
		mockRepo.SetupUpdateSuccess(productID, product)
		// Act
		err := useCase.Execute(ctx, productID, product)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := use_cases_mocks.NewMockProductRepository()
		useCase := use_cases.NewUpdateProductUseCase(mockRepo)

		productID := uuid.New()
		product := models_mothers.NewProductMother().MustBuild()
		expectedError := errors.New("Database connection failed")

		mockRepo.SetupUpdateError(productID, product, expectedError)

		// Act
		err := useCase.Execute(ctx, productID, product)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}
