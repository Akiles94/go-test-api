package dto_tests

import (
	"testing"

	"github.com/Akiles94/go-test-api/services/product/contexts/product/application/dto"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/domain/models/models_mothers"
	"github.com/stretchr/testify/assert"
)

func TestProductResponse_NewProductResponseFromDomainModel(t *testing.T) {
	t.Run("should convert from domain model", func(t *testing.T) {
		// Arrange
		domainModel := models_mothers.NewProductMother().MustBuild()

		// Act
		productResponse := dto.NewProductResponseFromDomainModel(domainModel)
		domainModelCategory := *domainModel.Category()
		// Assert
		assert.Equal(t, productResponse.Sku, domainModel.Sku())
		assert.Equal(t, productResponse.Name, domainModel.Name())
		assert.Equal(t, productResponse.Category.Name, domainModelCategory.Name())
		assert.Equal(t, productResponse.Category.Description, domainModelCategory.Description())
		assert.Equal(t, productResponse.Category.Slug, domainModelCategory.Slug())
		assert.Equal(t, productResponse.Price, domainModel.Price())
	})
}
