// gateway/routes/dynamic_routes.go
package routes

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"

	"github.com/Akiles94/go-test-api/gateway/config"
	"github.com/Akiles94/go-test-api/shared/infra/grpc/gen/registry"
	grpc_services "github.com/Akiles94/go-test-api/shared/infra/grpc/services"
	"github.com/gin-gonic/gin"
)

type DynamicRouter struct {
	ctx              context.Context
	apiGroup         *gin.RouterGroup
	serviceRegistry  *grpc_services.ServiceRegistryServer
	sharedProxy      *httputil.ReverseProxy
	routeToService   map[string]*registry.ServiceInfo
	registeredRoutes map[string][]string
	mutex            sync.RWMutex
	updateChan       chan *registry.ServiceUpdate
}

// NewDynamicRouter creates a new dynamic router instance
func NewDynamicRouter(ctx context.Context, apiGroup *gin.RouterGroup, serviceRegistry *grpc_services.ServiceRegistryServer) *DynamicRouter {
	// Create a shared proxy with custom director
	sharedProxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// Director will be set per request in the handler
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("Proxy error: %v", err)
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(`{"error": "service unavailable"}`))
		},
	}

	return &DynamicRouter{
		ctx:              ctx,
		apiGroup:         apiGroup,
		serviceRegistry:  serviceRegistry,
		sharedProxy:      sharedProxy,
		routeToService:   make(map[string]*registry.ServiceInfo),
		registeredRoutes: make(map[string][]string),
		updateChan:       make(chan *registry.ServiceUpdate, 100),
	}
}

// handleServiceUpdate processes service change notifications
func (dr *DynamicRouter) handleServiceUpdate(update *registry.ServiceUpdate) {
	dr.mutex.Lock()
	defer dr.mutex.Unlock()

	switch update.EventType {
	case registry.ServiceUpdateType_SERVICE_UPDATE_TYPE_ADDED:
		log.Printf("Adding service: %s", update.Service.Name)
		dr.addServiceRoutes(update.Service)

	case registry.ServiceUpdateType_SERVICE_UPDATE_TYPE_REMOVED:
		log.Printf("Removing service: %s", update.Service.Name)
		dr.removeServiceRoutes(update.Service.Name)

	case registry.ServiceUpdateType_SERVICE_UPDATE_TYPE_UPDATED:
		log.Printf("Updating service: %s", update.Service.Name)
		dr.updateServiceRoutes(update.Service)
	default:
		log.Printf("⚠️ Unknown update event type: %v for service %s", update.EventType, update.Service.Name)
	}
}

// addServiceRoutes dynamically adds routes for a service
func (dr *DynamicRouter) addServiceRoutes(service *registry.ServiceInfo) {
	// Parse service URL
	_, err := url.Parse(service.Url)
	if err != nil {
		log.Printf("Invalid service URL for %s: %v", service.Name, err)
		return
	}

	// Use existing API group
	var addedRoutes []string

	// Health endpoint
	if service.HealthEndpoint != "" {
		healthRouteKey := "GET " + "/health/" + service.Name
		if !dr.routeExists(healthRouteKey) {
			dr.routeToService[healthRouteKey] = service
			handler := dr.createHealthProxyHandler(service)

			dr.apiGroup.GET("/health/"+service.Name, handler)
			addedRoutes = append(addedRoutes, healthRouteKey)
			log.Printf("   Added health route: GET /api/v1/health/%s -> %s%s",
				service.Name, service.Url, service.HealthEndpoint)
		}
	}

	// Register specific routes if defined
	for _, route := range service.Routes {

		// Check for route conflicts BEFORE adding
		routeKey := route.Method + " " + route.Path
		if dr.routeExists(routeKey) {
			log.Printf("⚠️ Route conflict detected: %s (skipping)", routeKey)
			continue
		}

		dr.routeToService[routeKey] = service
		handler := dr.createSharedProxyHandler(routeKey)

		// Add middleware if needed
		if route.Protected {
			handler = dr.authMiddleware(handler)
		}
		if route.RateLimit > 0 {
			handler = dr.rateLimitMiddleware(route.RateLimit, handler)
		} else {
			handler = dr.rateLimitMiddleware(int32(config.Env.RateLimitCount), handler)
		}

		// Register route based on HTTP method
		switch strings.ToUpper(route.Method) {
		case "GET":
			dr.apiGroup.GET(route.Path, handler)
		case "POST":
			dr.apiGroup.POST(route.Path, handler)
		case "PUT":
			dr.apiGroup.PUT(route.Path, handler)
		case "DELETE":
			dr.apiGroup.DELETE(route.Path, handler)
		case "PATCH":
			dr.apiGroup.PATCH(route.Path, handler)
		default:
			dr.apiGroup.Any(route.Path, handler)
		}

		addedRoutes = append(addedRoutes, routeKey)
		log.Printf("   Added route: %s %s -> %s", route.Method, route.Path, service.Url)
	}

	// Track registered routes for later removal
	dr.registeredRoutes[service.Name] = addedRoutes

	log.Printf("Service %s registered with %d routes", service.Name, len(addedRoutes))
}

// Helper method to check route conflicts
func (dr *DynamicRouter) routeExists(routeKey string) bool {
	for _, routes := range dr.registeredRoutes {
		for _, route := range routes {
			if route == routeKey {
				return true
			}
		}
	}
	return false
}

// removeServiceRoutes removes all routes for a service
func (dr *DynamicRouter) removeServiceRoutes(serviceName string) {
	// Remove route mappings for this service
	routes, exists := dr.registeredRoutes[serviceName]
	if exists {
		delete(dr.registeredRoutes, serviceName)
		log.Printf("Service %s removed (%d routes)", serviceName, len(routes))
	}
}

// updateServiceRoutes updates routes for a service
func (dr *DynamicRouter) updateServiceRoutes(service *registry.ServiceInfo) {
	// Remove old routes and add new ones
	dr.removeServiceRoutes(service.Name)
	dr.addServiceRoutes(service)
}

// authMiddleware provides simple authentication middleware
func (dr *DynamicRouter) authMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "authorization required"})
			c.Abort()
			return
		}
		next(c)
	}
}

// rateLimitMiddleware provides simple rate limiting middleware
func (dr *DynamicRouter) rateLimitMiddleware(limit int32, next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Simple rate limiting implementation (add proper rate limiting later)
		c.Header("X-Rate-Limit", string(rune(limit)))
		next(c)
	}
}
