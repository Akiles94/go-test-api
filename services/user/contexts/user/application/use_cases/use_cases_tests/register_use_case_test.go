package use_cases_tests

import (
	"context"
	"errors"
	"testing"

	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/dto"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/use_cases"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/use_cases/use_cases_mocks"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterUseCase_Execute(t *testing.T) {
	t.Run("should register user successfully with valid data", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "John",
			LastName: "Doe",
			Email:    "john.doe@example.com",
			Password: "MySecurePass123!",
			Role:     int(models.RoleUser),
		}

		hashedPassword := "hashed_password_123"

		mockHasher.SetupHashPasswordSuccess(request.Password, hashedPassword)
		mockRepo.SetupGetUserByEmailError(request.Email, errors.New("User not found"))
		mockRepo.SetupCreateUserSuccess()

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		require.NoError(t, err)
		mockRepo.AssertExpectations(t)
		mockHasher.AssertExpectations(t)
	})

	t.Run("should return error when password hashing fails", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "John",
			LastName: "Doe",
			Email:    "john.doe@example.com",
			Password: "MySecurePass123!",
			Role:     int(models.RoleUser),
		}

		hashingError := errors.New("hashing failed")

		mockRepo.SetupCreateUserSuccess()
		mockHasher.SetupHashPasswordError(request.Password, hashingError)
		mockRepo.SetupGetUserByEmailError(request.Email, errors.New("User not found"))
		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, hashingError, err)

		mockHasher.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "CreateUser")
	})

	t.Run("should return error when repository create fails", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "John",
			LastName: "Doe",
			Email:    "john.doe@example.com",
			Password: "MySecurePass123!",
			Role:     int(models.RoleUser),
		}

		hashedPassword := "hashed_password_123"
		createError := errors.New("database connection failed")

		mockRepo.SetupGetUserByEmailError(request.Email, errors.New("User not found"))
		mockHasher.SetupHashPasswordSuccess(request.Password, hashedPassword)
		mockRepo.SetupCreateUserError(createError)

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, createError, err)

		mockRepo.AssertExpectations(t)
		mockHasher.AssertExpectations(t)
	})
}
