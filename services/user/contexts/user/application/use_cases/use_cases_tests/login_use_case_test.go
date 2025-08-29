package use_cases_tests

import (
	"context"
	"errors"
	"testing"

	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/dto"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/use_cases"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/use_cases/use_cases_mocks"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/domain/models/models_mothers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginUseCase_Execute(t *testing.T) {
	t.Run("should login successfully with valid credentials", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()
		mockAuth := use_cases_mocks.NewMockAuthService()

		user := models_mothers.NewUserMother().MustBuild()
		email := user.Email()
		password := "Password123!"
		userName := user.Name()

		expectedToken := "access_token_123"
		expectedRefreshToken := "refresh_token_123"

		mockRepo.SetupGetUserByEmailSuccess(email, &user)
		mockHasher.SetupValidatePasswordTrue(password, user.Password())
		mockAuth.SetupGenerateTokenSuccess(user.ID(), &email, &userName, expectedToken)
		mockAuth.SetupGenerateRefreshTokenSuccess(user.ID(), email, expectedRefreshToken)

		useCase := use_cases.NewLoginUseCase(mockRepo, mockHasher, mockAuth)

		request := dto.RegisterRequestDto{
			Email:    email,
			Password: password,
		}

		// Act
		response, err := useCase.Execute(context.Background(), request)

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, expectedToken, response.Token)
		assert.Equal(t, expectedRefreshToken, response.RefreshToken)

		mockRepo.AssertExpectations(t)
		mockHasher.AssertExpectations(t)
		mockAuth.AssertExpectations(t)
	})

	t.Run("should return error when user not found", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()
		mockAuth := use_cases_mocks.NewMockAuthService()

		email := "nonexistent@example.com"
		password := "Password123!"

		mockRepo.SetupGetUserByEmailNotFound(email, errors.New("User not found"))

		useCase := use_cases.NewLoginUseCase(mockRepo, mockHasher, mockAuth)

		request := dto.RegisterRequestDto{
			Email:    email,
			Password: password,
		}

		// Act
		response, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Contains(t, err.Error(), "User not found")

		mockRepo.AssertExpectations(t)
		mockHasher.AssertNotCalled(t, "ValidatePassword")
		mockAuth.AssertNotCalled(t, "GenerateToken")
		mockAuth.AssertNotCalled(t, "GenerateRefreshToken")
	})

	t.Run("should return error when password is invalid", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()
		mockAuth := use_cases_mocks.NewMockAuthService()

		user := models_mothers.NewUserMother().MustBuild()
		email := user.Email()
		wrongPassword := "WrongPassword123!"

		mockRepo.SetupGetUserByEmailSuccess(email, &user)
		mockHasher.SetupValidatePasswordFalse(wrongPassword, user.Password())

		useCase := use_cases.NewLoginUseCase(mockRepo, mockHasher, mockAuth)

		request := dto.RegisterRequestDto{
			Email:    email,
			Password: wrongPassword,
		}

		// Act
		response, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Contains(t, err.Error(), "Invalid credentials")

		mockRepo.AssertExpectations(t)
		mockHasher.AssertExpectations(t)
		mockAuth.AssertNotCalled(t, "GenerateToken")
		mockAuth.AssertNotCalled(t, "GenerateRefreshToken")
	})

	t.Run("should return error when token generation fails", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()
		mockAuth := use_cases_mocks.NewMockAuthService()

		user := models_mothers.NewUserMother().MustBuild()
		email := user.Email()
		password := "Password123!"
		userName := user.Name()

		tokenError := errors.New("token generation failed")

		mockRepo.SetupGetUserByEmailSuccess(email, &user)
		mockHasher.SetupValidatePasswordTrue(password, user.Password())
		mockAuth.SetupGenerateTokenError(user.ID(), &email, &userName, tokenError)

		useCase := use_cases.NewLoginUseCase(mockRepo, mockHasher, mockAuth)

		request := dto.RegisterRequestDto{
			Email:    email,
			Password: password,
		}

		// Act
		response, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, tokenError, err)

		mockRepo.AssertExpectations(t)
		mockHasher.AssertExpectations(t)
		mockAuth.AssertExpectations(t)
		mockAuth.AssertNotCalled(t, "GenerateRefreshToken")
	})

	t.Run("should return error when refresh token generation fails", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()
		mockAuth := use_cases_mocks.NewMockAuthService()

		user := models_mothers.NewUserMother().MustBuild()
		email := user.Email()
		password := "Password123!"
		userName := user.Name()

		expectedToken := "access_token_123"
		refreshTokenError := errors.New("refresh token generation failed")

		mockRepo.SetupGetUserByEmailSuccess(email, &user)
		mockHasher.SetupValidatePasswordTrue(password, user.Password())
		mockAuth.SetupGenerateTokenSuccess(user.ID(), &email, &userName, expectedToken)
		mockAuth.SetupGenerateRefreshTokenError(user.ID(), email, refreshTokenError)

		useCase := use_cases.NewLoginUseCase(mockRepo, mockHasher, mockAuth)

		request := dto.RegisterRequestDto{
			Email:    email,
			Password: password,
		}

		// Act
		response, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, refreshTokenError, err)

		mockRepo.AssertExpectations(t)
		mockHasher.AssertExpectations(t)
		mockAuth.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()
		mockAuth := use_cases_mocks.NewMockAuthService()

		email := "test@example.com"
		password := "Password123!"
		repositoryError := errors.New("database connection failed")

		mockRepo.SetupGetUserByEmailError(email, repositoryError)

		useCase := use_cases.NewLoginUseCase(mockRepo, mockHasher, mockAuth)

		request := dto.RegisterRequestDto{
			Email:    email,
			Password: password,
		}

		// Act
		response, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Contains(t, err.Error(), "database connection failed")

		mockRepo.AssertExpectations(t)
		mockHasher.AssertNotCalled(t, "ValidatePassword")
		mockAuth.AssertNotCalled(t, "GenerateToken")
		mockAuth.AssertNotCalled(t, "GenerateRefreshToken")
	})
}
