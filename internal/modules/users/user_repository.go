package users

import (
	"gym-api/internal/ent"
	"gym-api/internal/modules/shared/pagination"
)

type Repository interface {
	GetAll(p pagination.Params) ([]*ent.User, int, error)
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id string) (User, error)
}
