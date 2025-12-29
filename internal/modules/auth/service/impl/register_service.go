package impl

import (
	"net/http"
	"time"

	"gym-api/internal/modules/auth/models"
	"gym-api/internal/modules/sessions"
	"gym-api/internal/modules/shared/errors"
	"gym-api/internal/modules/shared/utils"
	"gym-api/internal/modules/users"

	"github.com/google/uuid"
)

func (s AuthService) Register(input models.RegisterRequest) (models.TokenResponse, error) {
	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return models.TokenResponse{}, errors.New().
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
		return models.TokenResponse{}, errors.New().
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
		return models.TokenResponse{}, errors.New().
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

	token, errToken := utils.GenerateToken(created.ID, expiresAtToken)
	if errToken != nil {
		return models.TokenResponse{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("Register").
			SetMessage("failed to generate access token").
			SetError(errToken)
	}

	tokenRefresh, errTokenRefresh := utils.GenerateToken(created.ID, expiresAtRefresh)
	if errTokenRefresh != nil {
		return models.TokenResponse{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("Register").
			SetMessage("failed to generate access token").
			SetError(errToken)
	}

	tokenResponse := models.TokenResponse{
		token,
		tokenRefresh,
	}
	return tokenResponse, nil
}
