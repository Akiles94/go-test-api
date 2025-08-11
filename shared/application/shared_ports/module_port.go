package shared_ports

import (
	"github.com/Akiles94/go-test-api/shared/infra/grpc/gen/registry"
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
	GetRouteDefinitions() []*registry.RouteInfo
	GetPathPrefix() string
}
