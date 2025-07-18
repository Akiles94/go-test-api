package adapters

import (
	"github.com/Akiles94/go-test-api/application/dto"
	"github.com/Akiles94/go-test-api/domain/models"
	"github.com/google/uuid"
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

func (pr *ProductRepository) GetPaginated(cursor *string, limit *int) (*dto.ProductsResponse, error) {
	var products []models.Product
	handledLimit := defaultLimit
	if limit != nil {
		handledLimit = *limit
	}
	query := pr.db.Order("id ASC").Limit(handledLimit + oneMore)
	if cursor != nil {
		parsedCursor, err := uuid.Parse(*cursor)
		if err != nil {
			return nil, err
		}
		query = query.Where("id > ?", parsedCursor)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	var nextCursor *string
	if len(products) > handledLimit {
		id := products[handledLimit].ID.String()
		nextCursor = &id
		products = products[:handledLimit]
	}

	return &dto.ProductsResponse{
		Products:   products,
		NextCursor: nextCursor,
	}, nil
}
func (pr *ProductRepository) GetByID(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	if err := pr.db.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}
func (pr *ProductRepository) Create(body *models.Product) error {
	id := uuid.New()
	body.ID = id
	return pr.db.Create(&body).Error
}
func (pr *ProductRepository) Update(id uuid.UUID, body models.Product) error {
	var storedProduct models.Product
	if err := pr.db.First(&storedProduct, id).Error; err != nil {
		return err
	}
	storedProduct.Name = body.Name
	storedProduct.Price = body.Price
	storedProduct.Sku = body.Sku
	storedProduct.Category = body.Category
	return pr.db.Save(storedProduct).Error
}
func (pr *ProductRepository) Delete(id uuid.UUID) error {
	return pr.db.Delete(&models.Product{}, id).Error
}
func (pr *ProductRepository) Patch(id uuid.UUID, body models.ProductPatch) error {
	var product models.Product
	if err := pr.db.First(&product, id).Error; err != nil {
		return err
	}

	if body.Sku != nil {
		product.Sku = *body.Sku
	}
	if body.Name != nil {
		product.Name = *body.Name
	}
	if body.Category != nil {
		product.Category = *body.Category
	}
	if body.Price != nil {
		product.Price = *body.Price
	}

	return pr.db.Save(&product).Error
}
