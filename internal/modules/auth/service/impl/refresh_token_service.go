package impl

import (
	"net/http"
	"time"

	"gym-api/internal/modules/auth/models"
	"gym-api/internal/modules/shared/errors"
	"gym-api/internal/modules/shared/utils"
)

func (s AuthService) RefreshToken(input models.RefreshTokenRequest) (models.TokenResponse, error) {
	result, err := utils.ParseToken(input.RefreshToken)
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

	token, errToken := utils.GenerateToken(session.ID, expiresAtToken)
	if errToken != nil {
		return models.TokenResponse{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("RefreshToken").
			SetMessage("failed to generate access token").
			SetError(err)
	}

	tokenRefresh, errTokenRefresh := utils.GenerateToken(session.ID, expiresAtRefresh)
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
