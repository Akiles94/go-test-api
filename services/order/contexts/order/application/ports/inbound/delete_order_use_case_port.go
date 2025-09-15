package inbound

import "github.com/google/uuid"

type DeleteOrderUseCasePort interface {
	Execute(id uuid.UUID) error
}
