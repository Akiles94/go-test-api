package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/dto"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/domain/models"
)

type RegisterUseCasePort interface {
	Execute(ctx context.Context, registerDto dto.RegisterRequestDto) (*models.User, error)
}
