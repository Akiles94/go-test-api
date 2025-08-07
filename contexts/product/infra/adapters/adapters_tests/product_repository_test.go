package adapters_tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/Akiles94/go-test-api/contexts/product/domain/models"
	"github.com/Akiles94/go-test-api/contexts/product/domain/models/models_mothers"
	"github.com/Akiles94/go-test-api/contexts/product/infra/adapters"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	postgres_driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ProductRepositoryTestSuite struct {
	suite.Suite
	container *postgres.PostgresContainer
	db        *gorm.DB
	repo      *adapters.ProductRepository
	ctx       context.Context
}

func (suite *ProductRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	// Create PostgreSQL container once for all tests
	container, err := postgres.Run(suite.ctx,
		"postgres:15-alpine",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2),
		),
	)
	suite.Require().NoError(err)
	suite.container = container

	// Get connection string and connect to database
	connStr, err := container.ConnectionString(suite.ctx, "sslmode=disable")
	suite.Require().NoError(err)

	db, err := gorm.Open(postgres_driver.Open(connStr), &gorm.Config{})
	suite.Require().NoError(err)

	// Auto-migrate schema
	err = db.AutoMigrate(&adapters.ProductEntity{})
	suite.Require().NoError(err)

	// Create repository singleton
	suite.repo = adapters.NewProductRepository(db)
	suite.db = db
}

func (suite *ProductRepositoryTestSuite) TearDownSuite() {
	if suite.container != nil {
		suite.container.Terminate(suite.ctx)
	}
}

func (suite *ProductRepositoryTestSuite) TearDownTest() {
	// Clean up all products after each test to ensure isolation
	// Access DB through the repository's internal DB connection
	if suite.repo != nil && suite.db != nil {
		suite.db.Exec("TRUNCATE TABLE product_entities")
	}
}

func (suite *ProductRepositoryTestSuite) TestCreate() {
	suite.Run("should create product successfully", func() {
		// Arrange
		product := models_mothers.NewProductMother().
			WithSku("POSTGRES-001").
			WithName("PostgreSQL Test Product").
			WithCategory("Database").
			WithPriceFloat(99.99).
			MustBuild()

		// Act
		err := suite.repo.Create(suite.ctx, product)

		// Assert
		suite.Require().NoError(err)

		// Verify through repository
		retrieved, err := suite.repo.GetByID(suite.ctx, product.ID())
		suite.Require().NoError(err)
		suite.NotNil(retrieved)
		suite.Equal(product.ID(), retrieved.ID())
		suite.Equal(product.Sku(), retrieved.Sku())
		suite.Equal(product.Name(), retrieved.Name())
		suite.Equal(product.Category(), retrieved.Category())
		suite.True(product.Price().Equal(retrieved.Price()))
	})

	suite.Run("should return error when creating duplicate ID", func() {
		// Arrange
		product := models_mothers.NewProductMother().MustBuild()

		// Create product first time
		err := suite.repo.Create(suite.ctx, product)
		suite.Require().NoError(err)

		// Act - try to create same product again
		err = suite.repo.Create(suite.ctx, product)

		// Assert
		suite.Error(err)
		suite.Contains(err.Error(), "duplicate key")
	})
}

func (suite *ProductRepositoryTestSuite) TestGetByID() {
	suite.Run("should get product successfully", func() {
		// Arrange
		originalProduct := models_mothers.NewProductMother().
			WithSku("GET-001").
			WithName("Get Test Product").
			WithCategory("Test Category").
			WithPriceFloat(150.75).
			MustBuild()

		err := suite.repo.Create(suite.ctx, originalProduct)
		suite.Require().NoError(err)

		// Act
		retrievedProduct, err := suite.repo.GetByID(suite.ctx, originalProduct.ID())

		// Assert
		suite.Require().NoError(err)
		suite.NotNil(retrievedProduct)
		suite.Equal(originalProduct.ID(), retrievedProduct.ID())
		suite.Equal(originalProduct.Sku(), retrievedProduct.Sku())
		suite.Equal(originalProduct.Name(), retrievedProduct.Name())
		suite.Equal(originalProduct.Category(), retrievedProduct.Category())
		suite.True(originalProduct.Price().Equal(retrievedProduct.Price()))
	})

	suite.Run("should return nil when product not found", func() {
		// Arrange
		nonExistentID := uuid.New()

		// Act
		product, err := suite.repo.GetByID(suite.ctx, nonExistentID)

		// Assert
		suite.NoError(err)
		suite.Nil(product)
	})
}

