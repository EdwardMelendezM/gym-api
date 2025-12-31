package impl

import (
	"context"
	"gym-api/internal/modules/users/models"
	"gym-api/internal/modules/users/repository"

	"gym-api/internal/utils/pagination"
)

type Service interface {
	ListUsers(ctx context.Context, p pagination.Params) (*pagination.Result[models.UserResponse], error)
	CreateUser(input models.CreateUserRequest) (models.User, error)
	GetUserById(id string) (models.User, error)
}

type service struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) Service {
	return &service{repo: repo}
}
