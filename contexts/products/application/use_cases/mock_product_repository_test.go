package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
	"github.com/Akiles94/go-test-api/contexts/shared/infra/handlers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(ctx context.Context, product models.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) GetByID(ctx context.Context, id uuid.UUID) (models.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductRepository) GetAll(ctx context.Context, cursor *string, limit *int) ([]models.Product, *string, error) {
	args := m.Called(ctx, cursor, limit)
	if args.Get(0) == nil {
		return nil, args.Get(1).(*string), args.Error(2)
	}
	return args.Get(0).([]models.Product), args.Get(1).(*string), args.Error(2)
}

func (m *MockProductRepository) Update(ctx context.Context, id uuid.UUID, product models.Product) error {
	args := m.Called(ctx, id, product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockProductRepository) Patch(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	args := m.Called(ctx, id, updates)
	return args.Error(0)
}

func NewMockProductRepository() *MockProductRepository {
	return &MockProductRepository{}
}

func (m *MockProductRepository) SetupCreateSuccess(product models.Product) *mock.Call {
	return m.On("Create", mock.Anything, product).Return(nil)
}

func (m *MockProductRepository) SetupCreateError(product models.Product, err error) *mock.Call {
	return m.On("Create", mock.Anything, product).Return(err)
}

func (m *MockProductRepository) SetupGetByIDNotFound(id uuid.UUID) *mock.Call {
	return m.On("GetByID", mock.Anything, id).Return(nil, handlers.ErrorCodeNotFound)
}

func (m *MockProductRepository) SetupGetByIDError(err error) *mock.Call {
	return m.On("GetByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, err)
}

func (m *MockProductRepository) SetupDeleteSuccess(id uuid.UUID) *mock.Call {
	return m.On("Delete", mock.Anything, id).Return(nil)
}

func (m *MockProductRepository) SetupDeleteError(id uuid.UUID, err error) *mock.Call {
	return m.On("Delete", mock.Anything, id).Return(err)
}

func (m *MockProductRepository) SetupPatchError(id uuid.UUID, err error) *mock.Call {
	return m.On("Patch", mock.Anything, id, mock.AnythingOfType("map[string]interface {}")).Return(err)
}

func (m *MockProductRepository) SetupUpdateSuccess(id uuid.UUID, product models.Product) *mock.Call {
	return m.On("Update", mock.Anything, id, product).Return(nil)
}

func (m *MockProductRepository) SetupUpdateError(id uuid.UUID, product models.Product, err error) *mock.Call {
	return m.On("Update", mock.Anything, id, product).Return(err)
}
