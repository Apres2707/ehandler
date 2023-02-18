package ehandler

import (
	"errors"
)

const systemErrorMsg = "Произошла ошибка при обработке запроса. Пожалуйста, повторите попытку позже."

// Handle is a method that returns one or another code and error message depending on whether it is allowed to
// throw system errors.
func Handle(err error, showSystem bool) (code, msg string) {
	if hasOnlyAppError(err) {
		return err.(AppError).Code(), err.Error()
	}

	if showSystem {
		return systemError, err.Error()
	}

	return systemError, systemErrorMsg
}

func hasOnlyAppError(err error) bool {
	if _, ok := err.(AppError); !ok {
		return false
	}

	parentErr := errors.Unwrap(err)
	if parentErr == nil {
		return true
	}

	if _, ok := parentErr.(AppError); !ok {
		return false
	}

	return hasOnlyAppError(parentErr)
}
