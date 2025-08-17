package inbound

import (
	"context"

	"github.com/google/uuid"
)

type DeleteCategoryUseCasePort interface {
	Execute(ctx context.Context, id uuid.UUID) error
}
