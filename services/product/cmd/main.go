package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Akiles94/go-test-api/services/product/config"
	category_module "github.com/Akiles94/go-test-api/services/product/contexts/category/infra/adapters/module"
	product_module "github.com/Akiles94/go-test-api/services/product/contexts/product/infra/adapters/module"
	"github.com/Akiles94/go-test-api/services/product/db"
	"github.com/Akiles94/go-test-api/services/product/shared/infra/adapters/repository"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/infra/middlewares"
)

func main() {
	// Load configuration
	config.LoadEnv()

	// Set Gin mode
	if config.Env.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize database
	database := db.Connect()

	if err := database.AutoMigrate(
		&repository.ProductEntity{},
		&repository.CategoryEntity{},
	); err != nil {
		log.Fatalf("❌ DB migration failed: %v", err)
	}

	// Initialize router
	router := gin.New()

	var modules []shared_ports.ModulePort

	// Product module
	productModule := product_module.NewProductModule(database)
	categoryModule := category_module.NewCategoryModule(database)
	modules = append(modules, productModule, categoryModule)

	// Start server
	startServer(router, modules)
}

func startServer(router *gin.Engine, modules []shared_ports.ModulePort) {
	// Product Health check
	router.GET("/products/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "product-service",
			"version": "1.0.0",
		})
	})

	// Add middlewares
	router.Use(middlewares.StructuredLogger())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.RequestIDMiddleware())
	router.Use(middlewares.ErrorHandlerMiddleware())
	router.Use(middlewares.SecurityHeadersMiddleware())
	apiV1 := router.Group("/api/v1")

	for _, item := range modules {
		switch mod := item.(type) {
		case *product_module.ProductModule:
			mod.RegisterRoutes(apiV1.Group(mod.GetPathPrefix()))
		case *category_module.CategoryModule:
			mod.RegisterRoutes(apiV1.Group(mod.GetPathPrefix()))
		}
	}

	server := &http.Server{
		Addr:         ":" + config.Env.ApiPort,
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Printf("🚀 Product service starting on port %s", config.Env.ApiPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down product service...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("✅ Product service stopped gracefully")
}
