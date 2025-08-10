package grpc_services

import (
	"context"
	"fmt"
	"log"

	"github.com/Akiles94/go-test-api/shared/application/shared_ports"
	"github.com/Akiles94/go-test-api/shared/infra/grpc/gen/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceRegistry struct {
	context     context.Context
	serviceName string
	serviceURL  string
	modules     []shared_ports.ModulePort
	client      registry.ServiceRegistryClient
	conn        *grpc.ClientConn
}

type ServiceRegistryConfig struct {
	Context        context.Context
	ServiceName    string
	ServiceVersion string
	ServiceURL     string
	HealthEndpoint string
	GatewayAddress string
	Modules        []shared_ports.ModulePort
}

func NewServiceRegistry(config ServiceRegistryConfig) (*ServiceRegistry, error) {
	conn, err := grpc.NewClient(config.GatewayAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gateway: %w", err)
	}

	client := registry.NewServiceRegistryClient(conn)

	return &ServiceRegistry{
		context:     config.Context,
		serviceName: config.ServiceName,
		serviceURL:  config.ServiceURL,
		modules:     config.Modules,
		client:      client,
		conn:        conn,
	}, nil
}
func (sr *ServiceRegistry) RegisterWithGateway(config ServiceRegistryConfig) error {
	allRoutes := make([]*registry.RouteInfo, 0)
	for _, module := range sr.modules {
		routes := module.GetRouteDefinitions()
		allRoutes = append(allRoutes, routes...)
	}
	serviceInfo := &registry.ServiceInfo{
		Name:           config.ServiceName,
		Version:        config.ServiceVersion,
		Url:            config.ServiceURL,
		HealthEndpoint: config.HealthEndpoint,
		Routes:         allRoutes,
	}
	req := &registry.RegisterServiceRequest{
		Service: serviceInfo,
	}
	resp, err := sr.client.RegisterService(sr.context, req)
	if err != nil {
		return fmt.Errorf("failed to register service: %w", err)
	}

	if !resp.Success {
		return fmt.Errorf("service registration failed: %s", resp.Message)
	}

	log.Printf("✅ Successfully registered %s with %d routes",
		config.ServiceName, len(allRoutes))

	return nil
}

func (sr *ServiceRegistry) DeregisterFromGateway() error {
	req := &registry.DeregisterServiceRequest{
		ServiceName: sr.serviceName,
	}
	_, err := sr.client.DeregisterService(sr.context, req)
	if err != nil {
		return fmt.Errorf("failed to deregister service: %w", err)
	}
	sr.conn.Close()
	return nil
}
