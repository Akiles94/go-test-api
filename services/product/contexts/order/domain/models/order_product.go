package models

type OrderProduct interface {
	ID() string
	Price() float64
	Stock() int
}
type orderProduct struct {
	id    string
	price float64
	stock int
}

func NewOrderProduct(id string, price float64, stock int) OrderProduct {
	return &orderProduct{
		id:    id,
		price: price,
		stock: stock,
	}
}

func (p *orderProduct) ID() string {
	return p.id
}

func (p *orderProduct) Price() float64 {
	return p.price
}

func (p *orderProduct) Stock() int {
	return p.stock
}
