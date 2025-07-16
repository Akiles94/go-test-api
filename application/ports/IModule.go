package ports

import "github.com/gin-gonic/gin"

type IModule interface {
	RegisterRoutes(router *gin.RouterGroup)
}
