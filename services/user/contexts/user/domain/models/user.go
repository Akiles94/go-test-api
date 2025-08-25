package models

import (
	"unicode"

	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
	"github.com/google/uuid"
)

var (
	ErrPassTooShort = value_objects.DomainError{
		Code:    "USER_PASSWORD_TOO_SHORT",
		Message: "User password must be at least 8 characters long",
	}
	ErrPassNoUppercase = value_objects.DomainError{
		Code:    "USER_PASSWORD_NO_UPPERCASE",
		Message: "User password must contain at least one uppercase letter",
	}
	ErrPassNoLowercase = value_objects.DomainError{
		Code:    "USER_PASSWORD_NO_LOWERCASE",
		Message: "User password must contain at least one lowercase letter",
	}
	ErrPassNoNumber = value_objects.DomainError{
		Code:    "USER_PASSWORD_NO_NUMBER",
		Message: "User password must contain at least one number",
	}
	ErrPassNoSpecialChar = value_objects.DomainError{
		Code:    "USER_PASSWORD_NO_SPECIAL_CHAR",
		Message: "User password must contain at least one special character",
	}
	ErrInvalidRole = value_objects.DomainError{
		Code:    "USER_INVALID_ROLE",
		Message: "Invalid user role",
	}
)

type User interface {
	ID() uuid.UUID
	Name() string
	LastName() string
	Email() string
	Password() string
	PasswordHash() string
	Role() Role
}

type user struct {
	id           uuid.UUID
	name         string
	lastName     string
	email        string
	password     string
	passwordHash string
	role         Role
}

func NewUser(id uuid.UUID, name, lastName, email, password, passwordHash string, role Role) (User, error) {
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	if !role.IsValid() {
		return nil, ErrInvalidRole
	}

	return &user{
		id:           id,
		name:         name,
		lastName:     lastName,
		email:        email,
		password:     password,
		passwordHash: passwordHash,
		role:         role,
	}, nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return ErrPassTooShort
	}

	var (
		hasUppercase   = false
		hasLowercase   = false
		hasNumber      = false
		hasSpecialChar = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}

	if !hasUppercase {
		return ErrPassNoUppercase
	}

	if !hasLowercase {
		return ErrPassNoLowercase
	}

	if !hasNumber {
		return ErrPassNoNumber
	}

	if !hasSpecialChar {
		return ErrPassNoSpecialChar
	}

	return nil
}

func (u *user) ID() uuid.UUID {
	return u.id
}

func (u *user) Name() string {
	return u.name
}

func (u *user) LastName() string {
	return u.lastName
}

func (u *user) Email() string {
	return u.email
}

func (u *user) Password() string {
	return u.password
}

func (u *user) PasswordHash() string {
	return u.passwordHash
}

func (u *user) Role() Role {
	return u.role
}
