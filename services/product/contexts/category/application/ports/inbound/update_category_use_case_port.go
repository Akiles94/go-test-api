package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/google/uuid"
)

type UpdateCategoryUseCasePort interface {
	Execute(ctx context.Context, id uuid.UUID, request dto.UpdateCategoryRequest) (*dto.CategoryResponse, error)
}
