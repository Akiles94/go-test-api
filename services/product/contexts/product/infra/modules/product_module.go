package modules

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/product/application/use_cases"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/infra/adapters"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/infra/handlers"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
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
	handler := handlers.NewProductHandler(
		createProductUseCase,
		updateProductUseCase,
		patchProductUseCase,
		deleteProductUseCase,
		getAllProductsUseCase,
		getOneProductUseCase)

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

func (m *ProductModule) GetServiceInfo() value_objects.ServiceInfo {
	return value_objects.ServiceInfo{
		Name:    "product-service",
		URL:     "http://localhost:8081", // or from config
		Health:  "/health",
		Version: "1.0.0",
		Status:  "healthy",
		Routes: []value_objects.RouteDefinition{
			{
				Method:    "GET",
				Path:      "/products",
				Protected: false,
				RateLimit: 100,
			},
			{
				Method:    "POST",
				Path:      "/products",
				Protected: true,
				RateLimit: 50,
			},
			{
				Method:    "GET",
				Path:      "/products/:id",
				Protected: false,
				RateLimit: 200,
			},
			{
				Method:    "PUT",
				Path:      "/products/:id",
				Protected: true,
				RateLimit: 30,
			},
			{
				Method:    "DELETE",
				Path:      "/products/:id",
				Protected: true,
				RateLimit: 10,
			},
		},
	}
}

func (pm *ProductModule) RegisterWithGateway(ctx context.Context, registry shared_ports.ServiceRegistryPort) error {
	serviceInfo := pm.GetServiceInfo()
	return registry.Register(ctx, serviceInfo)
}
