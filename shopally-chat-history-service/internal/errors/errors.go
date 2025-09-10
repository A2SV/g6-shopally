package errors

import "errors"

// Common application-specific errors.
var (
	ErrNotFound     = errors.New("resource not found")
	ErrInvalidInput = errors.New("invalid input provided")
	ErrConflict     = errors.New("resource conflict")
	ErrInternal     = errors.New("internal server error")
	ErrUnauthorized = errors.New("unauthorized access")
)

// Is checks if an error is of a specific type.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As finds the first error in a chain that matches type target.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
