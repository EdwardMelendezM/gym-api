package impl

import (
	models2 "gym-api/internal/modules/sessions/models"
	"time"

	"github.com/google/uuid"

	errorMessage "gym-api/internal/modules/auth/errors"
	"gym-api/internal/modules/auth/models"
	"gym-api/internal/utils/auth"
	"gym-api/internal/utils/errors"
)

func (s AuthService) Login(input models.LoginRequest) (models.TokenResponse, error) {
	user, err := s.users.FindUserByEmail(input.Email)
	if err != nil {
		return models.TokenResponse{}, errors.WithContext(
			errorMessage.ErrInvalidCredentials,
			"auth.service",
			"Login",
		)
	}

	if !auth.ComparePassword(user.Password, input.Password) {
		return models.TokenResponse{}, errors.WithContext(
			errorMessage.ErrInvalidCredentials,
			"auth.service",
			"Login",
		)
	}

	session := models2.Session{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	created, err := s.sessions.Create(session)
	if err != nil {
		return models.TokenResponse{}, errors.WrapWithContext(
			err,
			"ERROR_SESSION_CREATION",
			"failed to create session",
			"auth.service",
			"Login",
		)
	}

	token, err := auth.GenerateToken(created.ID, time.Hour)
	if err != nil {
		return models.TokenResponse{}, errors.WrapWithContext(
			err,
			"ERROR_GENERATE_TOKEN",
			"failed to generate access token",
			"auth.service",
			"Login",
		)
	}

	refresh, err := auth.GenerateToken(created.ID, 24*time.Hour)
	if err != nil {
		return models.TokenResponse{}, errors.WrapWithContext(
			err,
			"ERROR_GENERATE_TOKEN",
			"failed to generate refresh token",
			"auth.service",
			"Login",
		)
	}

	return models.TokenResponse{
		AccessToken:  token,
		RefreshToken: refresh,
	}, nil
}
