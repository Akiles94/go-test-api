package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	t.Run("should create product with valid data", func(t *testing.T) {
		// Arrange
		mother := NewProductMother()

		// Act
		product, err := mother.Build()

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, mother.id, product.ID())
		assert.Equal(t, mother.sku, product.Sku())
		assert.Equal(t, mother.name, product.Name())
		assert.Equal(t, mother.category, product.Category())
		assert.True(t, mother.price.Equal(product.Price()))
	})

	t.Run("should return error when price is negative", func(t *testing.T) {
		// Arrange & Act
		product, err := NewProductMother().WithPriceFloat(-1).Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, ErrProductPriceNegative, err)
	})

	t.Run("should return error when SKU is empty", func(t *testing.T) {
		// Arrange & Act
		product, err := NewProductMother().WithSku("").Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, ErrProductSkuEmpty, err)
	})

	t.Run("should return error when name is empty", func(t *testing.T) {
		// Arrange & Act
		product, err := NewProductMother().WithName("").Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, ErrProductNameEmpty, err)
	})

	t.Run("should return error when category is empty", func(t *testing.T) {
		// Arrange & Act
		product, err := NewProductMother().WithCategory("").Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, ErrProductCategoryEmpty, err)
	})

	t.Run("should return error when ID is nil", func(t *testing.T) {
		// Arrange & Act
		product, err := NewProductMother().WithID(uuid.Nil).Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, ErrProductIdNil, err)
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

		product := NewProductMother().
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
