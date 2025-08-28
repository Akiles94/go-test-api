package use_cases_mocks

import (
	"context"

	"github.com/Akiles94/go-test-api/services/user/contexts/user/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(models.User), args.Error(1)
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

// CreateUser utils
func (m *MockUserRepository) SetupCreateUserSuccess() *mock.Call {
	return m.On("CreateUser", mock.Anything, mock.AnythingOfType("*models.user")).Return(nil)
}

func (m *MockUserRepository) SetupCreateUserError(err error) *mock.Call {
	return m.On("CreateUser", mock.Anything, mock.AnythingOfType("*models.user")).Return(err)
}

// GetUserByEmail utils
func (m *MockUserRepository) SetupGetUserByEmailSuccess(email string, user models.User) *mock.Call {
	return m.On("GetUserByEmail", mock.Anything, email).Return(user, nil)
}

func (m *MockUserRepository) SetupGetUserByEmailNotFound(email string, err error) *mock.Call {
	return m.On("GetUserByEmail", mock.Anything, email).Return(nil, err)
}

func (m *MockUserRepository) SetupGetUserByEmailError(email string, err error) *mock.Call {
	return m.On("GetUserByEmail", mock.Anything, email).Return(nil, err)
}
