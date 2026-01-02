package impl

import (
	"context"
	errorMessage "gym-api/internal/modules/sessions/errors"
	"gym-api/internal/modules/sessions/models"
	"gym-api/internal/utils/errors"
)

func (r *sessionEntRepository) FindSessionByID(id string) (*models.Session, error) {
	row, err := r.client.Session.
		Get(context.Background(), id)

	if err != nil {
		return nil, errors.WithContext(
			errorMessage.ErrFindSessionById,
			"sessions.repository",
			"FindSessionByID",
		)
	}

	return &models.Session{
		ID:        row.ID,
		UserID:    row.UserID,
		ExpiresAt: row.ExpiresAt,
		CreatedAt: row.CreatedAt,
	}, nil
}
