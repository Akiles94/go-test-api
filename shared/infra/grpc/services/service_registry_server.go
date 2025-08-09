package grpc_services

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Akiles94/go-test-api/shared/infra/grpc/gen/registry"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServiceRegistryServer struct {
	registry.UnimplementedServiceRegistryServer
	services map[string]*registry.ServiceInfo
	mutex    sync.RWMutex
	watchers []chan *registry.ServiceUpdate
}

func NewServiceRegistryServer() *ServiceRegistryServer {
	return &ServiceRegistryServer{
		services: make(map[string]*registry.ServiceInfo),
		watchers: make([]chan *registry.ServiceUpdate, 0),
	}
}

func (srs *ServiceRegistryServer) RegisterService(ctx context.Context, req *registry.RegisterServiceRequest) (*registry.RegisterServiceResponse, error) {
	log.Printf("📝 Registering service: %s", req.Service.Name)

	srs.mutex.Lock()
	defer srs.mutex.Unlock()

	if req.Service.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "service name is required")
	}
	if req.Service.Version == "" {
		return nil, status.Error(codes.InvalidArgument, "service version is required")
	}
	if req.Service.Url == "" {
		return nil, status.Error(codes.InvalidArgument, "service url is required")
	}
	if len(req.Service.Routes) == 0 {
		return nil, status.Error(codes.InvalidArgument, "service routes are required")
	}

	// Check if service already exists
	serviceName := req.Service.Name
	existingService, exists := srs.services[serviceName]
	eventType := registry.ServiceUpdateType_SERVICE_UPDATE_TYPE_ADDED

	if exists {
		// Compare versions to detect updates
		if existingService.Version != req.Service.Version {
			eventType = registry.ServiceUpdateType_SERVICE_UPDATE_TYPE_UPDATED
			log.Printf("🔄 Service %s updated: %s -> %s", serviceName, existingService.Version, req.Service.Version)
		} else {
			log.Printf("📝 Service %s re-registered (same version)", serviceName)
			return &registry.RegisterServiceResponse{Success: true}, nil
		}
	}

	// Store/update service
	req.Service.RegisteredAt = timestamppb.New(time.Now())
	req.Service.Status = registry.ServiceStatus_SERVICE_STATUS_HEALTHY
	srs.services[serviceName] = req.Service

	// Notify watchers
	srs.notifyWatchers(eventType, req.Service)

	log.Printf("✅ Service registered: %s at %s", req.Service.Name, req.Service.Url)

	return &registry.RegisterServiceResponse{
		Success: true,
		Message: "service registered successfully",
	}, nil
}

func (srs *ServiceRegistryServer) GetServices(ctx context.Context, req *registry.GetServicesRequest) (*registry.GetServicesResponse, error) {
	srs.mutex.RLock()
	defer srs.mutex.RUnlock()

	services := make([]*registry.ServiceInfo, 0, len(srs.services))
	for _, service := range srs.services {
		services = append(services, service)
	}

	log.Printf("📋 Returning %d registered services", len(services))

	return &registry.GetServicesResponse{
		Services: services,
	}, nil
}

func (srs *ServiceRegistryServer) WatchServices(req *registry.WatchServicesRequest, stream registry.ServiceRegistry_WatchServicesServer) error {
	log.Println("👀 New service watcher connected")

	watcher := make(chan *registry.ServiceUpdate, 10)

	srs.mutex.Lock()
	srs.watchers = append(srs.watchers, watcher)
	srs.mutex.Unlock()

	// Send current services
	srs.mutex.RLock()
	for _, service := range srs.services {
		update := &registry.ServiceUpdate{
			EventType: registry.ServiceUpdateType_SERVICE_UPDATE_TYPE_ADDED,
			Service:   service,
			Timestamp: timestamppb.New(time.Now()),
		}
		if err := stream.Send(update); err != nil {
			srs.mutex.RUnlock()
			return err
		}
	}
	srs.mutex.RUnlock()

	// Listen for updates
	for {
		select {
		case update := <-watcher:
			if err := stream.Send(update); err != nil {
				return err
			}
		case <-stream.Context().Done():
			log.Println("👋 Service watcher disconnected")
			return stream.Context().Err()
		}
	}
}

func (srs *ServiceRegistryServer) notifyWatchers(eventType registry.ServiceUpdateType, service *registry.ServiceInfo) {
	update := &registry.ServiceUpdate{
		EventType: eventType,
		Service:   service,
		Timestamp: timestamppb.New(time.Now()),
	}

	for _, watcher := range srs.watchers {
		select {
		case watcher <- update:
		default:
			// Watcher buffer full, skip
		}
	}
}
