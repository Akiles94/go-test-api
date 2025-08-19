package module

import (
	"github.com/Akiles94/go-test-api/services/product/contexts/product/application/use_cases"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/infra/adapters/repository"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/infra/handlers"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/infra/grpc/gen/registry"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductModule struct {
	pathPrefix string
	handler    *handlers.ProductHandler
	routes     []shared_ports.RouteDefinition
}

const pathPrefix = "/products"

func NewProductModule(db *gorm.DB) *ProductModule {
	repo := repository.NewProductRepository(db)
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
	routes := []shared_ports.RouteDefinition{
		{
			Method:    "GET",
			Path:      "",
			Protected: false,
			Handler:   handler.GetPaginated,
			RateLimit: 10,
		},
		{
			Method:    "GET",
			Path:      "/:id",
			Protected: false,
			Handler:   handler.GetByID,
			RateLimit: 10,
		},
		{
			Method:    "POST",
			Path:      "",
			Protected: false,
			Handler:   handler.Create,
			RateLimit: 10,
		},
		{
			Method:    "PUT",
			Path:      "/:id",
			Protected: true,
			Handler:   handler.Update,
			RateLimit: 10,
		},
		{
			Method:    "PATCH",
			Path:      "/:id",
			Protected: true,
			Handler:   handler.Patch,
			RateLimit: 10,
		},
		{
			Method:    "DELETE",
			Path:      "/:id",
			Protected: true,
			Handler:   handler.Delete,
			RateLimit: 10,
		},
	}

	return &ProductModule{pathPrefix: pathPrefix, routes: routes}
}

func (pm *ProductModule) RegisterRoutes(router *gin.RouterGroup) {
	routes := pm.routes

	for _, route := range routes {
		switch route.Method {
		case "GET":
			router.GET(route.Path, route.Handler)
		case "POST":
			router.POST(route.Path, route.Handler)
		case "PUT":
			router.PUT(route.Path, route.Handler)
		case "PATCH":
			router.PATCH(route.Path, route.Handler)
		case "DELETE":
			router.DELETE(route.Path, route.Handler)
		}
	}
}

func (pm *ProductModule) GetRouteDefinitions() []*registry.RouteInfo {
	routeInfos := make([]*registry.RouteInfo, len(pm.routes))

	for i, route := range pm.routes {
		routeInfos[i] = &registry.RouteInfo{
			Method:    route.Method,
			Path:      pm.pathPrefix + route.Path,
			Protected: route.Protected,
			RateLimit: route.RateLimit,
		}
	}

	return routeInfos
}

func (pm *ProductModule) GetPathPrefix() string {
	return pm.pathPrefix
}
