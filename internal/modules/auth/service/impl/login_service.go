package impl

import (
	"time"

	"github.com/google/uuid"

	authError "gym-api/internal/modules/auth/errors"
	"gym-api/internal/modules/auth/models"
	models2 "gym-api/internal/modules/sessions/models"
	"gym-api/internal/utils/auth"
	"gym-api/internal/utils/errors"
)

func (s AuthService) Login(input models.LoginRequest) (models.TokenResponse, error) {
	user, err := s.users.FindUserByEmail(input.Email)
	if err != nil {
		return models.TokenResponse{}, authError.ErrorInvalidCredentials.
			SetLayer(string(errors.LayerService)).
			SetFunction("Login").
			SetError(err)
	}

	if !auth.ComparePassword(user.Password, input.Password) {
		return models.TokenResponse{}, authError.ErrorInvalidCredentials.
			SetLayer(string(errors.LayerService)).
			SetFunction("Login")
	}

	session := models2.Session{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	created, err := s.sessions.CreateSession(session)
	if err != nil {
		return models.TokenResponse{}, authError.ErrorInvalidCredentials.
			SetLayer(string(errors.LayerService)).
			SetFunction("Login").
			SetError(err)
	}

	token, err := auth.GenerateToken(created.ID, time.Hour)
	if err != nil {
		return models.TokenResponse{}, authError.ErrorGenerateToken.
			SetLayer(string(errors.LayerService)).
			SetFunction("Login").
			SetError(err)
	}

	refresh, err := auth.GenerateToken(created.ID, 24*time.Hour)
	if err != nil {
		return models.TokenResponse{}, authError.ErrorGenerateToken.
			SetLayer(string(errors.LayerService)).
			SetFunction("Login").
			SetError(err)
	}

	return models.TokenResponse{
		AccessToken:  token,
		RefreshToken: refresh,
	}, nil
}
