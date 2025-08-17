package models

import (
	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
	"github.com/google/uuid"
)

var (
	ErrCategoryNameEmpty = value_objects.DomainError{
		Code:    "CATEGORY_NAME_EMPTY",
		Message: "Category name cannot be empty",
	}

	ErrCategoryDescriptionEmpty = value_objects.DomainError{
		Code:    "CATEGORY_DESCRIPTION_EMPTY",
		Message: "Category description cannot be empty",
	}

	ErrCategoryIdNil = value_objects.DomainError{
		Code:    "CATEGORY_ID_NIL",
		Message: "Category ID cannot be nil",
	}
)

type Category interface {
	ID() uuid.UUID
	Name() string
	Description() string
	IsActive() bool
}

type category struct {
	id          uuid.UUID
	name        string
	description string
	isActive    bool
}

func NewCategory(id uuid.UUID, name, description string, isActive bool) (Category, error) {
	if name == "" {
		return nil, ErrCategoryNameEmpty
	}
	if description == "" {
		return nil, ErrCategoryDescriptionEmpty
	}
	if id == uuid.Nil {
		return nil, ErrCategoryIdNil
	}

	return &category{
		id:          id,
		name:        name,
		description: description,
		isActive:    isActive,
	}, nil
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

func (c *category) IsActive() bool {
	return c.isActive
}
