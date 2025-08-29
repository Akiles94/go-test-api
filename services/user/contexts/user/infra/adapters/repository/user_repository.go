package repository

import (
	"context"
	"errors"

	"github.com/Akiles94/go-test-api/services/user/contexts/user/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user models.User) error {
	userEntity := NewUserEntityFromDomain(user)
	return ur.db.WithContext(ctx).Create(&userEntity).Error
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var userEntity UserEntity
	if err := ur.db.WithContext(ctx).Where("email = ?", email).First(&userEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	user := userEntity.ToDomainModel()
	return &user, nil
}
