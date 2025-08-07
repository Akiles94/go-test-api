package use_cases_tests

import (
	"context"
	"errors"
	"testing"

	"github.com/Akiles94/go-test-api/services/product/contexts/product/application/use_cases"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/application/use_cases/use_cases_mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func PatchProductUseCase_Execute(t *testing.T) {
	t.Run("should patch product successfully", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := use_cases_mocks.NewMockProductRepository()
		useCase := use_cases.NewPatchProductUseCase(mockRepo)

		productID := uuid.New()
		patchData := map[string]interface{}{"name": "Updated Product Name"}

		// Act
		err := useCase.Execute(ctx, productID, patchData)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := use_cases_mocks.NewMockProductRepository()
		useCase := use_cases.NewPatchProductUseCase(mockRepo)

		productID := uuid.New()
		patchData := map[string]interface{}{"name": "Updated Product Name"}
		expectedError := errors.New("Database connection failed")

		mockRepo.SetupPatchError(productID, expectedError)

		// Act
		err := useCase.Execute(ctx, productID, patchData)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}
