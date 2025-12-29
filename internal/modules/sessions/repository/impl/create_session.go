package impl

import (
	"context"
	"gym-api/internal/modules/sessions/models"
	"gym-api/internal/modules/shared/errors"
	"net/http"
)

func (r *sessionEntRepository) Create(s models.Session) (models.Session, error) {
	row, err := r.client.Session.
		Create().
		SetID(s.ID).
		SetUserID(s.UserID).
		SetExpiresAt(s.ExpiresAt).
		Save(context.Background())

	if err != nil {
		return models.Session{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("sessions.repository").
			SetFunction("Create").
			SetMessage("failed to create session").
			SetError(err)
	}

	return models.Session{
		ID:        row.ID,
		UserID:    row.UserID,
		ExpiresAt: row.ExpiresAt,
		CreatedAt: row.CreatedAt,
	}, nil
}
