package modules

import (
	"github.com/Akiles94/go-test-api/products/application/use_cases"
	"github.com/Akiles94/go-test-api/products/infra/adapters"
	"github.com/Akiles94/go-test-api/products/infra/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductModule struct {
	handler *handlers.ProductHandler
}

func NewProductModule(db *gorm.DB) *ProductModule {
	repo := adapters.NewProductRepository(db)
	createProductUseCase := use_cases.NewCreateProductUseCase(repo)
	updateProductUseCase := use_cases.NewUpdateProductUseCase(repo)
	patchProductUseCase := use_cases.NewPatchProductUseCase(repo)
	deleteProductUseCase := use_cases.NewDeleteProductUseCase(repo)
	getAllProductsUseCase := use_cases.NewGetAllProductsUseCase(repo)
	getOneProductUseCase := use_cases.NewGetOneProductUseCase(repo)
	handler := handlers.NewProductHandler(createProductUseCase, updateProductUseCase, patchProductUseCase, deleteProductUseCase, getAllProductsUseCase, getOneProductUseCase)

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
