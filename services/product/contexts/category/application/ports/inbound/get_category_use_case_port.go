package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/google/uuid"
)

type GetCategoryUseCasePort interface {
	Execute(ctx context.Context, id uuid.UUID) (*dto.CategoryResponse, error)
}
