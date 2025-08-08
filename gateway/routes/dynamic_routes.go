package routes

import (
	"context"
	"fmt"
	"log"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
	"github.com/Akiles94/go-test-api/shared/infra/middlewares"
	"github.com/gin-gonic/gin"
)

type DynamicRouter struct {
	registry    shared_ports.ServiceRegistryPort
	authService shared_ports.AuthServicePort
	router      *gin.Engine
	services    map[string]value_objects.ServiceInfo
}

func NewDynamicRouter(registry shared_ports.ServiceRegistryPort, authService shared_ports.AuthServicePort) *DynamicRouter {
	return &DynamicRouter{
		registry:    registry,
		authService: authService,
		services:    make(map[string]value_objects.ServiceInfo),
	}
}

func (dr *DynamicRouter) SetupRoutes(router *gin.Engine) {
	dr.router = router

	// Load initial services
	dr.loadServices()

	// Watch for changes
	go dr.watchServices()
}

func (dr *DynamicRouter) loadServices() {
	services, err := dr.registry.GetServices(context.Background())
	if err != nil {
		log.Printf("❌ Failed to load services: %v", err)
		return
	}

	api := dr.router.Group("/api/v1")

	for _, service := range services {
		dr.registerServiceRoutes(api, service)
		dr.services[service.Name] = service
	}

	log.Printf("✅ Loaded %d services", len(services))
}

func (dr *DynamicRouter) registerServiceRoutes(api *gin.RouterGroup, service value_objects.ServiceInfo) {
	for _, route := range service.Routes {
		// Create route path
		fullPath := route.Path

		log.Printf("📡 Registering route: %s %s -> %s", route.Method, fullPath, service.URL)

		// Create handler with middleware chain
		handler := dr.createHandler(service)

		// Add middleware based on route definition
		middlewares := dr.buildMiddlewares(route)

		// Register route
		switch strings.ToUpper(route.Method) {
		case "GET":
			api.GET(fullPath, middlewares...).GET(fullPath, handler)
		case "POST":
			api.POST(fullPath, middlewares...).POST(fullPath, handler)
		case "PUT":
			api.PUT(fullPath, middlewares...).PUT(fullPath, handler)
		case "DELETE":
			api.DELETE(fullPath, middlewares...).DELETE(fullPath, handler)
		case "PATCH":
			api.PATCH(fullPath, middlewares...).PATCH(fullPath, handler)
		}
	}
}

func (dr *DynamicRouter) createHandler(service value_objects.ServiceInfo) gin.HandlerFunc {
	targetURL, _ := url.Parse(service.URL)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return func(c *gin.Context) {
		// Forward user context headers
		if userID, exists := c.Get("user_id"); exists {
			c.Request.Header.Set("X-User-ID", fmt.Sprintf("%v", userID))
		}
		if userEmail, exists := c.Get("user_email"); exists {
			c.Request.Header.Set("X-User-Email", fmt.Sprintf("%v", userEmail))
		}

		// Add service info headers
		c.Request.Header.Set("X-Service-Name", service.Name)
		c.Request.Header.Set("X-Service-Version", service.Version)

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func (dr *DynamicRouter) buildMiddlewares(route value_objects.RouteDefinition) []gin.HandlerFunc {
	var middlewareChain []gin.HandlerFunc

	// Auth middleware
	if route.Protected {
		middlewareChain = append(middlewareChain, middlewares.AuthMiddleware(dr.authService))
	}

	// Rate limiting middleware
	if route.RateLimit > 0 {
		middlewareChain = append(middlewareChain, middlewares.RateLimitMiddleware(route.RateLimit))
	}

	// Timeout middleware
	// if route.Timeout > 0 {
	//     middlewareChain = append(middlewareChain, middlewares.TimeoutMiddleware(time.Duration(route.Timeout)*time.Second))
	// }

	return middlewareChain
}

func (dr *DynamicRouter) watchServices() {
	watch, err := dr.registry.Watch(context.Background())
	if err != nil {
		log.Printf("❌ Failed to watch services: %v", err)
		return
	}

	for services := range watch {
		log.Println("🔄 Services changed, reloading routes...")
		dr.reloadRoutes(services)
	}
}

func (dr *DynamicRouter) reloadRoutes(services []value_objects.ServiceInfo) {
	// Clear existing routes (you'd need to implement this)
	// For now, just log the change
	for _, service := range services {
		if existing, exists := dr.services[service.Name]; exists {
			if existing.Version != service.Version {
				log.Printf("🔄 Service %s updated: %s -> %s", service.Name, existing.Version, service.Version)
			}
		} else {
			log.Printf("🆕 New service discovered: %s", service.Name)
		}
		dr.services[service.Name] = service
	}
}
