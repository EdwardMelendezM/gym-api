package impl

import (
	"gym-api/internal/ent"
	"gym-api/internal/modules/sessions/repository"
)

type sessionEntRepository struct {
	client *ent.Client
}

func NewSessionEntRepository(client *ent.Client) repository.SessionRepository {
	return &sessionEntRepository{client: client}
}
