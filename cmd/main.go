package main

import (
	"log"

	"github.com/Akiles94/go-test-api/config"
	"github.com/Akiles94/go-test-api/contexts/products/infra/adapters"
	"github.com/Akiles94/go-test-api/contexts/products/infra/modules"
	"github.com/Akiles94/go-test-api/contexts/shared/application/interfaces"
	"github.com/Akiles94/go-test-api/contexts/shared/infra/middlewares"
	"github.com/Akiles94/go-test-api/db"
	"github.com/gin-gonic/gin"
)

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
	api := router.Group("/api")

	var appModules []interfaces.Module

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
