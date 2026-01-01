package errors

type AppError struct {
	Code     string
	Status   int
	Layer    string
	Message  string
	Function string
	Err      error
}

func New() *AppError {
	return &AppError{}
}

func (e *AppError) SetCode(code string) *AppError {
	e.Code = code
	return e
}

func (e *AppError) SetStatus(status int) *AppError {
	e.Status = status
	return e
}

func (e *AppError) SetLayer(layer string) *AppError {
	e.Layer = layer
	return e
}

func (e *AppError) SetMessage(message string) *AppError {
	e.Message = message
	return e
}

func (e *AppError) SetFunction(fn string) *AppError {
	e.Function = fn
	return e
}

func (e *AppError) SetError(err error) *AppError {
	e.Err = err
	return e
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}
