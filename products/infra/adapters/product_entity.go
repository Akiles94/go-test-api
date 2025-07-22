package adapters

import (
	"github.com/Akiles94/go-test-api/products/domain/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductEntity struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Sku      string    `gorm:"unique"`
	Name     string
	Category string
	Price    decimal.Decimal `gorm:"type:decimal(10,2)"`
}

func (p *ProductEntity) ToDomainModel() *models.Product {
	product, err := models.NewProduct(
		p.ID,
		p.Sku,
		p.Name,
		p.Category,
		p.Price,
	)
	if err != nil {
		return nil
	}
	return &product
}
