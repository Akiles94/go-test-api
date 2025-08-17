package repository_tests

import (
	"testing"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models/models_mothers"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/infra/adapters/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCategoryEntity_ToDomain_ValidEntity_ShouldReturnCategory(t *testing.T) {
	// Arrange
	id := uuid.New()
	entity := repository.CategoryEntity{
		ID:          id,
		Name:        "Electronics",
		Description: "Electronic devices and accessories",
		IsActive:    true,
	}

	// Act
	category, err := entity.ToDomain()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, id, category.ID())
	assert.Equal(t, "Electronics", category.Name())
	assert.Equal(t, "Electronic devices and accessories", category.Description())
	assert.True(t, category.IsActive())
}

func TestCategoryEntity_ToDomain_EmptyName_ShouldReturnError(t *testing.T) {
	// Arrange
	entity := repository.CategoryEntity{
		ID:          uuid.New(),
		Name:        "",
		Description: "Description",
		IsActive:    true,
	}

	// Act
	category, err := entity.ToDomain()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, category)
}

func TestCategoryEntityFromDomain_ValidCategory_ShouldReturnEntity(t *testing.T) {
	// Arrange
	mother := models_mothers.NewCategoryMother()
	category := mother.ValidCategory()

	// Act
	entity := repository.CategoryEntityFromDomain(category)

	// Assert
	assert.NotNil(t, entity)
	assert.Equal(t, category.ID(), entity.ID)
	assert.Equal(t, category.Name(), entity.Name)
	assert.Equal(t, category.Description(), entity.Description)
	assert.Equal(t, category.IsActive(), entity.IsActive)
}

func TestCategoryEntity_TableName_ShouldReturnCorrectTableName(t *testing.T) {
	// Arrange
	entity := repository.CategoryEntity{}

	// Act
	tableName := entity.TableName()

	// Assert
	assert.Equal(t, "categories", tableName)
}
