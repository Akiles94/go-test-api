package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/dto"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/ports/outbound"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/domain/models"
	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
	"github.com/google/uuid"
)

var (
	errUserAlreadyExists = value_objects.DomainError{
		Code:    "USER_ALREADY_EXISTS",
		Message: "User already exists",
	}
)

type RegisterUseCase struct {
	repo       outbound.UserRepositoryPort
	passHasher outbound.PasswordHasherPort
}

func NewRegisterUseCase(repo outbound.UserRepositoryPort, passHasher outbound.PasswordHasherPort) *RegisterUseCase {
	return &RegisterUseCase{
		repo:       repo,
		passHasher: passHasher,
	}
}

func (ruc *RegisterUseCase) Execute(ctx context.Context, registerDto dto.RegisterRequestDto) (*models.User, error) {
	existingUser, _ := ruc.repo.GetUserByEmail(ctx, registerDto.Email)
	if existingUser != nil {
		return nil, errUserAlreadyExists
	}

	userId := uuid.New()
	passwordHash, err := ruc.passHasher.HashPassword(registerDto.Password)
	if err != nil {
		return nil, err
	}
	userRole := models.NewRole(registerDto.Role)
	user, err := models.NewUser(userId, registerDto.Name, registerDto.LastName, registerDto.Email, registerDto.Password, passwordHash, userRole)
	if err != nil {
		return nil, err
	}

	if err := ruc.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return &user, nil
}
