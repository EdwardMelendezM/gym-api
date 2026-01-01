package errors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FromValidation(err error) *AppError {
	var ve validator.ValidationErrors

	// Case 1: real validation
	if errors.As(err, &ve) {
		fields := make(map[string]string)

		for _, fe := range ve {
			field := toLowerCamel(fe.Field())
			fields[field] = normalizeRule(fe.Tag())
		}

		return BadRequest("VALIDATION_ERROR", "validation failed").
			SetFields(fields)
	}

	return BadRequest("INVALID_REQUEST", "invalid request body")
}

func normalizeRule(tag string) string {
	switch tag {
	case "required":
		return "required"
	case "email":
		return "invalid"
	case "min":
		return "min"
	default:
		return tag
	}
}

func toLowerCamel(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
