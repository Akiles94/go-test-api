package models

import (
	"github.com/google/uuid"
)

type Category interface {
	ID() uuid.UUID
	Name() string
	Description() string
	Slug() string
}

type category struct {
	id          uuid.UUID
	name        string
	description string
	slug        string
}

func NewCategory(id uuid.UUID, name, description, slug string) Category {
	return &category{
		id:          id,
		name:        name,
		description: description,
		slug:        slug,
	}
}

func (c *category) ID() uuid.UUID {
	return c.id
}

func (c *category) Name() string {
	return c.name
}

func (c *category) Description() string {
	return c.description
}

func (c *category) Slug() string {
	return c.slug
}
