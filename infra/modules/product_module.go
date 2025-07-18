package modules

import (
	"github.com/Akiles94/go-test-api/application/use_cases"
	"github.com/Akiles94/go-test-api/infra/adapters"
	"github.com/Akiles94/go-test-api/infra/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductModule struct {
	handler *handlers.ProductHandler
}

func NewProductModule(db *gorm.DB) *ProductModule {
	repo := adapters.NewProductRepository(db)
	useCase := use_cases.NewProductUseCase(repo)
	handler := handlers.NewProductHandler(useCase)

	return &ProductModule{handler: handler}
}

func (pm *ProductModule) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("", pm.handler.GetPaginated)
	router.GET("/:id", pm.handler.GetByID)
	router.POST("", pm.handler.Create)
	router.PUT("/:id", pm.handler.Update)
	router.PATCH("/:id", pm.handler.Patch)
	router.DELETE("/:id", pm.handler.Delete)
}
