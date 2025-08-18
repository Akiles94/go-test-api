package models_mothers

import (
	"github.com/Akiles94/go-test-api/services/product/contexts/product/domain/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductMother struct {
	Id         uuid.UUID
	Sku        string
	Name       string
	CategoryID uuid.UUID
	Category   *models.Category
	Price      decimal.Decimal
}

func NewProductMother() *ProductMother {
	category := models.NewCategory(uuid.New(), "Default Category", "default-category", "default-category")
	return &ProductMother{
		Id:         uuid.New(),
		Sku:        "DEFAULT-001",
		Name:       "Default Product",
		CategoryID: category.ID(),
		Category:   &category,
		Price:      decimal.NewFromFloat(99.99),
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

func (pm *ProductMother) WithCategoryID(categoryID uuid.UUID) *ProductMother {
	pm.CategoryID = categoryID
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
	return models.NewProduct(pm.Id, pm.Sku, pm.Name, pm.CategoryID, pm.Price, pm.Category)
}

func (pm *ProductMother) MustBuild() models.Product {
	product, err := pm.Build()
	if err != nil {
		panic("ProductMother.MustBuild failed: " + err.Error())
	}
	return product
}
