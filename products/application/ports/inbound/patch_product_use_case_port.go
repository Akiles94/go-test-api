package inbound

import (
	"github.com/google/uuid"
)

type PatchProductUseCasePort interface {
	Execute(id uuid.UUID, updates map[string]interface{}) error
}
