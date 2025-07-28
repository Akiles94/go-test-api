package dto

import (
	"testing"

	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductRequest_ToDomainModel(t *testing.T) {
	t.Run("should convert to domain model", func(t *testing.T) {
		// Arrange
		request := CreateProductRequest{
			Sku:      "DEFAULT-001",
			Name:     "Default Product",
			Category: "Electronics",
			Price:    99.99,
		}

		// Act
		domainModel := models.NewProductMother().MustBuild()

		// Assert
		assert.Equal(t, request.Sku, domainModel.Sku())
		assert.Equal(t, request.Name, domainModel.Name())
		assert.Equal(t, request.Category, domainModel.Category())
		assert.Equal(t, decimal.NewFromFloat(request.Price), domainModel.Price())
	})
}
