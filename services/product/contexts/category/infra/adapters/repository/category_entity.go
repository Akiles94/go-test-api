package repository

import (
	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models"
	"github.com/google/uuid"
)

type CategoryEntity struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Name        string    `gorm:"type:varchar(100);not null;unique"`
	Description string    `gorm:"type:varchar(500);not null"`
	IsActive    bool      `gorm:"type:boolean;not null;default:true"`
}

func (CategoryEntity) TableName() string {
	return "categories"
}

func (ce *CategoryEntity) ToDomain() (models.Category, error) {
	return models.NewCategory(ce.ID, ce.Name, ce.Description, ce.IsActive)
}

func CategoryEntityFromDomain(category models.Category) *CategoryEntity {
	return &CategoryEntity{
		ID:          category.ID(),
		Name:        category.Name(),
		Description: category.Description(),
		IsActive:    category.IsActive(),
	}
}
