package users

import (
	"gym-api/internal/errors"
	"net/http"
)

type Service interface {
	ListUsers() ([]User, error)
	CreateUser(input CreateUserRequest) (User, error)
	GetUserById(id string) (User, error)
}

type service struct {
	repo Repository
}

func NewUserService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) ListUsers() ([]User, error) {
	return s.repo.GetAll()
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
