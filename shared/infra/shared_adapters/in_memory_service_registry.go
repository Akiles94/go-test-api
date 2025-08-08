package shared_adapters

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
)

type InMemoryServiceRegistry struct {
	services map[string]value_objects.ServiceInfo
	mutex    sync.RWMutex
	watchers []chan []value_objects.ServiceInfo
}

func NewInMemoryServiceRegistry() shared_ports.ServiceRegistryPort {
	return &InMemoryServiceRegistry{
		services: make(map[string]value_objects.ServiceInfo),
		watchers: make([]chan []value_objects.ServiceInfo, 0),
	}
}

func (r *InMemoryServiceRegistry) Register(ctx context.Context, service value_objects.ServiceInfo) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Validate service info
	if service.Name == "" {
		return errors.New("service name is required")
	}
	if service.URL == "" {
		return errors.New("service URL is required")
	}

	// Set registration time
	service.RegisteredAt = time.Now()
	service.Status = value_objects.ServiceStatusHealthy

	// Store service
	r.services[service.Name] = service

	// Notify watchers
	r.notifyWatchers()

	return nil
}

func (r *InMemoryServiceRegistry) Deregister(ctx context.Context, serviceName string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.services[serviceName]; !exists {
		return errors.New("service not found")
	}

	delete(r.services, serviceName)

	// Notify watchers
	r.notifyWatchers()

	return nil
}

func (r *InMemoryServiceRegistry) GetServices(ctx context.Context) ([]value_objects.ServiceInfo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	services := make([]value_objects.ServiceInfo, 0, len(r.services))
	for _, service := range r.services {
		services = append(services, service)
	}

	return services, nil
}

func (r *InMemoryServiceRegistry) GetService(ctx context.Context, name string) (*value_objects.ServiceInfo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	service, exists := r.services[name]
	if !exists {
		return nil, errors.New("service not found")
	}

	return &service, nil
}

func (r *InMemoryServiceRegistry) Watch(ctx context.Context) (<-chan []value_objects.ServiceInfo, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	watcher := make(chan []value_objects.ServiceInfo, 10)
	r.watchers = append(r.watchers, watcher)

	// Send current services immediately
	go func() {
		services, _ := r.GetServices(ctx)
		select {
		case watcher <- services:
		case <-ctx.Done():
			return
		}
	}()

	// Handle context cancellation
	go func() {
		<-ctx.Done()
		r.removeWatcher(watcher)
		close(watcher)
	}()

	return watcher, nil
}

func (r *InMemoryServiceRegistry) notifyWatchers() {
	services := make([]value_objects.ServiceInfo, 0, len(r.services))
	for _, service := range r.services {
		services = append(services, service)
	}

	for _, watcher := range r.watchers {
		select {
		case watcher <- services:
		default:
			// Watcher buffer is full, skip
		}
	}
}

func (r *InMemoryServiceRegistry) removeWatcher(targetWatcher chan []value_objects.ServiceInfo) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, watcher := range r.watchers {
		if watcher == targetWatcher {
			r.watchers = append(r.watchers[:i], r.watchers[i+1:]...)
			break
		}
	}
}

// Health check functionality
func (r *InMemoryServiceRegistry) StartHealthCheck() {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker.C {
			r.performHealthCheck()
		}
	}()
}

func (r *InMemoryServiceRegistry) performHealthCheck() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	log.Println("🔍 Starting health check for all services...")

	for name, service := range r.services {
		go r.checkServiceHealth(name, service)
	}
}

func (r *InMemoryServiceRegistry) checkServiceHealth(serviceName string, service value_objects.ServiceInfo) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Build health check URL
	healthURL := service.URL
	if service.Health != "" {
		healthURL = fmt.Sprintf("%s%s", service.URL, service.Health)
	} else {
		healthURL = fmt.Sprintf("%s/health", service.URL) // Default health endpoint
	}

	// 2. Make HTTP request to health endpoint
	req, err := http.NewRequestWithContext(ctx, "GET", healthURL, nil)
	if err != nil {
		r.updateServiceStatus(serviceName, value_objects.ServiceStatusUnhealthy)
		log.Printf("❌ Health check failed for %s: %v", serviceName, err)
		return
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		r.updateServiceStatus(serviceName, value_objects.ServiceStatusUnhealthy)
		log.Printf("❌ Health check failed for %s: %v", serviceName, err)
		return
	}
	defer resp.Body.Close()

	// 3. Verificar respuesta
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		r.updateServiceStatus(serviceName, value_objects.ServiceStatusHealthy)
		log.Printf("✅ Health check passed for %s", serviceName)
	} else {
		r.updateServiceStatus(serviceName, value_objects.ServiceStatusUnhealthy)
		log.Printf("❌ Health check failed for %s: status code %d", serviceName, resp.StatusCode)
	}
}

func (r *InMemoryServiceRegistry) updateServiceStatus(serviceName string, status value_objects.ServiceStatus) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if service, exists := r.services[serviceName]; exists {
		service.Status = status
		service.LastHealthCheck = time.Now()
		r.services[serviceName] = service

		// Notify watchers of status change
		r.notifyWatchers()
	}
}
