package errors

import "gym-api/internal/utils/errors"

var (
	ErrInvalidCredentials = errors.Unauthorized(
		"INVALID_CREDENTIALS",
		"invalid credentials",
	)

	ErrTokenExpired = errors.Unauthorized(
		"TOKEN_EXPIRED",
		"token expired",
	)

	ErrSessionNotFound = errors.NotFound(
		"SESSION_NOT_FOUND",
		"session not found",
	)
)
