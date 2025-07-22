package inbound

import (
	"github.com/google/uuid"
)

type DeleteProductUseCasePort interface {
	Execute(id uuid.UUID) error
}
