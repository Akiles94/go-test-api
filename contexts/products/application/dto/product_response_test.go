package dto

import (
	"testing"

	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
	"github.com/stretchr/testify/assert"
)

func TestProductResponse_NewProductResponseFromDomainModel(t *testing.T) {
	t.Run("should convert from domain model", func(t *testing.T) {
		// Arrange
		domainModel := models.NewProductMother().MustBuild()

		// Act
		productResponse := NewProductResponseFromDomainModel(domainModel)

		// Assert
		assert.Equal(t, productResponse.Sku, domainModel.Sku())
		assert.Equal(t, productResponse.Name, domainModel.Name())
		assert.Equal(t, productResponse.Category, domainModel.Category())
		assert.Equal(t, productResponse.Price, domainModel.Price())
	})
}
