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
	"github.com/stretchr/testify/mock"
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
		mockRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(user models.User) bool {
			return user.Name() == request.Name &&
				user.LastName() == request.LastName &&
				user.Email() == request.Email &&
				user.Role() == models.Role(request.Role)
		})).Return(nil)

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		require.NoError(t, err)
		mockRepo.AssertExpectations(t)
		mockHasher.AssertExpectations(t)
	})

	t.Run("should return error when password is too short", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "John",
			LastName: "Doe",
			Email:    "john.doe@example.com",
			Password: "weak", // Password too short
			Role:     int(models.RoleUser),
		}

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, models.ErrPassTooShort, err)

		mockHasher.AssertNotCalled(t, "HashPassword")
		mockRepo.AssertNotCalled(t, "CreateUser")
	})

	t.Run("should return error when password has no uppercase", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "John",
			LastName: "Doe",
			Email:    "john.doe@example.com",
			Password: "mypassword123!", // No uppercase
			Role:     int(models.RoleUser),
		}

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, models.ErrPassNoUppercase, err)

		mockHasher.AssertNotCalled(t, "HashPassword")
		mockRepo.AssertNotCalled(t, "CreateUser")
	})

	t.Run("should return error when password has no lowercase", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "John",
			LastName: "Doe",
			Email:    "john.doe@example.com",
			Password: "MYPASSWORD123!", // No lowercase
			Role:     int(models.RoleUser),
		}

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, models.ErrPassNoLowercase, err)

		mockHasher.AssertNotCalled(t, "HashPassword")
		mockRepo.AssertNotCalled(t, "CreateUser")
	})

	t.Run("should return error when password has no number", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "John",
			LastName: "Doe",
			Email:    "john.doe@example.com",
			Password: "MyPassword!", // No number
			Role:     int(models.RoleUser),
		}

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, models.ErrPassNoNumber, err)

		mockHasher.AssertNotCalled(t, "HashPassword")
		mockRepo.AssertNotCalled(t, "CreateUser")
	})

	t.Run("should return error when password has no special character", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "John",
			LastName: "Doe",
			Email:    "john.doe@example.com",
			Password: "MyPassword123", // No special character
			Role:     int(models.RoleUser),
		}

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, models.ErrPassNoSpecialChar, err)

		mockHasher.AssertNotCalled(t, "HashPassword")
		mockRepo.AssertNotCalled(t, "CreateUser")
	})

	t.Run("should return error when role is invalid", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "John",
			LastName: "Doe",
			Email:    "john.doe@example.com",
			Password: "MySecurePass123!",
			Role:     999, // Invalid role
		}

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, models.ErrInvalidRole, err)

		mockHasher.AssertNotCalled(t, "HashPassword")
		mockRepo.AssertNotCalled(t, "CreateUser")
	})

	t.Run("should return error when name is empty", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "", // Empty name
			LastName: "Doe",
			Email:    "john.doe@example.com",
			Password: "MySecurePass123!",
			Role:     int(models.RoleUser),
		}

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		// Asumiendo que valida campos vacíos

		mockHasher.AssertNotCalled(t, "HashPassword")
		mockRepo.AssertNotCalled(t, "CreateUser")
	})

	t.Run("should return error when email is empty", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "John",
			LastName: "Doe",
			Email:    "", // Empty email
			Password: "MySecurePass123!",
			Role:     int(models.RoleUser),
		}

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)

		mockHasher.AssertNotCalled(t, "HashPassword")
		mockRepo.AssertNotCalled(t, "CreateUser")
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

		mockHasher.SetupHashPasswordError(request.Password, hashingError)

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

		mockHasher.SetupHashPasswordSuccess(request.Password, hashedPassword)
		mockRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("models.User")).Return(createError)

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, createError, err)

		mockRepo.AssertExpectations(t)
		mockHasher.AssertExpectations(t)
	})

	t.Run("should register admin user when role is admin", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "Admin",
			LastName: "User",
			Email:    "admin@example.com",
			Password: "AdminPass123!",
			Role:     int(models.RoleAdmin),
		}

		hashedPassword := "hashed_admin_password_123"

		mockHasher.SetupHashPasswordSuccess(request.Password, hashedPassword)
		mockRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(user models.User) bool {
			return user.Role() == models.RoleAdmin
		})).Return(nil)

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		require.NoError(t, err)
		mockRepo.AssertExpectations(t)
		mockHasher.AssertExpectations(t)
	})

	t.Run("should register guest user when role is guest", func(t *testing.T) {
		// Arrange
		mockRepo := use_cases_mocks.NewMockUserRepository()
		mockHasher := use_cases_mocks.NewMockPasswordHasher()

		request := dto.RegisterRequestDto{
			Name:     "Guest",
			LastName: "User",
			Email:    "guest@example.com",
			Password: "GuestPass123!",
			Role:     int(models.RoleGuest),
		}

		hashedPassword := "hashed_guest_password_123"

		mockHasher.SetupHashPasswordSuccess(request.Password, hashedPassword)
		mockRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(user models.User) bool {
			return user.Role() == models.RoleGuest
		})).Return(nil)

		useCase := use_cases.NewRegisterUseCase(mockRepo, mockHasher)

		// Act
		_, err := useCase.Execute(context.Background(), request)

		// Assert
		require.NoError(t, err)
		mockRepo.AssertExpectations(t)
		mockHasher.AssertExpectations(t)
	})
}
