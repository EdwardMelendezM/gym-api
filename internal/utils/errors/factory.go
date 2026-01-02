package errors

func ServiceError(layer string) *AppError {
	return New().SetLayer(layer)
}
