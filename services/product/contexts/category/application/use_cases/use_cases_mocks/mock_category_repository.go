package use_cases_mocks

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Create(ctx context.Context, category models.Category) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockCategoryRepository) GetByID(ctx context.Context, id uuid.UUID) (models.Category, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(models.Category), args.Error(1)
}

func (m *MockCategoryRepository) GetAll(ctx context.Context, cursor *string, limit int) ([]models.Category, *string, error) {
	args := m.Called(ctx, cursor, limit)
	if args.Get(0) == nil {
		return nil, args.Get(1).(*string), args.Error(2)
	}
	return args.Get(0).([]models.Category), args.Get(1).(*string), args.Error(2)
}

func (m *MockCategoryRepository) Update(ctx context.Context, category models.Category) error {
	args := m.Called(ctx, category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockCategoryRepository) ExistsByName(ctx context.Context, name string, excludeID *uuid.UUID) (bool, error) {
	args := m.Called(ctx, name, excludeID)
	return args.Bool(0), args.Error(1)
}
