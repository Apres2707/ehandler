package ehandler

const (
	systemError  = "systemError"
	processError = "processError"
)

// AppError is an interface that must be implemented by any type that should be considered a non-system error.
type AppError interface {
	error
	// Code is the method that returns the error code.
	Code() string
}

type baseError struct {
	msg,
	code string
	cause error
}

// ProcessError is a packaged type that is not a system error.
type ProcessError struct{ baseError }

// Error is a method that returns a textual representation of the error.
func (e baseError) Error() string {
	return e.msg
}

// Code returns the code that matches the error type.
func (e baseError) Code() string {
	return e.code
}

// Unwrap returns an error, if any, that has been wrapped with the current error.
func (e baseError) Unwrap() error {
	return e.cause
}

// NewProcessError is a method to create an error with type ProcessError.
func NewProcessError(msg string, cause error) ProcessError {
	return ProcessError{
		newBaseError(msg, processError, cause),
	}
}

func newBaseError(msg, code string, cause error) baseError {
	return baseError{
		msg:   msg,
		code:  code,
		cause: cause,
	}
}
