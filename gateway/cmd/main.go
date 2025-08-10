package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Akiles94/go-test-api/gateway/config"
	"github.com/Akiles94/go-test-api/gateway/routes"
	"github.com/Akiles94/go-test-api/shared/infra/grpc/gen/registry"
	grpc_services "github.com/Akiles94/go-test-api/shared/infra/grpc/services"
	"github.com/Akiles94/go-test-api/shared/infra/middlewares"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	log.Println("🚀 Starting Gateway...")

	config.LoadEnv()

	ctx := context.Background()

	// 1. Create gRPC Service Registry Server
	serviceRegistryServer := grpc_services.NewServiceRegistryServer()

	// 2. Start gRPC server in goroutine
	go startGRPCServer(serviceRegistryServer)

	// 3. Create HTTP router with Gin
	// 3. Create HTTP router with Gin
	router := gin.New()

	// 4. Add middlewares
	router.Use(middlewares.StructuredLogger())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.SecurityHeadersMiddleware())
	router.Use(func(c *gin.Context) {
		c.Header("X-Gateway", "go-test-api")
		c.Next()
	})

	// 5. Configure static routes
	setupGatewayRoutes(router, serviceRegistryServer)

	// 6. Configure dynamic routes
	dynamicRouter := routes.NewDynamicRouter(ctx, router, serviceRegistryServer)
	go dynamicRouter.WatchServices() // Watch for service changes

	// ✅ 7. Configure HTTP server properly
	httpServer := &http.Server{
		Addr:         ":" + config.Env.ApiPort,
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// ✅ 8. Start HTTP server in goroutine
	go func() {
		log.Printf("🌐 HTTP Gateway server starting on port %s", config.Env.ApiPort)
		log.Printf("📋 Environment: %s", config.Env.Mode)

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Failed to start HTTP server: %v", err)
		}
	}()

	// ✅ 9. Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down Gateway...")

	// Shutdown HTTP server with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 30*time.Second)
	defer shutdownCancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("❌ HTTP server forced to shutdown: %v", err)
	}

	log.Println("✅ Gateway stopped gracefully")
}

func startGRPCServer(serviceRegistryServer *grpc_services.ServiceRegistryServer) {
	// Create gRPC listener
	grpcAddress := fmt.Sprintf("%s:%s", config.Env.GRPCHost, config.Env.GRPCPort)
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("❌ Failed to listen on port %s: %v", config.Env.GRPCPort, err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register the service registry server
	registry.RegisterServiceRegistryServer(grpcServer, serviceRegistryServer)

	log.Printf("🔌 gRPC Service Registry server starting on %s", grpcAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve gRPC: %v", err)
	}
}

func setupGatewayRoutes(router *gin.Engine, serviceRegistry *grpc_services.ServiceRegistryServer) {
	// Gateway admin routes
	admin := router.Group("/gateway")

	// Endpoint to see all registered services
	admin.GET("/services", func(c *gin.Context) {
		services, err := serviceRegistry.GetServices(c.Request.Context(), nil)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"total_services": len(services.Services),
			"services":       services.Services,
		})
	})

	// Health check del Gateway
	admin.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "gateway",
			"version": "1.0.0",
		})
	})

	// Gateway info endpoint
	admin.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service":     "go-test-api-gateway",
			"version":     "1.0.0",
			"grpc_port":   config.Env.GRPCPort,
			"http_port":   config.Env.ApiPort,
			"description": "API Gateway with gRPC service registry",
		})
	})
}
