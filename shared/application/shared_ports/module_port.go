package shared_ports

import "github.com/gin-gonic/gin"

type ModulePort interface {
	RegisterRoutes(router *gin.RouterGroup)
}
