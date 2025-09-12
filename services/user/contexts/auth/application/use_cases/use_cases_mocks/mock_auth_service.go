package use_cases_mocks

import (
	"context"

	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) GenerateToken(ctx context.Context, userID uuid.UUID, email, name *string, role string) (*value_objects.JWTToken, error) {
	args := m.Called(ctx, userID, email, name, role)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*value_objects.JWTToken), args.Error(1)
}

func (m *MockAuthService) ValidateToken(ctx context.Context, token string) (*value_objects.UserClaims, error) {
	args := m.Called(ctx, token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*value_objects.UserClaims), args.Error(1)
}

func (m *MockAuthService) RefreshToken(ctx context.Context, refreshToken string) (*value_objects.JWTToken, error) {
	args := m.Called(ctx, refreshToken)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*value_objects.JWTToken), args.Error(1)
}

func (m *MockAuthService) GenerateRefreshToken(ctx context.Context, userID uuid.UUID, email, role string) (*value_objects.JWTToken, error) {
	args := m.Called(ctx, userID, email, role)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*value_objects.JWTToken), args.Error(1)
}

func NewMockAuthService() *MockAuthService {
	return &MockAuthService{}
}

// GenerateToken utils
func (m *MockAuthService) SetupGenerateTokenSuccess(userID uuid.UUID, email, name *string, token string) *mock.Call {
	jwtToken := value_objects.NewJWTToken(token, 3600) // 1 hour expiration
	return m.On("GenerateToken", mock.Anything, userID, email, name).Return(jwtToken, nil)
}

func (m *MockAuthService) SetupGenerateTokenError(userID uuid.UUID, email, name *string, err error) *mock.Call {
	return m.On("GenerateToken", mock.Anything, userID, email, name).Return(nil, err)
}

// ValidateToken utils
func (m *MockAuthService) SetupValidateTokenSuccess(token string, claims *value_objects.UserClaims) *mock.Call {
	return m.On("ValidateToken", mock.Anything, token).Return(claims, nil)
}

func (m *MockAuthService) SetupValidateTokenError(token string, err error) *mock.Call {
	return m.On("ValidateToken", mock.Anything, token).Return(nil, err)
}

// RefreshToken utils
func (m *MockAuthService) SetupRefreshTokenSuccess(refreshToken string, newToken string) *mock.Call {
	jwtToken := value_objects.NewJWTToken(newToken, 3600)
	return m.On("RefreshToken", mock.Anything, refreshToken).Return(jwtToken, nil)
}

func (m *MockAuthService) SetupRefreshTokenError(refreshToken string, err error) *mock.Call {
	return m.On("RefreshToken", mock.Anything, refreshToken).Return(nil, err)
}

// GenerateRefreshToken utils
func (m *MockAuthService) SetupGenerateRefreshTokenSuccess(userID uuid.UUID, email string, refreshToken string) *mock.Call {
	jwtToken := value_objects.NewJWTToken(refreshToken, 604800) // 7 days expiration
	return m.On("GenerateRefreshToken", mock.Anything, userID, email).Return(jwtToken, nil)
}

func (m *MockAuthService) SetupGenerateRefreshTokenError(userID uuid.UUID, email string, err error) *mock.Call {
	return m.On("GenerateRefreshToken", mock.Anything, userID, email).Return(nil, err)
}
