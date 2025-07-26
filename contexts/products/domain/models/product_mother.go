package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductMother struct {
	id       uuid.UUID
	sku      string
	name     string
	category string
	price    decimal.Decimal
}

func NewProductMother() *ProductMother {
	return &ProductMother{
		id:       uuid.New(),
		sku:      "DEFAULT-001",
		name:     "Default Product",
		category: "Electronics",
		price:    decimal.NewFromFloat(99.99),
	}
}

func (pm *ProductMother) WithID(id uuid.UUID) *ProductMother {
	pm.id = id
	return pm
}

func (pm *ProductMother) WithSku(sku string) *ProductMother {
	pm.sku = sku
	return pm
}

func (pm *ProductMother) WithName(name string) *ProductMother {
	pm.name = name
	return pm
}

func (pm *ProductMother) WithCategory(category string) *ProductMother {
	pm.category = category
	return pm
}

func (pm *ProductMother) WithPrice(price decimal.Decimal) *ProductMother {
	pm.price = price
	return pm
}

func (pm *ProductMother) WithPriceFloat(price float64) *ProductMother {
	pm.price = decimal.NewFromFloat(price)
	return pm
}

func (pm *ProductMother) Build() (Product, error) {
	return NewProduct(pm.id, pm.sku, pm.name, pm.category, pm.price)
}

func (pm *ProductMother) MustBuild() Product {
	product, err := pm.Build()
	if err != nil {
		panic("ProductMother.MustBuild failed: " + err.Error())
	}
	return product
}
