package main

import (
	"log"

	"github.com/Akiles94/go-test-api/config"
	"github.com/Akiles94/go-test-api/db"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/infra/adapters"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/infra/modules"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/infra/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Akiles94/go-test-api/docs"
)

// @title Go Test API
// @version 1.0
// @description A test API for products management
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
func main() {
	config.LoadEnv()

	database := db.Connect()

	if err := database.AutoMigrate(&adapters.ProductEntity{}); err != nil {
		log.Fatalf("❌ DB migration failed: %v", err)
	}

	if config.Env.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middlewares.RequestIDMiddleware())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.StructuredLogger())
	router.Use(middlewares.SecurityHeadersMiddleware())
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.RateLimitMiddleware())
	router.Use(middlewares.ErrorHandlerMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "go-test-api"})
	})
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})
	api := router.Group("/api/v1")

	var appModules []shared_ports.ModulePort

	productModule := modules.NewProductModule(database)
	appModules = append(appModules, productModule)

	for _, m := range appModules {
		switch mod := m.(type) {
		case *modules.ProductModule:
			mod.RegisterRoutes(api.Group("/products"))
		}
	}

	log.Printf("🚀 Server starting on port %s", config.Env.ApiPort)
	if err := router.Run(":" + config.Env.ApiPort); err != nil {
		log.Fatalf("❌ Error starting server: %v", err)
	}
}
