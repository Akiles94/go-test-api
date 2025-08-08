package main

import (
	"log"

	"github.com/Akiles94/go-test-api/gateway/config"
	"github.com/Akiles94/go-test-api/gateway/routes"
	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
	"github.com/Akiles94/go-test-api/shared/infra/middlewares"
	"github.com/Akiles94/go-test-api/shared/infra/shared_adapters"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.LoadEnv()

	if config.Env.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Services
	jwtService := shared_adapters.NewJWTService(config.Env.JWTSecret)
	authService := shared_adapters.NewAuthService(jwtService)

	// Service registry (in-memory for start, could be Redis/Consul later)
	registry := shared_adapters.NewInMemoryServiceRegistry()

	router := gin.New()

	// Static middlewares
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.StructuredLogger())
	router.Use(middlewares.SecurityHeadersMiddleware())
	router.Use(middlewares.ErrorHandlerMiddleware())

	// Health and docs (static routes)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "api-gateway"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Dynamic routes setup
	dynamicRouter := routes.NewDynamicRouter(registry, authService)
	dynamicRouter.SetupRoutes(router)

	// Start discovery endpoint for services to register
	router.POST("/gateway/register", func(c *gin.Context) {
		var serviceInfo value_objects.ServiceInfo
		if err := c.ShouldBindJSON(&serviceInfo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := registry.Register(c.Request.Context(), serviceInfo); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "service registered successfully"})
	})

	log.Printf("🚀 API Gateway starting on port %s", config.Env.ApiPort)
	if err := router.Run(":" + config.Env.ApiPort); err != nil {
		log.Fatalf("❌ Error starting server: %v", err)
	}
}
