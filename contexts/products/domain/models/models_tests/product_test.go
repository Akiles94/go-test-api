package models_tests

import (
	"testing"

	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
	"github.com/Akiles94/go-test-api/contexts/products/domain/models/models_mothers"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	t.Run("should create product with valid data", func(t *testing.T) {
		// Arrange
		mother := models_mothers.NewProductMother()

		// Act
		product, err := mother.Build()

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, mother.Id, product.ID())
		assert.Equal(t, mother.Sku, product.Sku())
		assert.Equal(t, mother.Name, product.Name())
		assert.Equal(t, mother.Category, product.Category())
		assert.True(t, mother.Price.Equal(product.Price()))
	})

	t.Run("should return error when price is negative", func(t *testing.T) {
		// Arrange & Act
		product, err := models_mothers.NewProductMother().WithPriceFloat(-1).Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, models.ErrProductPriceNegative, err)
	})

	t.Run("should return error when SKU is empty", func(t *testing.T) {
		// Arrange & Act
		product, err := models_mothers.NewProductMother().WithSku("").Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, models.ErrProductSkuEmpty, err)
	})

	t.Run("should return error when name is empty", func(t *testing.T) {
		// Arrange & Act
		product, err := models_mothers.NewProductMother().WithName("").Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, models.ErrProductNameEmpty, err)
	})

	t.Run("should return error when category is empty", func(t *testing.T) {
		// Arrange & Act
		product, err := models_mothers.NewProductMother().WithCategory("").Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, models.ErrProductCategoryEmpty, err)
	})

	t.Run("should return error when ID is nil", func(t *testing.T) {
		// Arrange & Act
		product, err := models_mothers.NewProductMother().WithID(uuid.Nil).Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, models.ErrProductIdNil, err)
	})
}

func TestProduct_Getters(t *testing.T) {
	t.Run("should return correct values from getters", func(t *testing.T) {
		// Arrange
		id := uuid.New()
		sku := "CUSTOM-001"
		name := "Custom Product"
		category := "Custom Category"
		price := decimal.NewFromFloat(123.45)

		product := models_mothers.NewProductMother().
			WithID(id).
			WithSku(sku).
			WithName(name).
			WithCategory(category).
			WithPrice(price).
			MustBuild()

		// Act & Assert
		assert.Equal(t, id, product.ID())
		assert.Equal(t, sku, product.Sku())
		assert.Equal(t, name, product.Name())
		assert.Equal(t, category, product.Category())
		assert.True(t, price.Equal(product.Price()))
	})
}
