package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Akiles94/go-test-api/services/product/config"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/infra/adapters/module"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/infra/adapters/repository"
	"github.com/Akiles94/go-test-api/services/product/db"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	grpc_services "github.com/Akiles94/go-test-api/shared/infra/grpc/services"
	"github.com/Akiles94/go-test-api/shared/infra/middlewares"
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
	// Load configuration
	config.LoadEnv()

	// Set Gin mode
	if config.Env.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize database
	database := db.Connect()

	if err := database.AutoMigrate(&repository.ProductEntity{}); err != nil {
		log.Fatalf("❌ DB migration failed: %v", err)
	}

	// Initialize router
	router := gin.New()

	var modules []shared_ports.ModulePort

	// Product module
	productModule := module.NewProductModule(database)
	modules = append(modules, productModule)

	serviceRegistryClientConfig := grpc_services.ServiceRegistryClientConfig{
		GatewayAddress: config.Env.GatewayGRPCAddress,
		Context:        context.Background(),
		ServiceName:    "product-service",
		ServiceVersion: "v1.0.0",
		ServiceURL:     fmt.Sprintf("http://%s:%s", config.Env.ServiceHost, config.Env.ApiPort),
		HealthEndpoint: "/health",
		Modules:        modules,
	}

	serviceRegistry, err := grpc_services.NewServiceRegistryClient(serviceRegistryClientConfig)
	if err != nil {
		log.Printf("⚠️ Failed to create service registry: %v", err)
		return
	}

	if err := serviceRegistry.RegisterWithGateway(); err != nil {
		log.Printf("⚠️ Failed to register with gateway: %v", err)
		log.Printf("Continuing without gateway registration...")
	}

	defer func() {
		if err := serviceRegistry.DeregisterFromGateway(); err != nil {
			log.Printf("⚠️ Failed to deregister from gateway: %v", err)
		}
	}()

	// Start server
	startServer(router, modules)
}

func startServer(router *gin.Engine, modules []shared_ports.ModulePort) {
	// Product Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "product-service",
			"version": "1.0.0",
		})
	})

	// Add middlewares
	router.Use(middlewares.StructuredLogger())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.RequestIDMiddleware())
	router.Use(middlewares.SecurityHeadersMiddleware())

	for _, item := range modules {
		switch mod := item.(type) {
		case *module.ProductModule:
			mod.RegisterRoutes(router.Group(mod.GetPathPrefix()))
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
