package sessions

import (
	"context"

	"gym-api/internal/ent"
)

type sessionEntRepository struct {
	client *ent.Client
}

func NewSessionEntRepository(client *ent.Client) SessionRepository {
	return &sessionEntRepository{client: client}
}

func (r *sessionEntRepository) Create(s Session) (Session, error) {
	row, err := r.client.Session.
		Create().
		SetID(s.ID).
		SetUserID(s.UserID).
		SetExpiresAt(s.ExpiresAt).
		Save(context.Background())

	if err != nil {
		return Session{}, err
	}

	return Session{
		ID:        row.ID,
		UserID:    row.UserID,
		ExpiresAt: row.ExpiresAt,
		CreatedAt: row.CreatedAt,
	}, nil
}

func (r *sessionEntRepository) FindByID(id string) (*Session, error) {
	row, err := r.client.Session.
		Get(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return &Session{
		ID:        row.ID,
		UserID:    row.UserID,
		ExpiresAt: row.ExpiresAt,
		CreatedAt: row.CreatedAt,
	}, nil
}

func (r *sessionEntRepository) Delete(id string) error {
	return r.client.Session.
		DeleteOneID(id).
		Exec(context.Background())
}
