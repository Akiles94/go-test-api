package repository_tests

import (
	"context"

	"github.com/Akiles94/go-test-api/services/user/contexts/user/domain/models/models_mothers"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/infra/adapters/repository"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	postgres_driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	container *postgres.PostgresContainer
	db        *gorm.DB
	repo      *repository.UserRepository
	ctx       context.Context
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
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
	err = db.AutoMigrate(
		&repository.UserEntity{},
	)
	suite.Require().NoError(err)

	// Create repository singleton
	suite.repo = repository.NewUserRepository(db)
	suite.db = db
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	if suite.container != nil {
		suite.container.Terminate(suite.ctx)
	}
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	// Clean up all products after each test to ensure isolation
	// Access DB through the repository's internal DB connection
	if suite.repo != nil && suite.db != nil {
		suite.db.Exec("TRUNCATE TABLE products")
		suite.db.Exec("TRUNCATE TABLE categories CASCADE")
	}
}

func (suite *UserRepositoryTestSuite) TestCreate() {
	suite.Run("should create user successfully", func() {
		// Arrange
		user := models_mothers.NewUserMother().MustBuild()

		// Act
		err := suite.repo.CreateUser(suite.ctx, user)

		// Assert
		suite.Require().NoError(err)

		// Verify through repository
		retrieved, err := suite.repo.GetUserByEmail(suite.ctx, user.Email())
		handledRetrieved := *retrieved
		suite.Require().NoError(err)
		suite.NotNil(retrieved)
		suite.Equal(user.ID(), handledRetrieved.ID())
		suite.Equal(user.Name(), handledRetrieved.Name())
		suite.Equal(user.Email(), handledRetrieved.Email())
	})

	suite.Run("should return error when creating duplicate ID", func() {
		// Arrange
		user := models_mothers.NewUserMother().MustBuild()

		// Create user first time
		err := suite.repo.CreateUser(suite.ctx, user)
		suite.Require().NoError(err)

		// Act - try to create same user again
		err = suite.repo.CreateUser(suite.ctx, user)

		// Assert
		suite.Error(err)
		suite.Contains(err.Error(), "duplicate key")
	})
}
