package routes

import (
	"log"

	"github.com/Akiles94/go-test-api/shared/infra/grpc/gen/registry"
)

// WatchServices listens for real-time service changes using direct notifications
func (dr *DynamicRouter) WatchServices() {
	log.Println("Starting real-time service watcher...")

	dr.serviceRegistry.RegisterUpdateListener(dr.updateChan)

	// Load initial services
	dr.loadInitialServices()

	// Listen for updates
	for update := range dr.updateChan {
		if update == nil {
			continue
		}
		log.Printf("📢 Service update: %s -> %s", update.EventType, update.Service.Name)
		dr.handleServiceUpdate(update)
	}
}

func (dr *DynamicRouter) loadInitialServices() {
	services, err := dr.serviceRegistry.GetServices(dr.ctx, nil)
	if err != nil {
		log.Printf("Failed to load initial services: %v", err)
		return
	}

	dr.mutex.Lock()
	defer dr.mutex.Unlock()

	for _, service := range services.Services {
		if service != nil && service.Name != "" && service.Status == registry.ServiceStatus_SERVICE_STATUS_HEALTHY {
			dr.addServiceRoutes(service)
		}
	}
}

func (dr *DynamicRouter) Close() {
	if dr.updateChan != nil {
		dr.serviceRegistry.UnregisterUpdateListener(dr.updateChan)
		close(dr.updateChan)
	}
}
