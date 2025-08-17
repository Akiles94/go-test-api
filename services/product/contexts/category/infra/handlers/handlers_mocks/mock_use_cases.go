package handlers_mocks

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockCreateCategoryUseCase struct {
	mock.Mock
}

func (m *MockCreateCategoryUseCase) Execute(ctx context.Context, request dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.CategoryResponse), args.Error(1)
}

type MockGetOneCategoryUseCase struct {
	mock.Mock
}

func (m *MockGetOneCategoryUseCase) Execute(ctx context.Context, id uuid.UUID) (*dto.CategoryResponse, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.CategoryResponse), args.Error(1)
}

type MockGetAllCategoriesUseCase struct {
	mock.Mock
}

func (m *MockGetAllCategoriesUseCase) Execute(ctx context.Context, cursor *string, limit int) (*dto.PaginatedCategoryResponse, error) {
	args := m.Called(ctx, cursor, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.PaginatedCategoryResponse), args.Error(1)
}

type MockUpdateCategoryUseCase struct {
	mock.Mock
}

func (m *MockUpdateCategoryUseCase) Execute(ctx context.Context, id uuid.UUID, request dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	args := m.Called(ctx, id, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.CategoryResponse), args.Error(1)
}

type MockPatchCategoryUseCase struct {
	mock.Mock
}

func (m *MockPatchCategoryUseCase) Execute(ctx context.Context, id uuid.UUID, request dto.PatchCategoryRequest) (*dto.CategoryResponse, error) {
	args := m.Called(ctx, id, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.CategoryResponse), args.Error(1)
}

type MockDeleteCategoryUseCase struct {
	mock.Mock
}

func (m *MockDeleteCategoryUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
