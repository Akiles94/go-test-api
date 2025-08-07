package use_cases_tests

import (
	"context"
	"testing"

	"github.com/Akiles94/go-test-api/services/product/contexts/product/application/use_cases"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/application/use_cases/use_cases_mocks"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/domain/models/models_mothers"
	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductUseCase_Execute(t *testing.T) {
	t.Run("should create product successfully", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := use_cases_mocks.NewMockProductRepository()
		product := models_mothers.NewProductMother().MustBuild()
		useCase := use_cases.NewCreateProductUseCase(mockRepo)

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
		mockRepo := use_cases_mocks.NewMockProductRepository()
		useCase := use_cases.NewCreateProductUseCase(mockRepo)

		product := models_mothers.NewProductMother().MustBuild()
		expectedError := value_objects.DomainError{
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
