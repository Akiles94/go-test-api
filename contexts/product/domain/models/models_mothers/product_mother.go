package models_mothers

import (
	"github.com/Akiles94/go-test-api/contexts/product/domain/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductMother struct {
	Id       uuid.UUID
	Sku      string
	Name     string
	Category string
	Price    decimal.Decimal
}

func NewProductMother() *ProductMother {
	return &ProductMother{
		Id:       uuid.New(),
		Sku:      "DEFAULT-001",
		Name:     "Default Product",
		Category: "Electronics",
		Price:    decimal.NewFromFloat(99.99),
	}
}

func (pm *ProductMother) WithID(id uuid.UUID) *ProductMother {
	pm.Id = id
	return pm
}

func (pm *ProductMother) WithSku(sku string) *ProductMother {
	pm.Sku = sku
	return pm
}

func (pm *ProductMother) WithName(name string) *ProductMother {
	pm.Name = name
	return pm
}

func (pm *ProductMother) WithCategory(category string) *ProductMother {
	pm.Category = category
	return pm
}

func (pm *ProductMother) WithPrice(price decimal.Decimal) *ProductMother {
	pm.Price = price
	return pm
}

func (pm *ProductMother) WithPriceFloat(price float64) *ProductMother {
	pm.Price = decimal.NewFromFloat(price)
	return pm
}

func (pm *ProductMother) Build() (models.Product, error) {
	return models.NewProduct(pm.Id, pm.Sku, pm.Name, pm.Category, pm.Price)
}

func (pm *ProductMother) MustBuild() models.Product {
	product, err := pm.Build()
	if err != nil {
		panic("ProductMother.MustBuild failed: " + err.Error())
	}
	return product
}
