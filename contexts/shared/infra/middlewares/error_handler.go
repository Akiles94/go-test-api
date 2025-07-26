package middlewares

import (
	"net/http"

	"github.com/Akiles94/go-test-api/contexts/shared/application/dto"
	"github.com/Akiles94/go-test-api/contexts/shared/domain/models"
	"github.com/Akiles94/go-test-api/contexts/shared/infra/handlers"
	"github.com/gin-gonic/gin"
)

const ErrorKey = "middleware_error"

const defaultDomainErrorStatus = http.StatusBadRequest

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			handleErrorResponse(c, err)
			return
		}

		if recovered := recover(); recovered != nil {
			var err error
			switch panicErr := recovered.(type) {
			case models.DomainError:
				err = panicErr
			case handlers.InfraError:
				err = panicErr
			case error:
				err = panicErr
			default:
				err = models.DomainError{
					Code:    string(handlers.ErrorCodePanicError),
					Message: "An unexpected error occurred",
				}
			}
			handleErrorResponse(c, err)
		}
	})
}

func handleErrorResponse(c *gin.Context, err error) {
	statusCode := getStatusCodeFromError(err)
	var errorResponse dto.ErrorResponse

	if _, ok := err.(models.DomainError); ok {
		errorResponse = dto.FromDomainError(err)
	}
	if _, ok := err.(handlers.InfraError); ok {
		errorResponse = dto.FromInfraError(err)
	}

	c.Header("Content-Type", "application/json")
	c.JSON(statusCode, errorResponse)
	c.Abort()
}

func getStatusCodeFromError(err error) int {
	if _, ok := err.(models.DomainError); ok {
		return defaultDomainErrorStatus
	}
	if infraErr, ok := err.(handlers.InfraError); ok {
		if statusCode, exists := handlers.InfraErrorStatusMap[handlers.ErrorCode(infraErr.Code)]; exists {
			return statusCode
		}
	}
	return http.StatusInternalServerError
}
