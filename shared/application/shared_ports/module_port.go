package shared_ports

import (
	"github.com/gin-gonic/gin"
)

type RouteDefinition struct {
	Method    string
	Path      string
	Protected bool
	Handler   gin.HandlerFunc
	RateLimit int32
}

type ModulePort interface {
	RegisterRoutes(router *gin.RouterGroup)
	GetPathPrefix() string
}
