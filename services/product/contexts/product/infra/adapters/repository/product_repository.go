package repository

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/product/domain/models"
	"github.com/Akiles94/go-test-api/services/product/shared/infra/adapters/repository"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const oneMore = 1
const defaultLimit = 10

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) GetAll(ctx context.Context, cursor *string, limit *int) ([]models.Product, *string, error) {
	var products []repository.ProductEntity
	handledLimit := defaultLimit
	if limit != nil {
		handledLimit = *limit
	}
	query := pr.db.WithContext(ctx).Preload("Category").Order("id ASC").Limit(handledLimit + oneMore)
	if cursor != nil {
		parsedCursor, err := uuid.Parse(*cursor)
		if err != nil {
			return nil, nil, err
		}
		query = query.Where("id > ?", parsedCursor)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, nil, err
	}

	var nextCursor *string
	if len(products) > handledLimit {
		lastItemInPage := products[handledLimit-1]
		cursorStr := lastItemInPage.ID.String()
		nextCursor = &cursorStr
		products = products[:handledLimit]
	}

	var productModels []models.Product
	for _, entity := range products {
		model := entity.ToDomainModel()
		productModels = append(productModels, *model)
	}

	return productModels, nextCursor, nil
}
func (pr *ProductRepository) GetByID(ctx context.Context, id uuid.UUID) (models.Product, error) {
	var entity repository.ProductEntity
	if err := pr.db.WithContext(ctx).Preload("Category").First(&entity, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	productModel := entity.ToDomainModel()
	return *productModel, nil
}
func (pr *ProductRepository) Create(ctx context.Context, product models.Product) error {
	productEntity := repository.NewProductEntityFromDomain(product)
	return pr.db.WithContext(ctx).Create(&productEntity).Error
}
func (pr *ProductRepository) Update(ctx context.Context, id uuid.UUID, product models.Product) error {
	var storedProduct repository.ProductEntity
	if err := pr.db.WithContext(ctx).First(&storedProduct, id).Error; err != nil {
		return err
	}
	storedProduct.Name = product.Name()
	storedProduct.Price = product.Price()
	storedProduct.Sku = product.Sku()
	storedProduct.CategoryID = product.CategoryID()
	return pr.db.WithContext(ctx).Save(storedProduct).Error
}
func (pr *ProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return pr.db.WithContext(ctx).Delete(&repository.ProductEntity{}, id).Error
}
func (pr *ProductRepository) Patch(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	var storedProduct repository.ProductEntity
	if err := pr.db.WithContext(ctx).First(&storedProduct, id).Error; err != nil {
		return err
	}

	if updates["sku"] != nil {
		storedProduct.Sku = updates["sku"].(string)
	}
	if updates["name"] != nil {
		storedProduct.Name = updates["name"].(string)
	}
	if updates["category_id"] != nil {
		storedProduct.CategoryID = updates["category_id"].(uuid.UUID)
	}
	if updates["price"] != nil {
		storedProduct.Price = updates["price"].(decimal.Decimal)
	}

	return pr.db.WithContext(ctx).Save(&storedProduct).Error
}
