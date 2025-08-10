package grpc_services

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Akiles94/go-test-api/shared/infra/grpc/gen/registry"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServiceRegistryServer struct {
	registry.UnimplementedServiceRegistryServer
	services        map[string]*registry.ServiceInfo
	mutex           sync.RWMutex
	updateListeners []chan<- *registry.ServiceUpdate
}

func NewServiceRegistryServer() *ServiceRegistryServer {
	return &ServiceRegistryServer{
		services:        make(map[string]*registry.ServiceInfo),
		updateListeners: make([]chan<- *registry.ServiceUpdate, 0),
	}
}

func (srs *ServiceRegistryServer) RegisterUpdateListener(ch chan<- *registry.ServiceUpdate) {
	srs.mutex.Lock()
	defer srs.mutex.Unlock()

	srs.updateListeners = append(srs.updateListeners, ch)
	log.Printf("📢 Registered new update listener (total: %d)", len(srs.updateListeners))
}

func (srs *ServiceRegistryServer) UnregisterUpdateListener(ch chan<- *registry.ServiceUpdate) {
	srs.mutex.Lock()
	defer srs.mutex.Unlock()

	for i, listener := range srs.updateListeners {
		if listener == ch {
			srs.updateListeners = append(srs.updateListeners[:i], srs.updateListeners[i+1:]...)
			log.Printf("📢 Unregistered update listener (total: %d)", len(srs.updateListeners))
			break
		}
	}
}

func (srs *ServiceRegistryServer) notifyUpdateListeners(update *registry.ServiceUpdate) {
	for i, listener := range srs.updateListeners {
		select {
		case listener <- update:
		default:
			log.Printf("⚠️ Update listener %d channel is full, skipping notification", i)
		}
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

	update := &registry.ServiceUpdate{
		EventType: eventType,
		Service:   req.Service,
		Timestamp: timestamppb.New(time.Now()),
	}
	srs.notifyUpdateListeners(update)

	log.Printf("✅ Service registered: %s at %s", req.Service.Name, req.Service.Url)

	return &registry.RegisterServiceResponse{
		Success: true,
		Message: "service registered successfully",
	}, nil
}

func (srs *ServiceRegistryServer) DeregisterService(ctx context.Context, req *registry.DeregisterServiceRequest) (*registry.DeregisterServiceResponse, error) {
	log.Printf("🗑️ Deregistering service: %s", req.ServiceName)

	srs.mutex.Lock()
	defer srs.mutex.Unlock()

	if req.ServiceName == "" {
		return nil, status.Error(codes.InvalidArgument, "service name is required")
	}

	// Check if service exists
	service, exists := srs.services[req.ServiceName]
	if !exists {
		return &registry.DeregisterServiceResponse{
			Success: false,
			Message: "service not found",
		}, nil
	}

	// Remove service
	delete(srs.services, req.ServiceName)

	update := &registry.ServiceUpdate{
		EventType: registry.ServiceUpdateType_SERVICE_UPDATE_TYPE_REMOVED,
		Service:   service,
		Timestamp: timestamppb.New(time.Now()),
	}
	srs.notifyUpdateListeners(update)

	log.Printf("✅ Service deregistered: %s", req.ServiceName)

	return &registry.DeregisterServiceResponse{
		Success: true,
		Message: "service deregistered successfully",
	}, nil
}

func (srs *ServiceRegistryServer) GetServices(ctx context.Context, _ *emptypb.Empty) (*registry.GetServicesResponse, error) {
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
