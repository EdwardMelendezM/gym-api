package impl

import (
	"gym-api/internal/modules/auth/models"
	"gym-api/internal/modules/sessions/repository"
	repository2 "gym-api/internal/modules/users/repository"
)

type Service interface {
	Register(input models.RegisterRequest) (models.TokenResponse, error)
	Login(input models.LoginRequest) (models.TokenResponse, error)
	RefreshToken(input models.RefreshTokenRequest) (models.TokenResponse, error)
}

type AuthService struct {
	users    repository2.Repository
	sessions repository.SessionRepository
}

func NewAuthService(
	usersRepo repository2.Repository,
	sessionsRepo repository.SessionRepository,
) Service {
	return &AuthService{
		users:    usersRepo,
		sessions: sessionsRepo,
	}
}