func (suite *ProductRepositoryTestSuite) TestGetAll() {
	suite.Run("should get all products without pagination", func() {
		// Arrange
		products := []models.Product{
			models_mothers.NewProductMother().WithSku("GETALL-001").WithName("Product 1").MustBuild(),
			models_mothers.NewProductMother().WithSku("GETALL-002").WithName("Product 2").MustBuild(),
			models_mothers.NewProductMother().WithSku("GETALL-003").WithName("Product 3").MustBuild(),
		}

		for _, product := range products {
			err := suite.repo.Create(suite.ctx, product)
			suite.Require().NoError(err)
		}

		// Act
		result, nextCursor, err := suite.repo.GetAll(suite.ctx, nil, nil)

		// Assert
		suite.NoError(err)
		suite.Len(result, 3)
		suite.Nil(nextCursor)
	})

	suite.Run("should handle pagination with limit", func() {
		// Arrange
		for i := 1; i <= 5; i++ {
			product := models_mothers.NewProductMother().
				WithSku(fmt.Sprintf("LIMIT-%03d", i)).
				WithName(fmt.Sprintf("Product %d", i)).
				MustBuild()
			err := suite.repo.Create(suite.ctx, product)
			suite.Require().NoError(err)
		}

		limit := 3

		// Act
		result, nextCursor, err := suite.repo.GetAll(suite.ctx, nil, &limit)

		// Assert
		suite.NoError(err)
		suite.Len(result, 3)
		suite.NotNil(nextCursor)
	})

	suite.Run("should handle cursor pagination", func() {
		// Arrange
		for i := 1; i <= 4; i++ {
			product := models_mothers.NewProductMother().
				WithSku(fmt.Sprintf("CURSOR-%03d", i)).
				WithName(fmt.Sprintf("Product %d", i)).
				MustBuild()
			err := suite.repo.Create(suite.ctx, product)
			suite.Require().NoError(err)
		}

		// Get first page
		limit := 2
		firstPage, cursor, err := suite.repo.GetAll(suite.ctx, nil, &limit)
		suite.Require().NoError(err)
		suite.Len(firstPage, 2)
		suite.NotNil(cursor)

		// Act - Get second page
		secondPage, _, err := suite.repo.GetAll(suite.ctx, cursor, &limit)

		// Assert
		suite.NoError(err)
		suite.Len(secondPage, 2)
		// Products should be different
		for _, firstProduct := range firstPage {
			for _, secondProduct := range secondPage {
				suite.NotEqual(firstProduct.ID(), secondProduct.ID())
			}
		}
	})
}

func (suite *ProductRepositoryTestSuite) TestUpdate() {
	suite.Run("should update product successfully", func() {
		// Arrange
		original := models_mothers.NewProductMother().
			WithSku("UPDATE-001").
			WithName("Original Name").
			WithCategory("Original Category").
			WithPriceFloat(100.00).
			MustBuild()

		err := suite.repo.Create(suite.ctx, original)
		suite.Require().NoError(err)

		updated := models_mothers.NewProductMother().
			WithID(original.ID()).
			WithSku("UPDATED-001").
			WithName("Updated Name").
			WithCategory("Updated Category").
			WithPriceFloat(200.00).
			MustBuild()

		// Act
		err = suite.repo.Update(suite.ctx, original.ID(), updated)

		// Assert
		suite.NoError(err)

		retrieved, err := suite.repo.GetByID(suite.ctx, original.ID())
		suite.Require().NoError(err)
		suite.Equal("UPDATED-001", retrieved.Sku())
		suite.Equal("Updated Name", retrieved.Name())
		suite.Equal("Updated Category", retrieved.Category())
		suite.True(decimal.NewFromFloat(200.00).Equal(retrieved.Price()))
	})

	suite.Run("should return error when updating non-existent product", func() {
		// Arrange
		nonExistentID := uuid.New()
		product := models_mothers.NewProductMother().WithID(nonExistentID).MustBuild()

		// Act
		err := suite.repo.Update(suite.ctx, nonExistentID, product)

		// Assert
		suite.Error(err)
		suite.Contains(err.Error(), "record not found")
	})
}

func (suite *ProductRepositoryTestSuite) TestPatch() {
	suite.Run("should patch product fields successfully", func() {
		// Arrange
		original := models_mothers.NewProductMother().
			WithSku("PATCH-001").
			WithName("Original Name").
			WithCategory("Original Category").
			WithPriceFloat(100.00).
			MustBuild()

		err := suite.repo.Create(suite.ctx, original)
		suite.Require().NoError(err)

		updates := map[string]interface{}{
			"name":  "Patched Name",
			"price": decimal.NewFromFloat(299.99),
		}

		// Act
		err = suite.repo.Patch(suite.ctx, original.ID(), updates)

		// Assert
		suite.NoError(err)

		updated, err := suite.repo.GetByID(suite.ctx, original.ID())
		suite.Require().NoError(err)
		suite.Equal("Patched Name", updated.Name())
		suite.True(decimal.NewFromFloat(299.99).Equal(updated.Price()))
		suite.Equal(original.Sku(), updated.Sku())
		suite.Equal(original.Category(), updated.Category())
	})

	suite.Run("should return error when patching non-existent product", func() {
		// Arrange
		nonExistentID := uuid.New()
		updates := map[string]interface{}{"name": "New Name"}

		// Act
		err := suite.repo.Patch(suite.ctx, nonExistentID, updates)

		// Assert
		suite.Error(err)
		suite.Contains(err.Error(), "record not found")
	})
}

func (suite *ProductRepositoryTestSuite) TestDelete() {
	suite.Run("should delete product successfully", func() {
		// Arrange
		product := models_mothers.NewProductMother().MustBuild()
		err := suite.repo.Create(suite.ctx, product)
		suite.Require().NoError(err)

		// Act
		err = suite.repo.Delete(suite.ctx, product.ID())

		// Assert
		suite.NoError(err)

		deleted, err := suite.repo.GetByID(suite.ctx, product.ID())
		suite.NoError(err)
		suite.Nil(deleted)
	})
}

func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}
