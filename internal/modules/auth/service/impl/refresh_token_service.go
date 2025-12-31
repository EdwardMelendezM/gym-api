package impl

import (
	"net/http"
	"time"

	"gym-api/internal/modules/auth/models"
	"gym-api/internal/utils/auth"
	"gym-api/internal/utils/errors"
)

func (s AuthService) RefreshToken(input models.RefreshTokenRequest) (models.TokenResponse, error) {
	result, err := auth.ParseToken(input.RefreshToken)
	if err != nil {
		return models.TokenResponse{}, errors.New().
			SetStatus(http.StatusUnauthorized).
			SetLayer("auth.service").
			SetFunction("RefreshToken").
			SetMessage("invalid refresh token").
			SetError(err)
	}

	session, err := s.sessions.FindByID(result.SessionID)
	if err != nil {
		return models.TokenResponse{}, errors.New().
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

	token, errToken := auth.GenerateToken(session.ID, expiresAtToken)
	if errToken != nil {
		return models.TokenResponse{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("RefreshToken").
			SetMessage("failed to generate access token").
			SetError(err)
	}

	tokenRefresh, errTokenRefresh := auth.GenerateToken(session.ID, expiresAtRefresh)
	if errTokenRefresh != nil {
		return models.TokenResponse{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("RefreshToken").
			SetMessage("failed to generate refresh token").
			SetError(err)
	}

	tokenResponse := models.TokenResponse{
		token,
		tokenRefresh,
	}
	return tokenResponse, nil
}
