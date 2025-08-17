package use_cases_tests

import (
	"context"
	"testing"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/use_cases"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/use_cases/use_cases_mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCategoryUseCase_ValidRequest_ShouldReturnCategoryResponse(t *testing.T) {
	// Arrange
	mockRepo := &use_cases_mocks.MockCategoryRepository{}
	useCase := use_cases.NewCreateCategoryUseCase(mockRepo)

	request := dto.CreateCategoryRequest{
		Name:        "Electronics",
		Description: "Electronic devices and accessories",
		IsActive:    true,
	}

	mockRepo.On("ExistsByName", mock.Anything, request.Name, (*uuid.UUID)(nil)).Return(false, nil)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.category")).Return(nil)

	// Act
	response, err := useCase.Execute(context.Background(), request)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, request.Name, response.Name)
	assert.Equal(t, request.Description, response.Description)
	assert.Equal(t, request.IsActive, response.IsActive)
	mockRepo.AssertExpectations(t)
}

func TestCreateCategoryUseCase_EmptyName_ShouldReturnError(t *testing.T) {
	// Arrange
	mockRepo := &use_cases_mocks.MockCategoryRepository{}
	useCase := use_cases.NewCreateCategoryUseCase(mockRepo)

	request := dto.CreateCategoryRequest{
		Name:        "",
		Description: "Description",
		IsActive:    true,
	}

	// Act
	response, err := useCase.Execute(context.Background(), request)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, response)
	mockRepo.AssertExpectations(t)
}
