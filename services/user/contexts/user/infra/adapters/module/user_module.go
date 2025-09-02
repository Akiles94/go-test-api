package module

import (
	"github.com/Akiles94/go-test-api/services/user/contexts/user/application/use_cases"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/infra/adapters/hasher"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/infra/adapters/repository"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/infra/handlers"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/infra/shared_adapters"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserModule struct {
	pathPrefix string
	handler    *handlers.UserHandler
	routes     []shared_ports.RouteDefinition
}

const pathPrefix = "/users"

func NewUserModule(db *gorm.DB) *UserModule {
	repo := repository.NewUserRepository(db)
	hasher := hasher.NewHasher()
	jwtService := shared_adapters.NewJWTService("some-secret-key")
	authorizer := shared_adapters.NewAuthService(jwtService)
	loginUseCase := use_cases.NewLoginUseCase(repo, hasher, authorizer)
	registerUseCase := use_cases.NewRegisterUseCase(repo, hasher)
	handler := handlers.NewUserHandler(loginUseCase, registerUseCase)
	routes := []shared_ports.RouteDefinition{
		{
			Method:    "POST",
			Path:      "/auth/login",
			Protected: false,
			Handler:   handler.Login,
			RateLimit: 10,
		},
		{
			Method:    "POST",
			Path:      "/auth/register",
			Protected: false,
			Handler:   handler.Register,
			RateLimit: 10,
		},
		{
			Method:    "GET",
			Path:      "/auth/validate",
			Protected: false,
			Handler:   handler.ValidateToken,
			RateLimit: 10,
		},
	}

	return &UserModule{pathPrefix: pathPrefix, handler: handler, routes: routes}
}

func (um *UserModule) RegisterRoutes(router *gin.RouterGroup) {
	routes := um.routes

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

func (um *UserModule) GetPathPrefix() string {
	return um.pathPrefix
}
