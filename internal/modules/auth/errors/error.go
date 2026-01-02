package errors

import (
	"gym-api/internal/utils/errors"
	"net/http"
)

var (
	ErrorUserNotFound = errors.New().
				SetStatus(http.StatusNotFound).
				SetCode(ErrUserNotFoundCode).
				SetMessage("user not found")

	ErrUserAlreadyExists = errors.New().
				SetStatus(http.StatusConflict).
				SetCode(ErrUserAlreadyExistsCode).
				SetMessage("user not found")

	ErrorInvalidCredentials = errors.New().
				SetStatus(http.StatusUnauthorized).
				SetCode(ErrInvalidCredentialsCode).
				SetMessage("invalid credentials")

	ErrorGenerateToken = errors.New().
				SetStatus(http.StatusUnauthorized).
				SetCode(ErrGenerateTokenCode).
				SetMessage("invalid token")

	ErrorGenerateHashPassword = errors.New().
					SetStatus(http.StatusInternalServerError).
					SetCode(ErrGenerateHashCode).
					SetMessage("error generating hash password")
)
