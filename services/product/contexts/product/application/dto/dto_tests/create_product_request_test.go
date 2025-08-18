package dto_tests

import (
	"testing"

	"github.com/Akiles94/go-test-api/services/product/contexts/product/application/dto"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductRequest_ToDomainModel(t *testing.T) {
	t.Run("should convert to domain model", func(t *testing.T) {
		// Arrange
		request := dto.CreateProductRequest{
			Sku:        "DEFAULT-001",
			Name:       "Default Product",
			CategoryID: uuid.New(),
			Price:      99.99,
		}

		// Act
		domainModel, _ := request.ToDomainModel()

		// Assert
		assert.Equal(t, request.Sku, domainModel.Sku())
		assert.Equal(t, request.Name, domainModel.Name())
		assert.Equal(t, request.CategoryID, domainModel.CategoryID())
		assert.Equal(t, decimal.NewFromFloat(request.Price), domainModel.Price())
	})
}
