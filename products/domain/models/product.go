package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product interface {
	ID() uuid.UUID
	Sku() string
	Name() string
	Category() string
	Price() decimal.Decimal
	SetPrice(decimal.Decimal) error
	SetName(string) error
	SetCategory(string) error
}

type product struct {
	id       uuid.UUID
	sku      string
	name     string
	category string
	price    decimal.Decimal
}

func NewProduct(id uuid.UUID, sku, name, category string, price decimal.Decimal) (Product, error) {
	if price.IsNegative() {
		return nil, errors.New("price cannot be negative")
	}
	if sku == "" {
		return nil, errors.New("sku cannot be empty")
	}
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if category == "" {
		return nil, errors.New("category cannot be empty")
	}
	if id == uuid.Nil {
		return nil, errors.New("id cannot be nil")
	}

	return &product{
		id:       id,
		sku:      sku,
		name:     name,
		category: category,
		price:    price,
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

func (p *product) Category() string {
	return p.category
}

func (p *product) Price() decimal.Decimal {
	return p.price
}

func (p *product) SetPrice(price decimal.Decimal) error {
	if price.IsNegative() {
		return errors.New("price cannot be negative")
	}
	p.price = price
	return nil
}

func (p *product) SetName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	p.name = name
	return nil
}

func (p *product) SetCategory(category string) error {
	if category == "" {
		return errors.New("category cannot be empty")
	}
	p.category = category
	return nil
}
