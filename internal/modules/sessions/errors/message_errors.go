package errors

import "gym-api/internal/utils/errors"

var (
	ErrCreateSession = errors.Internal(
		ErrCreateSessionCode,
		"failed to create session",
	)

	ErrFindSessionById = errors.Internal(
		ErrFindSessionByIdCode,
		"failed to find session by id",
	)
)
