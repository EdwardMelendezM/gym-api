package repository

import (
	"context"
	"gym-api/internal/ent"
	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/pagination"
)

type Repository interface {
	GetAll(ctx context.Context, p pagination.Params) ([]*ent.User, int, error)
	Create(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindById(id string) (models.User, error)
}
