package adapters

import (
	"context"
	"fmt"
	"testing"

	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	postgres_driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupPostgresDB(t *testing.T) *gorm.DB {
	t.Helper()

	ctx := context.Background()

	// Create PostgreSQL container
	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15-alpine"),
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2),
		),
	)
	require.NoError(t, err)

	// Get connection string
	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	// Connect to database
	db, err := gorm.Open(postgres_driver.Open(connStr), &gorm.Config{})
	require.NoError(t, err)

	// Auto-migrate the schema
	err = db.AutoMigrate(&ProductEntity{})
	require.NoError(t, err)

	// Cleanup when test finishes
	t.Cleanup(func() {
		postgresContainer.Terminate(ctx)
	})

	return db
}

func TestProductRepository_Create(t *testing.T) {
	t.Run("should create product successfully", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		product := models.NewProductMother().
			WithSku("POSTGRES-001").
			WithName("PostgreSQL Test Product").
			WithCategory("Database").
			WithPriceFloat(99.99).
			MustBuild()

		// Act
		err := repo.Create(ctx, product)

		// Assert
		require.NoError(t, err)

		// Verify product was created
		var entity ProductEntity
		err = db.First(&entity, product.ID()).Error
		require.NoError(t, err)
		assert.Equal(t, product.ID(), entity.ID)
		assert.Equal(t, product.Sku(), entity.Sku)
		assert.Equal(t, product.Name(), entity.Name)
		assert.Equal(t, product.Category(), entity.Category)
		assert.True(t, product.Price().Equal(entity.Price))
	})

	t.Run("should return error when creating duplicate ID", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		product := models.NewProductMother().MustBuild()

		// Create product first time
		err := repo.Create(ctx, product)
		require.NoError(t, err)

		// Act - try to create same product again
		err = repo.Create(ctx, product)

		// Assert
		assert.Error(t, err)
		// PostgreSQL should return unique constraint violation
		assert.Contains(t, err.Error(), "duplicate key")
	})
}

func TestProductRepository_GetByID(t *testing.T) {
	t.Run("should get product by ID successfully", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		// Create product first
		originalProduct := models.NewProductMother().
			WithSku("GET-001").
			WithName("Get Test Product").
			WithCategory("Test Category").
			WithPriceFloat(150.75).
			MustBuild()

		err := repo.Create(ctx, originalProduct)
		require.NoError(t, err)

		// Act
		retrievedProduct, err := repo.GetByID(ctx, originalProduct.ID())

		// Assert
		require.NoError(t, err)
		require.NotNil(t, retrievedProduct)
		assert.Equal(t, originalProduct.ID(), retrievedProduct.ID())
		assert.Equal(t, originalProduct.Sku(), retrievedProduct.Sku())
		assert.Equal(t, originalProduct.Name(), retrievedProduct.Name())
		assert.Equal(t, originalProduct.Category(), retrievedProduct.Category())
		assert.True(t, originalProduct.Price().Equal(retrievedProduct.Price()))
	})

	t.Run("should return nil when product not found", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		nonExistentID := uuid.New()

		// Act
		product, err := repo.GetByID(ctx, nonExistentID)

		// Assert
		require.NoError(t, err)
		assert.Nil(t, product)
	})
}

func TestProductRepository_GetAll(t *testing.T) {
	t.Run("should get all products without pagination", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		// Create multiple products
		products := []models.Product{
			models.NewProductMother().WithSku("GETALL-001").WithName("Product 1").MustBuild(),
			models.NewProductMother().WithSku("GETALL-002").WithName("Product 2").MustBuild(),
			models.NewProductMother().WithSku("GETALL-003").WithName("Product 3").MustBuild(),
		}

		for _, product := range products {
			err := repo.Create(ctx, product)
			require.NoError(t, err)
		}

		// Act
		result, nextCursor, err := repo.GetAll(ctx, nil, nil)

		// Assert
		require.NoError(t, err)
		assert.Len(t, result, 3)
		assert.Nil(t, nextCursor) // No pagination needed
	})

	t.Run("should handle pagination with limit", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		// Create 5 products
		for i := 1; i <= 5; i++ {
			product := models.NewProductMother().
				WithSku(fmt.Sprintf("LIMIT-%03d", i)).
				WithName(fmt.Sprintf("Product %d", i)).
				MustBuild()
			err := repo.Create(ctx, product)
			require.NoError(t, err)
		}

		limit := 3

		// Act
		result, nextCursor, err := repo.GetAll(ctx, nil, &limit)

		// Assert
		require.NoError(t, err)
		assert.Len(t, result, 3)
		assert.NotNil(t, nextCursor) // Should have next cursor
	})

	t.Run("should handle cursor pagination", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		// Create products in order
		var productIDs []uuid.UUID
		for i := 1; i <= 4; i++ {
			product := models.NewProductMother().
				WithSku(fmt.Sprintf("CURSOR-%03d", i)).
				WithName(fmt.Sprintf("Product %d", i)).
				MustBuild()
			err := repo.Create(ctx, product)
			require.NoError(t, err)
			productIDs = append(productIDs, product.ID())
		}

		// Get first page
		limit := 2
		firstPage, cursor, err := repo.GetAll(ctx, nil, &limit)
		require.NoError(t, err)
		require.Len(t, firstPage, 2)
		require.NotNil(t, cursor)

		// Act - Get second page using cursor
		secondPage, _, err := repo.GetAll(ctx, cursor, &limit)

		// Assert
		require.NoError(t, err)
		assert.Len(t, secondPage, 2)
		// Products should be different from first page
		for _, firstProduct := range firstPage {
			for _, secondProduct := range secondPage {
				assert.NotEqual(t, firstProduct.ID(), secondProduct.ID())
			}
		}
	})
}

