package users

import "gym-api/internal/errors"

type Service interface {
	ListUsers() ([]User, error)
	CreateUser(input CreateUserRequest) (User, error)
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
		return User{}, errors.Internal("could not create user")
	}

	return created, nil
}
