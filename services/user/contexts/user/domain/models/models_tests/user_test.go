package models_tests

import (
	"testing"

	"github.com/Akiles94/go-test-api/services/user/contexts/user/domain/models"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/domain/models/models_mothers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	t.Run("should create user with valid data", func(t *testing.T) {
		// Arrange
		mother := models_mothers.NewUserMother()

		// Act
		user, err := mother.Build()

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, mother.ID, user.ID())
		assert.Equal(t, mother.Name, user.Name())
		assert.Equal(t, mother.LastName, user.LastName())
		assert.Equal(t, mother.Email, user.Email())
		assert.Equal(t, mother.Password, user.Password())
		assert.Equal(t, mother.Role, user.Role())
	})

	t.Run("should return error when password is too short", func(t *testing.T) {
		// Arrange & Act
		user, err := models_mothers.NewUserMother().WithPassword("Pass1!").Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, models.ErrPassTooShort, err)
	})

	t.Run("should return error when password has no uppercase", func(t *testing.T) {
		// Arrange & Act
		user, err := models_mothers.NewUserMother().WithPassword("password123!").Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, models.ErrPassNoUppercase, err)
	})

	t.Run("should return error when password has no lowercase", func(t *testing.T) {
		// Arrange & Act
		user, err := models_mothers.NewUserMother().WithPassword("PASSWORD123!").Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, models.ErrPassNoLowercase, err)
	})

	t.Run("should return error when password has no number", func(t *testing.T) {
		// Arrange & Act
		user, err := models_mothers.NewUserMother().WithPassword("Password!").Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, models.ErrPassNoNumber, err)
	})

	t.Run("should return error when password has no special character", func(t *testing.T) {
		// Arrange & Act
		user, err := models_mothers.NewUserMother().WithPassword("Password123").Build()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, models.ErrPassNoSpecialChar, err)
	})

	t.Run("should return error when role is invalid", func(t *testing.T) {
		// Arrange
		mother := models_mothers.NewUserMother()
		invalidRole := models.Role(999)
		hashedPassword := "hashed_password_123"

		// Act
		user, err := models.NewUser(mother.ID, mother.Name, mother.LastName, mother.Email, mother.Password, hashedPassword, invalidRole)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, models.ErrInvalidRole, err)
	})

	t.Run("should create user with admin role", func(t *testing.T) {
		// Arrange & Act
		user, err := models_mothers.NewUserMother().WithRole(models.RoleAdmin).Build()

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, models.RoleAdmin, user.Role())
	})

	t.Run("should create user with guest role", func(t *testing.T) {
		// Arrange & Act
		user, err := models_mothers.NewUserMother().WithRole(models.RoleGuest).Build()

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, models.RoleGuest, user.Role())
	})
}

func TestUser_Getters(t *testing.T) {
	t.Run("should return correct values from getters", func(t *testing.T) {
		// Arrange
		id := uuid.New()
		name := "John"
		lastName := "Doe"
		email := "john.doe@example.com"
		password := "MySecurePass123!"
		role := models.RoleAdmin

		user := models_mothers.NewUserMother().
			WithID(id).
			WithName(name).
			WithLastName(lastName).
			WithEmail(email).
			WithPassword(password).
			WithRole(role).
			MustBuild()

		// Act & Assert
		assert.Equal(t, id, user.ID())
		assert.Equal(t, name, user.Name())
		assert.Equal(t, lastName, user.LastName())
		assert.Equal(t, email, user.Email())
		assert.Equal(t, password, user.Password())
		assert.Equal(t, role, user.Role())
	})
}

func TestUser_PasswordValidation(t *testing.T) {
	testCases := []struct {
		name        string
		password    string
		expectedErr error
	}{
		{
			name:        "valid password with all requirements",
			password:    "MySecurePass123!",
			expectedErr: nil,
		},
		{
			name:        "valid password with symbols",
			password:    "TestPass1@#$",
			expectedErr: nil,
		},
		{
			name:        "valid password with punctuation",
			password:    "MyPass123.",
			expectedErr: nil,
		},
		{
			name:        "password too short",
			password:    "Pass1!",
			expectedErr: models.ErrPassTooShort,
		},
		{
			name:        "password with only 7 characters",
			password:    "Pass12!",
			expectedErr: models.ErrPassTooShort,
		},
		{
			name:        "password without uppercase",
			password:    "mypassword123!",
			expectedErr: models.ErrPassNoUppercase,
		},
		{
			name:        "password without lowercase",
			password:    "MYPASSWORD123!",
			expectedErr: models.ErrPassNoLowercase,
		},
		{
			name:        "password without number",
			password:    "MyPassword!",
			expectedErr: models.ErrPassNoNumber,
		},
		{
			name:        "password without special character",
			password:    "MyPassword123",
			expectedErr: models.ErrPassNoSpecialChar,
		},
		{
			name:        "password with only letters",
			password:    "MyPasswordOnly",
			expectedErr: models.ErrPassNoNumber,
		},
		{
			name:        "password with only numbers and letters",
			password:    "MyPassword123",
			expectedErr: models.ErrPassNoSpecialChar,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange & Act
			user, err := models_mothers.NewUserMother().WithPassword(tc.password).Build()

			// Assert
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Nil(t, user)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tc.password, user.Password())
			}
		})
	}
}

func TestUser_RoleValidation(t *testing.T) {
	testCases := []struct {
		name        string
		role        models.Role
		shouldError bool
	}{
		{
			name:        "valid admin role",
			role:        models.RoleAdmin,
			shouldError: false,
		},
		{
			name:        "valid user role",
			role:        models.RoleUser,
			shouldError: false,
		},
		{
			name:        "valid guest role",
			role:        models.RoleGuest,
			shouldError: false,
		},
		{
			name:        "invalid role - negative",
			role:        models.Role(-1),
			shouldError: true,
		},
		{
			name:        "invalid role - too high",
			role:        models.Role(999),
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange & Act
			user, err := models_mothers.NewUserMother().WithRole(tc.role).Build()

			// Assert
			if tc.shouldError {
				assert.Error(t, err)
				assert.Nil(t, user)
				assert.Equal(t, models.ErrInvalidRole, err)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tc.role, user.Role())
			}
		})
	}
}

func TestRole_String(t *testing.T) {
	testCases := []struct {
		role     models.Role
		expected string
	}{
		{models.RoleAdmin, "admin"},
		{models.RoleUser, "user"},
		{models.RoleGuest, "guest"},
		{models.Role(999), "unknown"},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			// Act & Assert
			assert.Equal(t, tc.expected, tc.role.String())
		})
	}
}

func TestRole_IsValid(t *testing.T) {
	testCases := []struct {
		role     models.Role
		expected bool
	}{
		{models.RoleAdmin, true},
		{models.RoleUser, true},
		{models.RoleGuest, true},
		{models.Role(-1), false},
		{models.Role(999), false},
		{models.Role(3), false},
	}

	for _, tc := range testCases {
		t.Run(tc.role.String(), func(t *testing.T) {
			// Act & Assert
			assert.Equal(t, tc.expected, tc.role.IsValid())
		})
	}
}
