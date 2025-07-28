package use_cases

import (
	"context"
	"errors"
	"testing"

	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
	"github.com/stretchr/testify/assert"
)

func GetAllProductsUseCase_Execute(t *testing.T) {
	t.Run("should get all products successfully", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := NewMockProductRepository()
		useCase := NewGetAllProductsUseCase(mockRepo)

		cursor := "someCursor"
		limit := 10

		expectedProducts := []models.Product{
			models.NewProductMother().MustBuild(),
			models.NewProductMother().MustBuild(),
		}
		nextCursor := "nextCursor"

		mockRepo.On("GetAll", ctx, &cursor, &limit).Return(expectedProducts, &nextCursor, nil)

		// Act
		products, next, err := useCase.Execute(ctx, &cursor, &limit)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedProducts, products)
		assert.Equal(t, &nextCursor, next)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := NewMockProductRepository()
		useCase := NewGetAllProductsUseCase(mockRepo)

		cursor := "someCursor"
		limit := 10
		expectedError := errors.New("Database connection failed")

		mockRepo.On("GetAll", ctx, &cursor, &limit).Return(nil, nil, expectedError)

		// Act
		products, next, err := useCase.Execute(ctx, &cursor, &limit)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, products)
		assert.Nil(t, next)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}
