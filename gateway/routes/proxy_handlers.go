package routes

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/Akiles94/go-test-api/shared/infra/grpc/gen/registry"
	"github.com/gin-gonic/gin"
)

func (dr *DynamicRouter) createSharedProxyHandler(routeKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get service info for this specific route
		dr.mutex.RLock()
		service, exists := dr.routeToService[routeKey]
		dr.mutex.RUnlock()

		if !exists {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "route no longer available",
				"route": routeKey,
			})
			return
		}

		// Parse service URL
		serviceURL, err := url.Parse(service.Url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "invalid service configuration",
			})
			return
		}

		// Add service information headers
		c.Header("X-Service-Name", service.Name)
		c.Header("X-Service-Version", service.Version)

		// Create a custom director for this request
		dr.sharedProxy.Director = func(req *http.Request) {
			req.URL.Scheme = serviceURL.Scheme
			req.URL.Host = serviceURL.Host
			req.URL.Path = strings.TrimPrefix(req.URL.Path, "/api/v1")
			if req.URL.Path == "" {
				req.URL.Path = "/"
			}

			// Preserve query parameters
			if serviceURL.RawQuery == "" || req.URL.RawQuery == "" {
				req.URL.RawQuery = serviceURL.RawQuery + req.URL.RawQuery
			} else {
				req.URL.RawQuery = serviceURL.RawQuery + "&" + req.URL.RawQuery
			}
		}

		// Proxy the request to the service
		dr.sharedProxy.ServeHTTP(c.Writer, c.Request)
	}
}

func (dr *DynamicRouter) createHealthProxyHandler(service *registry.ServiceInfo) gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceURL, err := url.Parse(service.Url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "invalid service configuration",
			})
			return
		}

		c.Header("X-Service-Name", service.Name)
		c.Header("X-Service-Version", service.Version)
		c.Header("X-Health-Check", "true")

		dr.sharedProxy.Director = func(req *http.Request) {
			req.URL.Scheme = serviceURL.Scheme
			req.URL.Host = serviceURL.Host
			req.URL.Path = service.HealthEndpoint
			req.URL.RawQuery = ""
		}

		dr.sharedProxy.ServeHTTP(c.Writer, c.Request)
	}
}
