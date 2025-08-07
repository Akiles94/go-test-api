package inbound

import (
	"context"

	"github.com/google/uuid"
)

type DeleteProductUseCasePort interface {
	Execute(ctx context.Context, id uuid.UUID) error
}
