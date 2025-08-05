package use_cases_tests

import (
	"context"
	"testing"

	"github.com/Akiles94/go-test-api/contexts/products/application/use_cases"
	"github.com/Akiles94/go-test-api/contexts/products/application/use_cases/use_cases_mocks"
	shared_models "github.com/Akiles94/go-test-api/contexts/shared/domain/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteProductUseCase_Execute(t *testing.T) {
	t.Run("should delete product successfully", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := use_cases_mocks.NewMockProductRepository()
		useCase := use_cases.NewDeleteProductUseCase(mockRepo)

		productID := uuid.New()
		mockRepo.SetupDeleteSuccess(productID)
		// Act
		err := useCase.Execute(ctx, productID)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := use_cases_mocks.NewMockProductRepository()
		useCase := use_cases.NewDeleteProductUseCase(mockRepo)

		productID := uuid.New()
		expectedError := shared_models.DomainError{
			Code:    "DATABASE_ERROR",
			Message: "Database connection failed",
		}

		mockRepo.SetupDeleteError(productID, expectedError)

		// Act
		err := useCase.Execute(ctx, productID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}
