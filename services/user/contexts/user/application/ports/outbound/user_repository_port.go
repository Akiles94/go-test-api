package outbound

import (
	"context"

	"github.com/Akiles94/go-test-api/services/user/contexts/user/domain/models"
)

type UserRepositoryPort interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}
