package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
)

type GetAllCategoriesUseCasePort interface {
	Execute(ctx context.Context, cursor *string, limit int) (*dto.PaginatedCategoryResponse, error)
}