func TestProductRepository_Update(t *testing.T) {
	t.Run("should update product successfully", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		// Create original product
		original := models.NewProductMother().
			WithSku("UPDATE-001").
			WithName("Original Name").
			WithCategory("Original Category").
			WithPriceFloat(100.00).
			MustBuild()

		err := repo.Create(ctx, original)
		require.NoError(t, err)

		// Create updated product with same ID
		updated := models.NewProductMother().
			WithID(original.ID()).
			WithSku("UPDATED-001").
			WithName("Updated Name").
			WithCategory("Updated Category").
			WithPriceFloat(200.00).
			MustBuild()

		// Act
		err = repo.Update(ctx, original.ID(), updated)

		// Assert
		require.NoError(t, err)

		// Verify update
		retrieved, err := repo.GetByID(ctx, original.ID())
		require.NoError(t, err)
		assert.Equal(t, "UPDATED-001", retrieved.Sku())
		assert.Equal(t, "Updated Name", retrieved.Name())
		assert.Equal(t, "Updated Category", retrieved.Category())
		assert.True(t, decimal.NewFromFloat(200.00).Equal(retrieved.Price()))
	})

	t.Run("should return error when updating non-existent product", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		nonExistentID := uuid.New()
		product := models.NewProductMother().WithID(nonExistentID).MustBuild()

		// Act
		err := repo.Update(ctx, nonExistentID, product)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})
}

func TestProductRepository_Patch(t *testing.T) {
	t.Run("should patch product fields successfully", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		// Create original product
		original := models.NewProductMother().
			WithSku("PATCH-001").
			WithName("Original Name").
			WithCategory("Original Category").
			WithPriceFloat(100.00).
			MustBuild()

		err := repo.Create(ctx, original)
		require.NoError(t, err)

		updates := map[string]interface{}{
			"name":  "Patched Name",
			"price": decimal.NewFromFloat(299.99),
		}

		// Act
		err = repo.Patch(ctx, original.ID(), updates)

		// Assert
		require.NoError(t, err)

		// Verify patch
		updated, err := repo.GetByID(ctx, original.ID())
		require.NoError(t, err)
		assert.Equal(t, "Patched Name", updated.Name())
		assert.True(t, decimal.NewFromFloat(299.99).Equal(updated.Price()))
		// Unchanged fields
		assert.Equal(t, original.Sku(), updated.Sku())
		assert.Equal(t, original.Category(), updated.Category())
	})

	t.Run("should patch only specified fields", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		original := models.NewProductMother().MustBuild()
		err := repo.Create(ctx, original)
		require.NoError(t, err)

		updates := map[string]interface{}{
			"sku": "NEW-SKU-001",
		}

		// Act
		err = repo.Patch(ctx, original.ID(), updates)

		// Assert
		require.NoError(t, err)

		updated, err := repo.GetByID(ctx, original.ID())
		require.NoError(t, err)
		assert.Equal(t, "NEW-SKU-001", updated.Sku())
		// Other fields unchanged
		assert.Equal(t, original.Name(), updated.Name())
		assert.Equal(t, original.Category(), updated.Category())
		assert.True(t, original.Price().Equal(updated.Price()))
	})

	t.Run("should return error when patching non-existent product", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		nonExistentID := uuid.New()
		updates := map[string]interface{}{
			"name": "New Name",
		}

		// Act
		err := repo.Patch(ctx, nonExistentID, updates)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record not found")
	})
}

func TestProductRepository_Delete(t *testing.T) {
	t.Run("should delete product successfully", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		product := models.NewProductMother().MustBuild()
		err := repo.Create(ctx, product)
		require.NoError(t, err)

		// Act
		err = repo.Delete(ctx, product.ID())

		// Assert
		require.NoError(t, err)

		// Verify deletion
		deleted, err := repo.GetByID(ctx, product.ID())
		require.NoError(t, err)
		assert.Nil(t, deleted)
	})

	t.Run("should not return error when deleting non-existent product", func(t *testing.T) {
		// Arrange
		db := setupPostgresDB(t)
		repo := NewProductRepository(db)
		ctx := context.Background()

		nonExistentID := uuid.New()

		// Act
		err := repo.Delete(ctx, nonExistentID)

		// Assert
		require.NoError(t, err) // GORM doesn't error on delete non-existent
	})
}
