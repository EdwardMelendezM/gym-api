package users

import (
	"gym-api/internal/errors"
	"gym-api/internal/modules/shared/pagination"
	"math"
	"net/http"
)

type Service interface {
	ListUsers(p pagination.Params) (*pagination.Result[UserResponse], error)
	CreateUser(input CreateUserRequest) (User, error)
	GetUserById(id string) (User, error)
}

type service struct {
	repo Repository
}

func NewUserService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) ListUsers(p pagination.Params) (*pagination.Result[UserResponse], error) {
	users, total, err := s.repo.GetAll(p)
	if err != nil {
		return nil, err
	}

	out := make([]UserResponse, 0, len(users))
	for _, u := range users {
		out = append(out, UserResponse{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(p.PageSize)))

	return &pagination.Result[UserResponse]{
		Data: out,
		Meta: pagination.Meta{
			Page:       p.Page,
			PageSize:   p.PageSize,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *service) CreateUser(input CreateUserRequest) (User, error) {
	user := User{
		FirstName: &input.FirstName,
		LastName:  &input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}

	created, err := s.repo.Create(user)
	if err != nil {
		return User{}, errors.New().
			SetStatus(http.StatusBadRequest).
			SetLayer("users.service").
			SetFunction("CreateUser").
			SetMessage(err.Error())
	}

	return created, nil
}

func (s *service) GetUserById(id string) (User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return User{}, errors.New().
			SetStatus(http.StatusNotFound).
			SetLayer("users.service").
			SetFunction("GetUserById").
			SetMessage("user not found")
	}

	return user, nil
}
