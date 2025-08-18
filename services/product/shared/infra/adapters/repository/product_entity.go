package repository

import (
	"github.com/Akiles94/go-test-api/services/product/contexts/product/domain/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductEntity struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Sku        string    `gorm:"unique"`
	Name       string
	Price      decimal.Decimal `gorm:"type:decimal(10,2)"`
	CategoryID uuid.UUID
	Category   *CategoryEntity `gorm:"foreignKey:CategoryID;references:ID"`
}

func (ProductEntity) TableName() string {
	return "products"
}

func NewProductEntityFromDomain(product models.Product) *ProductEntity {
	return &ProductEntity{
		ID:         product.ID(),
		Sku:        product.Sku(),
		Name:       product.Name(),
		Price:      product.Price(),
		CategoryID: product.CategoryID(),
	}
}

func (p *ProductEntity) ToDomainModel() *models.Product {
	var category *models.Category
	if p.Category != nil {
		newCategory := models.NewCategory(
			p.Category.ID,
			p.Category.Name,
			p.Category.Description,
			p.Category.Slug,
		)
		category = &newCategory
	}

	product, err := models.NewProduct(
		p.ID,
		p.Sku,
		p.Name,
		p.CategoryID,
		p.Price,
		category,
	)
	if err != nil {
		return nil
	}
	return &product
}
