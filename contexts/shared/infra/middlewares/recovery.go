package middlewares

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, recovered interface{}) {
		logrus.WithFields(logrus.Fields{
			"panic":      recovered,
			"stack":      string(debug.Stack()),
			"path":       c.Request.URL.Path,
			"method":     c.Request.Method,
			"request_id": c.GetHeader("X-Request-ID"),
		}).Error("Panic recovered")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":      "internal server error",
			"request_id": c.GetHeader("X-Request-ID"),
		})
	})
}
