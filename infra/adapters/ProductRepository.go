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

func (r *ProductRepository) GetPaginated(cursor *string, limit *int) (dto.ProductsResponse, error) {
	var products []models.Product
	handledLimit := defaultLimit
	if limit != nil {
		handledLimit = *limit
	}
	query := r.db.Order("id ASC").Limit(handledLimit + oneMore)
	if cursor != nil {
		parsedCursor, err := uuid.Parse(*cursor)
		if err != nil {
			return dto.ProductsResponse{}, err
		}
		query = query.Where("id > ?", parsedCursor)
	}

	if err := query.Find(&products).Error; err != nil {
		return dto.ProductsResponse{}, err
	}

	var nextCursor *string
	if len(products) > handledLimit {
		id := products[handledLimit].ID.String()
		nextCursor = &id
		products = products[:handledLimit]
	}

	return dto.ProductsResponse{
		Products:   products,
		NextCursor: nextCursor,
	}, nil
}
func (r *ProductRepository) GetByID(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
func (r *ProductRepository) Create(product models.Product) error {
	return r.db.Create(&product).Error
}
func (r *ProductRepository) Update(id uuid.UUID, product *models.Product) error {
	var storeProduct models.Product
	if err := r.db.First(&storeProduct, id).Error; err != nil {
		return err
	}
	return r.db.Save(product).Error
}
func (r *ProductRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Product{}, id).Error
}
func (r *ProductRepository) PatchProduct(id uuid.UUID, p *models.ProductPatch) error {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return err
	}

	if p.Sku != nil {
		product.Sku = *p.Sku
	}
	if p.Name != nil {
		product.Name = *p.Name
	}
	if p.Category != nil {
		product.Category = *p.Category
	}
	if p.Price != nil {
		product.Price = *p.Price
	}

	return r.db.Save(&product).Error
}
