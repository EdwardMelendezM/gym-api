package impl

import (
	"context"
	"net/http"

	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/errors"
)

func (r *entRepository) CreateUser(user models.User) (models.User, error) {
	row, err := r.client.User.
		Create().
		SetNillableFirstName(user.FirstName).
		SetNillableLastName(user.LastName).
		SetNillableFullName(user.FullName).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(context.Background())

	if err != nil {
		return models.User{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("users.repository").
			SetFunction("CreateUser").
			SetMessage("failed to insert user").
			SetError(err)
	}

	return models.User{
		ID:        row.ID,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		FullName:  row.FullName,
		Email:     row.Email,
		Password:  row.Password,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}, nil
}
