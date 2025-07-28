package adapters

import (
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func ProductEntityToDomainModelTest(t *testing.T) {
	//Arrange
	productEntity := &ProductEntity{
		ID:    uuid.New(),
		Name:  "Test Product",
		Price: decimal.NewFromFloat(100),
	}

	//Act
	product := productEntity.ToDomainModel()
	if product == nil {
		t.Fatal("Expected product to be not nil")
	}
	handledProduct := *product

	//Assert
	assert.Equal(t, productEntity.ID, handledProduct.ID())
	assert.Equal(t, productEntity.Name, handledProduct.Name())
	assert.Equal(t, productEntity.Price, handledProduct.Price())
	assert.Equal(t, productEntity.Category, handledProduct.Category())
	assert.Equal(t, productEntity.Sku, handledProduct.Sku())
}
