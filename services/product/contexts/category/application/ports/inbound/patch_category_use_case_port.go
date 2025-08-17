package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/google/uuid"
)

type PatchCategoryUseCasePort interface {
	Execute(ctx context.Context, id uuid.UUID, request dto.PatchCategoryRequest) (*dto.CategoryResponse, error)
}
