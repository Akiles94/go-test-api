package use_cases_mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockPasswordHasher struct {
	mock.Mock
}

func (m *MockPasswordHasher) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordHasher) ValidatePassword(password, hash string) bool {
	args := m.Called(password, hash)
	return args.Bool(0)
}

func NewMockPasswordHasher() *MockPasswordHasher {
	return &MockPasswordHasher{}
}

// HashPassword utils
func (m *MockPasswordHasher) SetupHashPasswordSuccess(password, hash string) *mock.Call {
	return m.On("HashPassword", password).Return(hash, nil)
}

func (m *MockPasswordHasher) SetupHashPasswordError(password string, err error) *mock.Call {
	return m.On("HashPassword", password).Return("", err)
}

// ValidatePassword utils
func (m *MockPasswordHasher) SetupValidatePasswordTrue(password, hash string) *mock.Call {
	return m.On("ValidatePassword", password, hash).Return(true)
}

func (m *MockPasswordHasher) SetupValidatePasswordFalse(password, hash string) *mock.Call {
	return m.On("ValidatePassword", password, hash).Return(false)
}
