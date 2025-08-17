package dto_tests

import (
	"testing"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategoryRequest_ValidRequest(t *testing.T) {
	// Arrange
	request := dto.CreateCategoryRequest{
		Name:        "Electronics",
		Description: "Electronic devices and accessories",
		IsActive:    true,
	}

	// Act & Assert
	assert.Equal(t, "Electronics", request.Name)
	assert.Equal(t, "Electronic devices and accessories", request.Description)
	assert.True(t, request.IsActive)
}

func TestCreateCategoryRequest_EmptyFields(t *testing.T) {
	// Arrange
	request := dto.CreateCategoryRequest{
		Name:        "",
		Description: "",
		IsActive:    false,
	}

	// Act & Assert
	assert.Empty(t, request.Name)
	assert.Empty(t, request.Description)
	assert.False(t, request.IsActive)
}
