package main

import (
	"log"

	"github.com/Akiles94/go-test-api/config"
	"github.com/Akiles94/go-test-api/db"
	"github.com/Akiles94/go-test-api/domain/models"
	"github.com/Akiles94/go-test-api/infra/modules"
	"github.com/Akiles94/go-test-api/interfaces"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database := db.Connect()

	if err := database.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("❌ DB migration failed: %v", err)
	}

	var appModules []interfaces.Module

	productModule := modules.NewProductModule(database)
	appModules = append(appModules, productModule)

	router := gin.Default()
	api := router.Group("/api")

	for _, m := range appModules {
		switch mod := m.(type) {
		case *modules.ProductModule:
			mod.RegisterRoutes(api.Group("/products"))
		}
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("❌ Error starting server: %v", err)
	}
}
