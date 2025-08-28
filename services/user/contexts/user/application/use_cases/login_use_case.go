package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/dto"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/ports/outbound"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
)

var (
	errUserNotFound = value_objects.DomainError{
		Code:    "USER_NOT_FOUND",
		Message: "User not found",
	}
	errInvalidPassword = value_objects.DomainError{
		Code:    "INVALID_CREDENTIALS",
		Message: "Invalid credentials",
	}
)

type LoginUseCase struct {
	repo       outbound.UserRepositoryPort
	passHasher outbound.PasswordHasherPort
	authorizer shared_ports.AuthServicePort
}

func NewLoginUseCase(repo outbound.UserRepositoryPort, passHasher outbound.PasswordHasherPort, authorizer shared_ports.AuthServicePort) *LoginUseCase {
	return &LoginUseCase{
		repo:       repo,
		passHasher: passHasher,
		authorizer: authorizer,
	}
}

func (luc *LoginUseCase) Execute(ctx context.Context, registerDto dto.RegisterRequestDto) (*dto.LoginResponseDto, error) {
	email := registerDto.Email
	user, err := luc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	userPass := user.PasswordHash()
	if isValid := luc.passHasher.ValidatePassword(registerDto.Password, userPass); !isValid {
		return nil, errInvalidPassword
	}
	userName := user.Name()
	token, err := luc.authorizer.GenerateToken(ctx, user.ID(), &email, &userName)
	if err != nil {
		return nil, err
	}
	refreshToken, err := luc.authorizer.GenerateRefreshToken(ctx, user.ID(), email)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponseDto{
		Token:        token.AccessToken,
		RefreshToken: refreshToken.AccessToken,
	}, nil
}
