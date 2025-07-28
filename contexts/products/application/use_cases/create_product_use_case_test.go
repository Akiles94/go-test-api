package use_cases

import (
	"context"
	"testing"

	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
	shared_models "github.com/Akiles94/go-test-api/contexts/shared/domain/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductUseCase_Execute(t *testing.T) {
	t.Run("should create product successfully", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := NewMockProductRepository()
		product := models.NewProductMother().MustBuild()
		useCase := NewCreateProductUseCase(mockRepo)

		mockRepo.SetupCreateSuccess(product)

		// Act
		err := useCase.Execute(ctx, product)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := NewMockProductRepository()
		useCase := NewCreateProductUseCase(mockRepo)

		product := models.NewProductMother().MustBuild()
		expectedError := shared_models.DomainError{
			Code:    "DATABASE_ERROR",
			Message: "Database connection failed",
		}

		// Using helper method for cleaner setup
		mockRepo.SetupCreateError(product, expectedError)

		// Act
		err := useCase.Execute(ctx, product)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}
