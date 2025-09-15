package models

import "github.com/google/uuid"

type product struct {
	id    uuid.UUID
	price float64
	stock int
}

type Product interface {
	ID() uuid.UUID
	Price() float64
	Stock() int
}

func NewProduct(id uuid.UUID, price float64, stock int) Product {
	return &product{
		id:    id,
		price: price,
		stock: stock,
	}
}

func (p *product) ID() uuid.UUID {
	return p.id
}

func (p *product) Price() float64 {
	return p.price
}

func (p *product) Stock() int {
	return p.stock
}
