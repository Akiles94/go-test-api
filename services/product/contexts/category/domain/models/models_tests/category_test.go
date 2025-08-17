package models_tests

import (
	"testing"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models/models_mothers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewCategory_ValidCategory_ShouldReturnCategory(t *testing.T) {
	// Arrange
	id := uuid.New()
	name := "Electronics"
	description := "Electronic devices and accessories"
	isActive := true

	// Act
	category, err := models.NewCategory(id, name, description, isActive)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, id, category.ID())
	assert.Equal(t, name, category.Name())
	assert.Equal(t, description, category.Description())
	assert.Equal(t, isActive, category.IsActive())
}

func TestNewCategory_EmptyName_ShouldReturnError(t *testing.T) {
	// Arrange
	mother := models_mothers.NewCategoryMother()

	// Act
	category, err := mother.CategoryWithEmptyName()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Equal(t, models.ErrCategoryNameEmpty, err)
}

func TestNewCategory_EmptyDescription_ShouldReturnError(t *testing.T) {
	// Arrange
	mother := models_mothers.NewCategoryMother()

	// Act
	category, err := mother.CategoryWithEmptyDescription()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Equal(t, models.ErrCategoryDescriptionEmpty, err)
}

func TestNewCategory_NilId_ShouldReturnError(t *testing.T) {
	// Arrange
	mother := models_mothers.NewCategoryMother()

	// Act
	category, err := mother.CategoryWithNilId()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, category)
	assert.Equal(t, models.ErrCategoryIdNil, err)
}

func TestCategory_GetMethods_ShouldReturnCorrectValues(t *testing.T) {
	// Arrange
	mother := models_mothers.NewCategoryMother()
	category := mother.ValidCategory()

	// Act & Assert
	assert.NotEqual(t, uuid.Nil, category.ID())
	assert.NotEmpty(t, category.Name())
	assert.NotEmpty(t, category.Description())
	assert.True(t, category.IsActive())
}

func TestCategory_InactiveCategory_ShouldReturnFalse(t *testing.T) {
	// Arrange
	mother := models_mothers.NewCategoryMother()
	category := mother.InactiveCategory()

	// Act & Assert
	assert.False(t, category.IsActive())
}
