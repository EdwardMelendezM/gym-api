package repository

import "gym-api/internal/modules/sessions/models"

type SessionRepository interface {
	CreateSession(session models.Session) (models.Session, error)
	FindSessionByID(id string) (*models.Session, error)
	DeleteSessionById(id string) error
}
