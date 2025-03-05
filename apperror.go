package apperror

import (
	"encoding/json"
	"fmt"
)

type AppError struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	HTTPStatus int         `json:"http_status"`
	Details    interface{} `json:"details,omitempty"`
	stackTrace *stackTrace
	Err        error `json:"-"`
}

func NewAppError(code string, message string, httpStatus int, err error) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
		stackTrace: captureStackTrace(),
		Err:        err,
	}
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// StackTrace returns the stack trace associated with the AppError.
// It provides detailed information about the sequence of function calls
// that led to the error, which can be useful for debugging purposes.
func (e *AppError) StackTrace() *stackTrace {
	return e.stackTrace
}

func (e *AppError) MarshalJSON() ([]byte, error) {
	type Alias AppError // Avoid infinite recursion

	return json.Marshal(&struct {
		*Alias
		StackTrace *stackTrace `json:"stack_trace,omitempty"`
	}{
		Alias:      (*Alias)(e),
		StackTrace: e.stackTrace,
	})
}
