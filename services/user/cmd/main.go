package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Akiles94/go-test-api/services/product/db"
	"github.com/Akiles94/go-test-api/services/user/config"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/infra/adapters/module"
	"github.com/Akiles94/go-test-api/services/user/contexts/user/infra/adapters/repository"
	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/infra/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	if config.Env.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	database := db.Connect()

	if err := database.AutoMigrate(
		&repository.UserEntity{},
	); err != nil {
		log.Fatalf("❌ DB migration failed: %v", err)
	}

	router := gin.New()

	var modules []shared_ports.ModulePort

	userModule := module.NewUserModule(database)
	modules = append(modules, userModule)

	startServer(router, modules)
}

func startServer(router *gin.Engine, modules []shared_ports.ModulePort) {
	// User Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "user-service",
			"version": "1.0.0",
		})
	})

	// Add middlewares
	router.Use(middlewares.StructuredLogger())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.RequestIDMiddleware())
	router.Use(middlewares.ErrorHandlerMiddleware())
	router.Use(middlewares.SecurityHeadersMiddleware())

	for _, item := range modules {
		switch mod := item.(type) {
		case *module.UserModule:
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
		log.Printf("🚀 User service starting on port %s", config.Env.ApiPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down user service...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("✅ User service stopped gracefully")
}
