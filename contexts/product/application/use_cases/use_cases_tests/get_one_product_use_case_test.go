package use_cases_tests

import (
	"context"
	"testing"

	"github.com/Akiles94/go-test-api/contexts/product/application/use_cases"
	"github.com/Akiles94/go-test-api/contexts/product/application/use_cases/use_cases_mocks"
	"github.com/Akiles94/go-test-api/contexts/product/domain/models/models_mothers"
	shared_models "github.com/Akiles94/go-test-api/contexts/shared/domain/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func GetOneProductUseCase_Execute(t *testing.T) {
	t.Run("should return product by ID successfully", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := use_cases_mocks.NewMockProductRepository()
		useCase := use_cases.NewGetOneProductUseCase(mockRepo)

		productID := uuid.New()
		expectedProduct := models_mothers.NewProductMother().MustBuild()

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
		mockRepo := use_cases_mocks.NewMockProductRepository()
		useCase := use_cases.NewGetOneProductUseCase(mockRepo)

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
