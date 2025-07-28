package use_cases

import (
	"context"
	"errors"
	"testing"

	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func UpdateProductUseCase_Execute(t *testing.T) {
	t.Run("should update product successfully", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := NewMockProductRepository()
		useCase := NewUpdateProductUseCase(mockRepo)

		productID := uuid.New()
		product := models.NewProductMother().MustBuild()
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
		mockRepo := NewMockProductRepository()
		useCase := NewUpdateProductUseCase(mockRepo)

		productID := uuid.New()
		product := models.NewProductMother().MustBuild()
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
