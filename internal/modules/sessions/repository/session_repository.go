package repository

import "gym-api/internal/modules/sessions/models"

type SessionRepository interface {
	Create(session models.Session) (models.Session, error)
	FindByID(id string) (*models.Session, error)
	Delete(id string) error
}
