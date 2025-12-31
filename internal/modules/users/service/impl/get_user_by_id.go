package impl

import (
	"net/http"

	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/errors"
)

func (s *service) GetUserById(id string) (models.User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return models.User{}, errors.New().
			SetStatus(http.StatusNotFound).
			SetLayer("users.service").
			SetFunction("GetUserById").
			SetMessage("user not found")
	}

	return user, nil
}
