package repository

import (
	"github.com/Akiles94/go-test-api/services/user/contexts/user/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name     string    `gorm:"type:varchar(100);not null"`
	LastName string    `gorm:"type:varchar(100);not null"`
	Email    string    `gorm:"type:varchar(100);not null;unique"`
	Password string    `gorm:"type:varchar(100);not null"`
	Role     int       `gorm:"type:int;not null"`
}

func (UserEntity) TableName() string {
	return "users"
}

func NewUserEntityFromDomain(user models.User) *UserEntity {
	return &UserEntity{
		ID:       user.ID(),
		Name:     user.Name(),
		LastName: user.LastName(),
		Email:    user.Email(),
		Password: user.Password(),
		Role:     int(user.Role()),
	}
}

func (ue *UserEntity) ToDomainModel() models.User {
	user, err := models.NewUser(
		ue.ID,
		ue.Name,
		ue.LastName,
		ue.Email,
		ue.Password,
		ue.Password,
		models.Role(ue.Role),
	)
	if err != nil {
		return nil
	}
	return user
}
