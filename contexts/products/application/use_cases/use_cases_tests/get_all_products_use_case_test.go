package use_cases_tests

import (
	"context"
	"errors"
	"testing"

	"github.com/Akiles94/go-test-api/contexts/products/application/use_cases"
	"github.com/Akiles94/go-test-api/contexts/products/application/use_cases/use_cases_mocks"
	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
	"github.com/Akiles94/go-test-api/contexts/products/domain/models/models_mothers"
	"github.com/stretchr/testify/assert"
)

func GetAllProductsUseCase_Execute(t *testing.T) {
	t.Run("should get all products successfully", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := use_cases_mocks.NewMockProductRepository()
		useCase := use_cases.NewGetAllProductsUseCase(mockRepo)

		cursor := "someCursor"
		limit := 10

		expectedProducts := []models.Product{
			models_mothers.NewProductMother().MustBuild(),
			models_mothers.NewProductMother().MustBuild(),
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
		mockRepo := use_cases_mocks.NewMockProductRepository()
		useCase := use_cases.NewGetAllProductsUseCase(mockRepo)

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
