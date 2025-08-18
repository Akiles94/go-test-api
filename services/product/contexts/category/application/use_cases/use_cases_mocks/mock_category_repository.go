package use_cases_mocks

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/domain/models"
	"github.com/Akiles94/go-test-api/shared/infra/shared_handlers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

func NewMockCategoryRepository() *MockCategoryRepository {
	return &MockCategoryRepository{}
}

// Create utils
func (m *MockCategoryRepository) SetupCreateSuccess(category models.Category) *mock.Call {
	return m.On("Create", mock.Anything, category).Return(nil)
}

func (m *MockCategoryRepository) SetupCreateError(category models.Category, err error) *mock.Call {
	return m.On("Create", mock.Anything, category).Return(err)
}

// GetByID utils
func (m *MockCategoryRepository) SetupGetByIDSuccess(category models.Category) *mock.Call {
	return m.On("GetByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(category, nil)
}

func (m *MockCategoryRepository) SetupGetByIDNotFound(id uuid.UUID) *mock.Call {
	return m.On("GetByID", mock.Anything, id).Return(nil, shared_handlers.ErrorCodeNotFound)
}

func (m *MockCategoryRepository) SetupGetByIDError() *mock.Call {
	return m.On("GetByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, assert.AnError)
}

// GetAll utils
func (m *MockCategoryRepository) SetupGetAllSuccess(categories []models.Category, nextCursor *string) *mock.Call {
	return m.On("GetAll", mock.Anything, mock.AnythingOfType("*string"), mock.AnythingOfType("int")).Return(categories, nextCursor, nil)
}

func (m *MockCategoryRepository) SetupGetAllError(err error) *mock.Call {
	return m.On("GetAll", mock.Anything, mock.AnythingOfType("*string"), mock.AnythingOfType("int")).Return(nil, nil, err)
}

// Update utils
func (m *MockCategoryRepository) SetupUpdateSuccess(category models.Category) *mock.Call {
	return m.On("Update", mock.Anything, category).Return(nil)
}

func (m *MockCategoryRepository) SetupUpdateError(category models.Category, err error) *mock.Call {
	return m.On("Update", mock.Anything, category).Return(err)
}

// Delete utils
func (m *MockCategoryRepository) SetupDeleteSuccess() *mock.Call {
	return m.On("Delete", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil)
}

func (m *MockCategoryRepository) SetupDeleteError() *mock.Call {
	return m.On("Delete", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(assert.AnError)
}

// ExistsByName utils
func (m *MockCategoryRepository) SetupExistsByNameTrue(name string, excludeID *uuid.UUID) *mock.Call {
	return m.On("ExistsByName", mock.Anything, name, excludeID).Return(true, nil)
}

func (m *MockCategoryRepository) SetupExistsByNameFalse(name string, excludeID *uuid.UUID) *mock.Call {
	return m.On("ExistsByName", mock.Anything, name, excludeID).Return(false, nil)
}

func (m *MockCategoryRepository) SetupExistsByNameError(name string, excludeID *uuid.UUID, err error) *mock.Call {
	return m.On("ExistsByName", mock.Anything, name, excludeID).Return(false, err)
}

// Método específico para manejar strings vacíos
func (m *MockCategoryRepository) SetupExistsByNameForEmptyString() *mock.Call {
	return m.On("ExistsByName", mock.Anything, "", (*uuid.UUID)(nil)).Return(false, nil)
}

// Método genérico para cualquier string
func (m *MockCategoryRepository) SetupExistsByNameAny(exists bool, err error) *mock.Call {
	return m.On("ExistsByName", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(exists, err)
}
