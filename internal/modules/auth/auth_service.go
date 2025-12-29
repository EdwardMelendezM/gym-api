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
	Register(input RegisterRequest) (TokenResponse, error)
	Login(input LoginRequest) (TokenResponse, error)
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

func (s *service) Register(input RegisterRequest) (TokenResponse, error) {
	hash, err := utils2.HashPassword(input.Password)
	if err != nil {
		return TokenResponse{}, err
	}

	user, err := s.users.Create(users.User{
		FirstName: &input.FirstName,
		LastName:  &input.LastName,
		Email:     input.Email,
		Password:  hash,
	})
	if err != nil {
		return TokenResponse{}, err
	}

	session := sessions.Session{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	created, err := s.sessions.Create(session)
	if err != nil {
		return TokenResponse{}, err
	}

	// One day expiration
	expiresAtToken := 1 * time.Hour
	// One day expiration
	expiresAtRefresh := 24 * time.Hour

	token, errToken := utils2.GenerateToken(created.ID, expiresAtToken)
	if errToken != nil {
		return TokenResponse{}, errToken
	}

	tokenRefresh, errTokenRefresh := utils2.GenerateToken(created.ID, expiresAtRefresh)
	if errTokenRefresh != nil {
		return TokenResponse{}, errTokenRefresh
	}

	tokenResponse := TokenResponse{
		token,
		tokenRefresh,
	}
	return tokenResponse, nil
}

func (s *service) Login(input LoginRequest) (TokenResponse, error) {
	user, err := s.users.FindByEmail(input.Email)
	if err != nil {
		return TokenResponse{}, errors.New("invalid credentials")
	}

	if !utils2.ComparePassword(user.Password, input.Password) {
		return TokenResponse{}, errors.New("invalid credentials")
	}

	// One hour expiration
	expiresAtToken := 1 * time.Hour
	// One day expiration
	expiresAtRefresh := 24 * time.Hour

	token, errToken := utils2.GenerateToken(user.ID, expiresAtToken)
	if errToken != nil {
		return TokenResponse{}, errToken
	}

	tokenRefresh, errTokenRefresh := utils2.GenerateToken(user.ID, expiresAtRefresh)
	if errTokenRefresh != nil {
		return TokenResponse{}, errTokenRefresh
	}

	tokenResponse := TokenResponse{
		token,
		tokenRefresh,
	}
	return tokenResponse, nil
}
