package apperror

type AppError struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	HTTPStatus int         `json:"http_status"`
	Details    interface{} `json:"details,omitempty"`
	StackTrace []string    `json:"stack_trace,omitempty"`
	Err        error       `json:"-"`
}

func NewAppError(code string, message string, httpStatus int, err error) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
		StackTrace: captureStackTrace(),
		Err:        err,
	}
}
