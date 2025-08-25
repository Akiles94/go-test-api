package inbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/dto"
)

type LoginUseCasePort interface {
	Execute(ctx context.Context, loginDto dto.LoginRequestDto) (*dto.LoginResponseDto, error)
}
