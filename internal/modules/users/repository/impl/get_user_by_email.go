package impl

import (
	"context"
	"gym-api/internal/ent/user"
	authError "gym-api/internal/modules/auth/errors"
	"gym-api/internal/modules/users/models"
	"gym-api/internal/utils/errors"
)

func (r *entRepository) FindUserByEmail(email string) (models.User, error) {
	row, err := r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(context.Background())

	if err != nil {
		return models.User{}, errors.Internal(string(errors.CodeInternalError), "error finding user by email").
			SetLayer(string(errors.LayerRepository)).
			SetFunction("FindUserByEmail").
			SetError(err)
	}

	if row == nil {
		return models.User{}, authError.ErrorUserNotFound.
			SetLayer(string(errors.LayerRepository)).
			SetFunction("FindUserByEmail")
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
