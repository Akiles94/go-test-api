package handlers_mocks

import (
	"context"

	"github.com/Akiles94/go-test-api/contexts/products/domain/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockCreateProductUseCase struct {
	mock.Mock
}

func (m *MockCreateProductUseCase) Execute(ctx context.Context, product models.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

type MockUpdateProductUseCase struct {
	mock.Mock
}

func (m *MockUpdateProductUseCase) Execute(ctx context.Context, id uuid.UUID, product models.Product) error {
	args := m.Called(ctx, id, product)
	return args.Error(0)
}

type MockPatchProductUseCase struct {
	mock.Mock
}

func (m *MockPatchProductUseCase) Execute(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	args := m.Called(ctx, id, updates)
	return args.Error(0)
}

type MockDeleteProductUseCase struct {
	mock.Mock
}

func (m *MockDeleteProductUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type MockGetAllProductsUseCase struct {
	mock.Mock
}

func (m *MockGetAllProductsUseCase) Execute(ctx context.Context, cursor *string, limit *int) ([]models.Product, *string, error) {
	args := m.Called(ctx, cursor, limit)
	return args.Get(0).([]models.Product), args.Get(1).(*string), args.Error(2)
}

type MockGetOneProductUseCase struct {
	mock.Mock
}

func (m *MockGetOneProductUseCase) Execute(ctx context.Context, id uuid.UUID) (models.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(models.Product), args.Error(1)
}
