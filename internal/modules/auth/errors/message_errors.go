package errors

import "gym-api/internal/utils/errors"

var serviceBase = errors.ServiceError("auth.service")
var repositoryBase = errors.ServiceError("auth.repository")

var (
	ErrInvalidCredentials = errors.Unauthorized(
		ErrInvalidCredentialsCode,
		"invalid credentials",
	)

	ErrGenerateToken = errors.Unauthorized(
		ErrGenerateTokenCode,
		"invalid token",
	)

	ErrGenerateHashPassword = errors.Internal(
		ErrGenerateHashCode,
		"invalid token",
	)
)
