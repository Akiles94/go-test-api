package module

import (
	"github.com/Akiles94/go-test-api/services/user/contexts/auth/application/use_cases"
	"github.com/Akiles94/go-test-api/services/user/contexts/auth/infra/adapters/hasher"
	"github.com/Akiles94/go-test-api/services/user/contexts/auth/infra/adapters/repository"
	"github.com/Akiles94/go-test-api/services/user/contexts/auth/infra/handlers"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/infra/shared_adapters"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthModule struct {
	pathPrefix string
	handler    *handlers.AuthHandler
	routes     []shared_ports.RouteDefinition
}

const pathPrefix = "/auth"

func NewAuthModule(db *gorm.DB) *AuthModule {
	repo := repository.NewUserRepository(db)
	hasher := hasher.NewHasher()
	jwtService := shared_adapters.NewJWTService("some-secret-key")
	authorizer := shared_adapters.NewAuthService(jwtService)
	loginUseCase := use_cases.NewLoginUseCase(repo, hasher, authorizer)
	registerUseCase := use_cases.NewRegisterUseCase(repo, hasher)
	handler := handlers.NewAuthHandler(loginUseCase, registerUseCase, jwtService)
	routes := []shared_ports.RouteDefinition{
		{
			Method:    "POST",
			Path:      "/login",
			Protected: false,
			Handler:   handler.Login,
			RateLimit: 10,
		},
		{
			Method:    "POST",
			Path:      "/register",
			Protected: false,
			Handler:   handler.Register,
			RateLimit: 10,
		},
		{
			Method:    "GET",
			Path:      "/validate",
			Protected: false,
			Handler:   handler.ValidateToken,
			RateLimit: 10,
		},
	}
	return &AuthModule{pathPrefix: pathPrefix, handler: handler, routes: routes}
}

func (am *AuthModule) RegisterRoutes(router *gin.RouterGroup) {
	routes := am.routes

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

func (am *AuthModule) GetPathPrefix() string {
	return am.pathPrefix
}
