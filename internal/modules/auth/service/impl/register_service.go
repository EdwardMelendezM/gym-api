package impl

import (
	"net/http"
	"time"

	authError "gym-api/internal/modules/auth/errors"
	"gym-api/internal/modules/auth/models"
	models2 "gym-api/internal/modules/sessions/models"
	models3 "gym-api/internal/modules/users/models"
	"gym-api/internal/utils/auth"
	"gym-api/internal/utils/errors"

	"github.com/google/uuid"
)

func (s AuthService) Register(input models.RegisterRequest) (models.TokenResponse, error) {
	hash, err := auth.HashPassword(input.Password)
	if err != nil {
		return models.TokenResponse{}, authError.ErrorGenerateHashPassword.
			SetLayer("auth.service").
			SetFunction("Register").
			SetError(err)

	}

	_, err = s.users.FindUserByEmail(input.Email)
	if err != nil {
		return models.TokenResponse{}, authError.ErrUserAlreadyExists.
			SetLayer("auth.service").
			SetFunction("Register").
			SetError(err)
	}

	fullName := input.FirstName + " " + input.LastName
	user, err := s.users.CreateUser(models3.User{
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

	// CreateUser session
	session := models2.Session{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	created, err := s.sessions.CreateSession(session)
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

	token, errToken := auth.GenerateToken(created.ID, expiresAtToken)
	if errToken != nil {
		return models.TokenResponse{}, errors.New().
			SetStatus(http.StatusInternalServerError).
			SetLayer("auth.service").
			SetFunction("Register").
			SetMessage("failed to generate access token").
			SetError(errToken)
	}

	tokenRefresh, errTokenRefresh := auth.GenerateToken(created.ID, expiresAtRefresh)
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
