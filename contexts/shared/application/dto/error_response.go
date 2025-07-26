package dto

import (
	"github.com/Akiles94/go-test-api/contexts/shared/domain/models"
	shared_handlers "github.com/Akiles94/go-test-api/contexts/shared/infra/handlers"
)

type ErrorResponse struct {
	Error   string                 `json:"error"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func FromDomainError(domainErr error) ErrorResponse {
	if de, ok := domainErr.(models.DomainError); ok {
		return ErrorResponse{
			Error:   de.Code,
			Message: de.Message,
			Details: de.Details,
		}
	}

	return ErrorResponse{
		Error:   "INTERNAL_ERROR",
		Message: "An unexpected error occurred",
	}
}

func FromInfraError(infraErr error) ErrorResponse {
	if ie, ok := infraErr.(shared_handlers.InfraError); ok {
		return ErrorResponse{
			Error:   string(ie.Code),
			Message: ie.Message,
			Details: nil,
		}
	}

	return ErrorResponse{
		Error:   "INTERNAL_ERROR",
		Message: "An unexpected error occurred",
	}
}
