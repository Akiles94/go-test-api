package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
)

type CreateCategoryUseCasePort interface {
	Execute(ctx context.Context, request dto.CreateCategoryRequest) (*dto.CategoryResponse, error)
}
