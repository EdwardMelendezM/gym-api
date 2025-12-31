package impl

import (
	"context"
	"gym-api/internal/modules/sessions/models"
	"gym-api/internal/utils/errors"
	"net/http"
)

func (r *sessionEntRepository) FindByID(id string) (*models.Session, error) {
	row, err := r.client.Session.
		Get(context.Background(), id)

	if err != nil {
		return nil, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("sessions.repository").
			SetFunction("FindByID").
			SetMessage("failed to find by id session").
			SetError(err)
	}

	return &models.Session{
		ID:        row.ID,
		UserID:    row.UserID,
		ExpiresAt: row.ExpiresAt,
		CreatedAt: row.CreatedAt,
	}, nil
}
