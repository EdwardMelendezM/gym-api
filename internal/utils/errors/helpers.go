package errors

import "net/http"

func BadRequest(code, msg string) *AppError {
	return New().
		SetStatus(http.StatusBadRequest).
		SetCode(string(CodeBadRequest)).
		SetMessage(msg)
}

func Unauthorized(code, msg string) *AppError {
	return New().
		SetStatus(http.StatusUnauthorized).
		SetCode(code).
		SetMessage(msg)
}

func Forbidden(code, msg string) *AppError {
	return New().
		SetStatus(http.StatusForbidden).
		SetCode(code).
		SetMessage(msg)
}

func NotFound(code, msg string) *AppError {
	return New().
		SetStatus(http.StatusNotFound).
		SetCode(code).
		SetMessage(msg)
}

func Conflict(code, msg string) *AppError {
	return New().
		SetStatus(http.StatusConflict).
		SetCode(code).
		SetMessage(msg)
}

func Internal(code, msg string) *AppError {
	return New().
		SetStatus(http.StatusInternalServerError).
		SetCode(code).
		SetMessage(msg)
}

func WithContext(err *AppError, layer, fn string) *AppError {
	return err.
		SetLayer(layer).
		SetFunction(fn)
}

func Wrap(err error, msg string, code string) *AppError {
	return Internal(msg, string(CodeBadRequest)).SetError(err)
}

func WrapWithContext(err error, msg, layer, fn string) *AppError {
	return Internal(msg, string(CodeInternalError)).
		SetLayer(layer).
		SetFunction(fn).
		SetError(err)
}

func WrapError(err *AppError) *AppError {
	return err
}
