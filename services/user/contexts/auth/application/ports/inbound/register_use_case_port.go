package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/user/contexts/auth/application/dto"
)

type RegisterUseCasePort interface {
	Execute(ctx context.Context, registerDto dto.RegisterRequestDto) (*dto.RegisterResponseDto, error)
}
