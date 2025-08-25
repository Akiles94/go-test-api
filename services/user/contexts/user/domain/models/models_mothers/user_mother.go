package models_mothers

import (
	"github.com/Akiles94/go-test-api/services/user/contexts/user/domain/models"
	"github.com/google/uuid"
)

type UserMother struct {
	ID           uuid.UUID
	Name         string
	LastName     string
	Email        string
	Password     string
	PasswordHash string
	Role         models.Role
}

func NewUserMother() *UserMother {
	return &UserMother{
		ID:           uuid.New(),
		Name:         "John",
		LastName:     "Doe",
		Email:        "john.doe@example.com",
		Password:     "Password123!",
		PasswordHash: "$2a$10$EIXZQ1z5Q5Q5Q5Q5Q5Q5QO",
		Role:         models.RoleUser,
	}
}

func (um *UserMother) WithID(id uuid.UUID) *UserMother {
	um.ID = id
	return um
}

func (um *UserMother) WithName(name string) *UserMother {
	um.Name = name
	return um
}

func (um *UserMother) WithLastName(lastName string) *UserMother {
	um.LastName = lastName
	return um
}

func (um *UserMother) WithEmail(email string) *UserMother {
	um.Email = email
	return um
}

func (um *UserMother) WithPassword(password string) *UserMother {
	um.Password = password
	return um
}

func (um *UserMother) WithPasswordHash(passwordHash string) *UserMother {
	um.PasswordHash = passwordHash
	return um
}

func (um *UserMother) WithRole(role models.Role) *UserMother {
	um.Role = role
	return um
}

func (um *UserMother) Build() (models.User, error) {
	return models.NewUser(
		um.ID,
		um.Name,
		um.LastName,
		um.Email,
		um.Password,
		um.PasswordHash,
		um.Role,
	)
}

func (um *UserMother) MustBuild() models.User {
	user, err := um.Build()
	if err != nil {
		panic(err)
	}
	return user
}
