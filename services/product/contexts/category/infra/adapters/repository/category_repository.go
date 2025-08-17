package repository

import (
	"context"
	"errors"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/outbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const oneMore = 1
const defaultLimit = 10

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) outbound.CategoryRepositoryPort {
	return &CategoryRepository{db: db}
}

func (cr *CategoryRepository) Create(ctx context.Context, category models.Category) error {
	entity := CategoryEntityFromDomain(category)
	if err := cr.db.WithContext(ctx).Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CategoryRepository) GetByID(ctx context.Context, id uuid.UUID) (models.Category, error) {
	var entity CategoryEntity
	if err := cr.db.WithContext(ctx).Where("id = ?", id).First(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return entity.ToDomain()
}

func (cr *CategoryRepository) GetAll(ctx context.Context, cursor *string, limit int) ([]models.Category, *string, error) {
	var entities []CategoryEntity
	handledLimit := limit
	if handledLimit <= 0 {
		handledLimit = defaultLimit
	}

	query := cr.db.WithContext(ctx).Order("id ASC").Limit(handledLimit + oneMore)
	if cursor != nil {
		parsedCursor, err := uuid.Parse(*cursor)
		if err != nil {
			return nil, nil, err
		}
		query = query.Where("id > ?", parsedCursor)
	}

	if err := query.Find(&entities).Error; err != nil {
		return nil, nil, err
	}

	var nextCursor *string
	if len(entities) > handledLimit {
		lastItemInPage := entities[handledLimit-1]
		cursorStr := lastItemInPage.ID.String()
		nextCursor = &cursorStr
		entities = entities[:handledLimit]
	}

	var categoryModels []models.Category
	for _, entity := range entities {
		model, err := entity.ToDomain()
		if err != nil {
			return nil, nil, err
		}
		categoryModels = append(categoryModels, model)
	}

	return categoryModels, nextCursor, nil
}

func (cr *CategoryRepository) Update(ctx context.Context, category models.Category) error {
	entity := CategoryEntityFromDomain(category)
	result := cr.db.WithContext(ctx).Where("id = ?", entity.ID).Updates(entity)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (cr *CategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := cr.db.WithContext(ctx).Where("id = ?", id).Delete(&CategoryEntity{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (cr *CategoryRepository) ExistsByName(ctx context.Context, name string, excludeID *uuid.UUID) (bool, error) {
	var count int64
	query := cr.db.WithContext(ctx).Model(&CategoryEntity{}).Where("name = ?", name)

	if excludeID != nil {
		query = query.Where("id != ?", *excludeID)
	}

	if err := query.Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
