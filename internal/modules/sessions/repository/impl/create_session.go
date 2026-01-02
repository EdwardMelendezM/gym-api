package impl

import (
	"context"
	errorMessage "gym-api/internal/modules/sessions/errors"
	"gym-api/internal/modules/sessions/models"
	"gym-api/internal/utils/errors"
)

func (r *sessionEntRepository) CreateSession(s models.Session) (models.Session, error) {
	row, err := r.client.Session.
		Create().
		SetID(s.ID).
		SetUserID(s.UserID).
		SetExpiresAt(s.ExpiresAt).
		Save(context.Background())

	if err != nil {
		return models.Session{}, errors.WithContext(
			errorMessage.ErrCreateSession,
			"sessions.repository",
			"CreateUser",
		)
	}

	return models.Session{
		ID:        row.ID,
		UserID:    row.UserID,
		ExpiresAt: row.ExpiresAt,
		CreatedAt: row.CreatedAt,
	}, nil
}
