package shared_handlers

import (
	"net/http"
)

type InfraError struct {
	Code    ErrorCode
	Message string
}

func (e InfraError) Error() string {
	return e.Message
}

type ErrorCode string

const (
	ErrorCodeNotFound        ErrorCode = "NOT_FOUND"
	ErrorCodeUnauthorized    ErrorCode = "UNAUTHORIZED"
	ErrorCodeForbidden       ErrorCode = "FORBIDDEN"
	ErrorCodeValidationError ErrorCode = "VALIDATION_ERROR"
	ErrorCodePanicError      ErrorCode = "PANIC_ERROR"
	ErrorCodeBadRequest      ErrorCode = "BAD_REQUEST"
	ErrorCodeInternalError   ErrorCode = "INTERNAL_ERROR"
)

var (
	ErrInvalidCursor = InfraError{
		Code:    ErrorCodeValidationError,
		Message: "Invalid cursor format",
	}

	ErrInvalidLimit = InfraError{
		Code:    ErrorCodeValidationError,
		Message: "Invalid limit value",
	}

	ErrInvalidUUID = InfraError{
		Code:    ErrorCodeValidationError,
		Message: "Invalid UUID format",
	}

	ErrInvalidPayload = InfraError{
		Code:    ErrorCodeBadRequest,
		Message: "invalid payload",
	}

	ErrNotFound = InfraError{
		Code:    ErrorCodeNotFound,
		Message: "Resource not found",
	}
)

var InfraErrorStatusMap = map[ErrorCode]int{
	ErrorCodeNotFound:        http.StatusNotFound,
	ErrorCodeUnauthorized:    http.StatusUnauthorized,
	ErrorCodeForbidden:       http.StatusForbidden,
	ErrorCodeValidationError: http.StatusBadRequest,
	ErrorCodePanicError:      http.StatusInternalServerError,
	ErrorCodeBadRequest:      http.StatusBadRequest,
	ErrorCodeInternalError:   http.StatusInternalServerError,
}
