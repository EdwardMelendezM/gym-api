package errors

type Layers string

const (
	LayerService    Layers = "service"
	LayerRepository Layers = "repository"
	LayerHandler    Layers = "handler"
)

type DefaultCodes string

const (
	CodeInternalError DefaultCodes = "INTERNAL_ERROR"
	CodeBadRequest    DefaultCodes = "BAD_REQUEST"
	CodeUnauthorized  DefaultCodes = "UNAUTHORIZED"
	CodeForbidden     DefaultCodes = "FORBIDDEN"
	CodeNotFound      DefaultCodes = "NOT_FOUND"
	CodeConflict      DefaultCodes = "CONFLICT"
)

type AppError struct {
	Code     string
	Status   int
	Layer    string
	Message  string
	Fields   map[string]string
	Function string
	Err      error
}

func New() *AppError {
	return &AppError{}
}

func (e *AppError) SetFields(fields map[string]string) *AppError {
	e.Fields = fields
	return e
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
