package use_cases

import (
	"context"
	"testing"

	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
	shared_models "github.com/Akiles94/go-test-api/contexts/shared/domain/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func GetOneProductUseCase_Execute(t *testing.T) {
	t.Run("should return product by ID successfully", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := NewMockProductRepository()
		useCase := NewGetOneProductUseCase(mockRepo)

		productID := uuid.New()
		expectedProduct := models.NewProductMother().MustBuild()

		mockRepo.On("GetByID", ctx, productID).Return(expectedProduct, nil)

		// Act
		product, err := useCase.Execute(ctx, productID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedProduct, product)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when product not found", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := NewMockProductRepository()
		useCase := NewGetOneProductUseCase(mockRepo)

		productID := uuid.New()
		expectedError := shared_models.DomainError{
			Code:    "NOT_FOUND",
			Message: "Product not found",
		}

		mockRepo.On("GetByID", ctx, productID).Return(nil, expectedError)

		// Act
		product, err := useCase.Execute(ctx, productID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, product)
		mockRepo.AssertExpectations(t)
	})
}
