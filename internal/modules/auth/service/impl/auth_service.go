package impl

import (
	"gym-api/internal/modules/auth/models"
	"gym-api/internal/modules/sessions"
	"gym-api/internal/modules/users"
)

type Service interface {
	Register(input models.RegisterRequest) (models.TokenResponse, error)
	Login(input models.LoginRequest) (models.TokenResponse, error)
	RefreshToken(input models.RefreshTokenRequest) (models.TokenResponse, error)
}

type AuthService struct {
	users    users.Repository
	sessions sessions.SessionRepository
}

func NewAuthService(
	usersRepo users.Repository,
	sessionsRepo sessions.SessionRepository,
) Service {
	return &AuthService{
		users:    usersRepo,
		sessions: sessionsRepo,
	}
}
