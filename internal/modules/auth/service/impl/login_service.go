package impl

import (
	models2 "gym-api/internal/modules/sessions/models"
	"net/http"
	"time"

	"github.com/google/uuid"

	"gym-api/internal/modules/auth/models"
	auth "gym-api/internal/utils/auth"
	"gym-api/internal/utils/errors"
)

func (s AuthService) Login(input models.LoginRequest) (models.TokenResponse, error) {
	user, err := s.users.FindByEmail(input.Email)
	if err != nil {
		return models.TokenResponse{}, errors.New().
			SetStatus(http.StatusUnauthorized).
			SetLayer("auth.service").
			SetFunction("Login").
			SetMessage("invalid credentials")
	}

	if !auth.ComparePassword(user.Password, input.Password) {
		return models.TokenResponse{}, errors.New().
			SetStatus(http.StatusUnauthorized).
			SetLayer("auth.service").
			SetFunction("Login").
			SetMessage("invalid credentials")
	}

	// Create session
	session := models2.Session{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	created, err := s.sessions.Create(session)
	if err != nil {
		return models.TokenResponse{}, errors.New().
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

	token, errToken := auth.GenerateToken(created.ID, expiresAtToken)
	if errToken != nil {
		return models.TokenResponse{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("Login").
			SetMessage("failed to generate access token").
			SetError(err)
	}

	tokenRefresh, errTokenRefresh := auth.GenerateToken(created.ID, expiresAtRefresh)
	if errTokenRefresh != nil {
		return models.TokenResponse{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("Login").
			SetMessage("failed to generate refresh token").
			SetError(err)
	}

	tokenResponse := models.TokenResponse{
		token,
		tokenRefresh,
	}
	return tokenResponse, nil
}
