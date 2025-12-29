package auth

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	appErrors "gym-api/internal/errors"
	"gym-api/internal/modules/sessions"
	utils2 "gym-api/internal/modules/shared/utils"
	"gym-api/internal/modules/users"
)

type AuthService interface {
	Register(input RegisterRequest) (TokenResponse, error)
	Login(input LoginRequest) (TokenResponse, error)
	RefreshToken(input RefreshTokenRequest) (TokenResponse, error)
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
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("Register").
			SetMessage("failed to hash password").
			SetError(err)
	}

	fullName := input.FirstName + " " + input.LastName

	user, err := s.users.Create(users.User{
		FirstName: &input.FirstName,
		FullName:  &fullName,
		LastName:  &input.LastName,
		Email:     input.Email,
		Password:  hash,
	})
	if err != nil {
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusConflict).
			SetLayer("auth.service").
			SetFunction("Register").
			SetMessage("email already registered").
			SetError(err)
	}

	// Create session
	session := sessions.Session{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	created, err := s.sessions.Create(session)
	if err != nil {
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("Register").
			SetMessage("failed to create session").
			SetError(err)
	}

	// One day expiration
	expiresAtToken := 1 * time.Hour
	// One day expiration
	expiresAtRefresh := 24 * time.Hour

	token, errToken := utils2.GenerateToken(created.ID, expiresAtToken)
	if errToken != nil {
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("Register").
			SetMessage("failed to generate access token").
			SetError(errToken)
	}

	tokenRefresh, errTokenRefresh := utils2.GenerateToken(created.ID, expiresAtRefresh)
	if errTokenRefresh != nil {
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("Register").
			SetMessage("failed to generate access token").
			SetError(errToken)
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
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusUnauthorized).
			SetLayer("auth.service").
			SetFunction("Login").
			SetMessage("invalid credentials")
	}

	if !utils2.ComparePassword(user.Password, input.Password) {
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusUnauthorized).
			SetLayer("auth.service").
			SetFunction("Login").
			SetMessage("invalid credentials")
	}

	// Create session
	session := sessions.Session{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	created, err := s.sessions.Create(session)
	if err != nil {
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("Register").
			SetMessage("failed to create session").
			SetError(err)
	}

	// One hour expiration
	expiresAtToken := 1 * time.Hour
	// One day expiration
	expiresAtRefresh := 24 * time.Hour

	token, errToken := utils2.GenerateToken(created.ID, expiresAtToken)
	if errToken != nil {
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("Login").
			SetMessage("failed to generate access token").
			SetError(err)
	}

	tokenRefresh, errTokenRefresh := utils2.GenerateToken(created.ID, expiresAtRefresh)
	if errTokenRefresh != nil {
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("Login").
			SetMessage("failed to generate refresh token").
			SetError(err)
	}

	tokenResponse := TokenResponse{
		token,
		tokenRefresh,
	}
	return tokenResponse, nil
}

func (s *service) RefreshToken(input RefreshTokenRequest) (TokenResponse, error) {
	result, err := utils2.ParseToken(input.RefreshToken)
	if err != nil {
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusUnauthorized).
			SetLayer("auth.service").
			SetFunction("RefreshToken").
			SetMessage("invalid refresh token").
			SetError(err)
	}

	session, err := s.sessions.FindByID(result.SessionID)
	if err != nil {
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusUnauthorized).
			SetLayer("auth.service").
			SetFunction("RefreshToken").
			SetMessage("invalid refresh token").
			SetError(err)
	}

	// One hour expiration
	expiresAtToken := 1 * time.Hour
	// One day expiration
	expiresAtRefresh := 24 * time.Hour

	token, errToken := utils2.GenerateToken(session.ID, expiresAtToken)
	if errToken != nil {
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("RefreshToken").
			SetMessage("failed to generate access token").
			SetError(err)
	}

	tokenRefresh, errTokenRefresh := utils2.GenerateToken(session.ID, expiresAtRefresh)
	if errTokenRefresh != nil {
		return TokenResponse{}, appErrors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("RefreshToken").
			SetMessage("failed to generate refresh token").
			SetError(err)
	}

	tokenResponse := TokenResponse{
		token,
		tokenRefresh,
	}
	return tokenResponse, nil
}
