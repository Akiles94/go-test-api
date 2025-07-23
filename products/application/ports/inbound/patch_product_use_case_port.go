package inbound

import (
	"context"

	"github.com/google/uuid"
)

type PatchProductUseCasePort interface {
	Execute(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
}
