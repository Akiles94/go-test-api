package inbound

import (
	"github.com/Akiles94/go-test-api/products/application/dto"
	"github.com/google/uuid"
)

type PatchProductUseCasePort interface {
	Execute(id uuid.UUID, product dto.ProductPatchBody) error
}
