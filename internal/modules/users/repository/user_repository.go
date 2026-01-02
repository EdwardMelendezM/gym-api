package repository

import (
	"context"

	"gym-api/internal/ent"
	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/pagination"
)

type Repository interface {
	GetUsersPaginated(ctx context.Context, p pagination.Params) ([]*ent.User, error)
	GetTotalUsers(ctx context.Context) (int, error)
	CreateUser(user models.User) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	FindUserById(id string) (models.User, error)
}
