package models_mothers

import (
	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models"
	"github.com/google/uuid"
)

type CategoryMother struct{}

func NewCategoryMother() *CategoryMother {
	return &CategoryMother{}
}

func (cm *CategoryMother) ValidCategory() models.Category {
	category, _ := models.NewCategory(
		uuid.New(),
		"Electronics",
		"Electronic devices and accessories",
		true,
	)
	return category
}

func (cm *CategoryMother) ValidCategoryWithId(id uuid.UUID) models.Category {
	category, _ := models.NewCategory(
		id,
		"Electronics",
		"Electronic devices and accessories",
		true,
	)
	return category
}

func (cm *CategoryMother) ValidCategoryWithName(name string) models.Category {
	category, _ := models.NewCategory(
		uuid.New(),
		name,
		"Category description",
		true,
	)
	return category
}

func (cm *CategoryMother) InactiveCategory() models.Category {
	category, _ := models.NewCategory(
		uuid.New(),
		"Inactive Category",
		"This category is inactive",
		false,
	)
	return category
}

func (cm *CategoryMother) CategoryWithEmptyName() (models.Category, error) {
	return models.NewCategory(
		uuid.New(),
		"",
		"Description for empty name category",
		true,
	)
}

func (cm *CategoryMother) CategoryWithEmptyDescription() (models.Category, error) {
	return models.NewCategory(
		uuid.New(),
		"Valid Name",
		"",
		true,
	)
}

func (cm *CategoryMother) CategoryWithNilId() (models.Category, error) {
	return models.NewCategory(
		uuid.Nil,
		"Valid Name",
		"Valid description",
		true,
	)
}
