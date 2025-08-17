package dto_tests

import (
	"testing"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCategoryResponse_ValidResponse(t *testing.T) {
	// Arrange
	id := uuid.New()
	response := dto.CategoryResponse{
		ID:          id,
		Name:        "Electronics",
		Description: "Electronic devices and accessories",
		IsActive:    true,
	}

	// Act & Assert
	assert.Equal(t, id, response.ID)
	assert.Equal(t, "Electronics", response.Name)
	assert.Equal(t, "Electronic devices and accessories", response.Description)
	assert.True(t, response.IsActive)
}

func TestCategoryResponse_EmptyResponse(t *testing.T) {
	// Arrange
	response := dto.CategoryResponse{}

	// Act & Assert
	assert.Equal(t, uuid.Nil, response.ID)
	assert.Empty(t, response.Name)
	assert.Empty(t, response.Description)
	assert.False(t, response.IsActive)
}
