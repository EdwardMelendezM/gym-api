package impl

import (
	"context"
	"net/http"

	"gym-api/internal/ent/user"
	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/errors"
)

func (r *entRepository) FindUserByEmail(email string) (models.User, error) {
	row, err := r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(context.Background())

	if err != nil {
		return models.User{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("users.repository").
			SetFunction("FindUserByEmail").
			SetMessage("failed to find by email").
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
