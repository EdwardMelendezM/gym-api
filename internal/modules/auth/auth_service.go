package auth

import (
	"errors"
	"gym-api/internal/modules/sessions"
	utils2 "gym-api/internal/modules/shared/utils"
	"gym-api/internal/modules/users"
	"time"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(input RegisterRequest) (string, error)
	Login(input LoginRequest) (string, error)
}

type service struct {
	users    users.Repository
	sessions sessions.SessionRepository
}

func NewAuthService(
	usersRepo users.Repository,
	sessionsRepo sessions.SessionRepository,
) AuthService {
	return &service{
		users:    usersRepo,
		sessions: sessionsRepo,
	}
}

func (s *service) Register(input RegisterRequest) (string, error) {
	hash, err := utils2.HashPassword(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.users.Create(users.User{
		FirstName: &input.FirstName,
		LastName:  &input.LastName,
		Email:     input.Email,
		Password:  hash,
	})
	if err != nil {
		return "", err
	}

	session := sessions.Session{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	created, err := s.sessions.Create(session)
	if err != nil {
		return "", err
	}

	return utils2.GenerateToken(created.ID)
}

func (s *service) Login(input LoginRequest) (string, error) {
	user, err := s.users.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils2.ComparePassword(user.Password, input.Password) {
		return "", errors.New("invalid credentials")
	}

	return utils2.GenerateToken(user.ID)
}
