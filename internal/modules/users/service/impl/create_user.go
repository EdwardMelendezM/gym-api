package impl

import (
	"net/http"

	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/errors"
)

func (s *service) CreateUser(input models.CreateUserRequest) (models.User, error) {
	user := models.User{
		FirstName: &input.FirstName,
		LastName:  &input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}

	created, err := s.repo.CreateUser(user)
	if err != nil {
		return models.User{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("users.service").
			SetFunction("CreateUser").
			SetMessage(err.Error())
	}

	return created, nil
}
