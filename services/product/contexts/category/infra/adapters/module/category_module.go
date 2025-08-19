package module

import (
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/use_cases"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/infra/adapters/repository"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/infra/handlers"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/infra/grpc/gen/registry"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryModule struct {
	pathPrefix string
	handler    *handlers.CategoryHandler
	routes     []shared_ports.RouteDefinition
}

const pathPrefix = "/categories"

func NewCategoryModule(db *gorm.DB) *CategoryModule {
	repo := repository.NewCategoryRepository(db)
	createCategoryUseCase := use_cases.NewCreateCategoryUseCase(repo)
	updateCategoryUseCase := use_cases.NewUpdateCategoryUseCase(repo)
	patchCategoryUseCase := use_cases.NewPatchCategoryUseCase(repo)
	deleteCategoryUseCase := use_cases.NewDeleteCategoryUseCase(repo)
	getAllCategoriesUseCase := use_cases.NewGetAllCategoriesUseCase(repo)
	getOneCategoryUseCase := use_cases.NewGetOneCategoryUseCase(repo)
	handler := handlers.NewCategoryHandler(
		createCategoryUseCase,
		updateCategoryUseCase,
		patchCategoryUseCase,
		deleteCategoryUseCase,
		getAllCategoriesUseCase,
		getOneCategoryUseCase)
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

	return &CategoryModule{pathPrefix: pathPrefix, routes: routes, handler: handler}
}

func (cm *CategoryModule) RegisterRoutes(router *gin.RouterGroup) {
	routes := cm.routes

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

func (cm *CategoryModule) GetRouteDefinitions() []*registry.RouteInfo {
	routeInfos := make([]*registry.RouteInfo, len(cm.routes))

	for i, route := range cm.routes {
		routeInfos[i] = &registry.RouteInfo{
			Method:    route.Method,
			Path:      cm.pathPrefix + route.Path,
			Protected: route.Protected,
			RateLimit: route.RateLimit,
		}
	}

	return routeInfos
}

func (cm *CategoryModule) GetPathPrefix() string {
	return cm.pathPrefix
}
