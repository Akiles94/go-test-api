package models

import (
	"github.com/Akiles94/go-test-api/shared/domain/value_objects"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

var (
	ErrProductPriceNegative = value_objects.DomainError{
		Code:    "PRODUCT_PRICE_NEGATIVE",
		Message: "Product price cannot be negative",
	}

	ErrProductSkuEmpty = value_objects.DomainError{
		Code:    "PRODUCT_SKU_EMPTY",
		Message: "Product SKU cannot be empty",
	}

	ErrProductNameEmpty = value_objects.DomainError{
		Code:    "PRODUCT_NAME_EMPTY",
		Message: "Product name cannot be empty",
	}

	ErrProductCategoryEmpty = value_objects.DomainError{
		Code:    "PRODUCT_CATEGORY_EMPTY",
		Message: "Product category cannot be empty",
	}

	ErrProductIdNil = value_objects.DomainError{
		Code:    "PRODUCT_ID_NIL",
		Message: "Product ID cannot be nil",
	}
)

type Product interface {
	ID() uuid.UUID
	Sku() string
	Name() string
	CategoryID() uuid.UUID
	Category() *Category
	Price() decimal.Decimal
}

type product struct {
	id         uuid.UUID
	sku        string
	name       string
	categoryID uuid.UUID
	category   *Category
	price      decimal.Decimal
}

func NewProduct(id uuid.UUID, sku, name string, categoryID uuid.UUID, price decimal.Decimal, category *Category) (Product, error) {
	if price.IsNegative() {
		return nil, ErrProductPriceNegative
	}
	if sku == "" {
		return nil, ErrProductSkuEmpty
	}
	if name == "" {
		return nil, ErrProductNameEmpty
	}
	if categoryID == uuid.Nil {
		return nil, ErrProductCategoryEmpty
	}
	if id == uuid.Nil {
		return nil, ErrProductIdNil
	}

	return &product{
		id:         id,
		sku:        sku,
		name:       name,
		categoryID: categoryID,
		price:      price,
		category:   category,
	}, nil
}

func (p *product) ID() uuid.UUID {
	return p.id
}

func (p *product) Sku() string {
	return p.sku
}

func (p *product) Name() string {
	return p.name
}

func (p *product) CategoryID() uuid.UUID {
	return p.categoryID
}

func (p *product) Category() *Category {
	return p.category
}

func (p *product) Price() decimal.Decimal {
	return p.price
}
