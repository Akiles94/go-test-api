package grpc_services_tests

import (
	"context"
	"testing"

	"github.com/Akiles94/go-test-api/shared/infra/grpc/gen/registry"
	grpc_services "github.com/Akiles94/go-test-api/shared/infra/grpc/services"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestRegisterAndGetServices(t *testing.T) {
	server := grpc_services.NewServiceRegistryServer()

	// Simulate service registration
	service := &registry.ServiceInfo{
		Name:    "product-service",
		Version: "1.0.0",
		Url:     "http://product-service:8081",
		Routes:  []*registry.RouteInfo{{Path: "/test", Method: "GET"}}, // Add required routes
	}
	req := &registry.RegisterServiceRequest{Service: service}
	_, err := server.RegisterService(context.Background(), req)
	assert.NoError(t, err)

	// Get registered services
	resp, err := server.GetServices(context.Background(), &emptypb.Empty{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, len(resp.Services))
	assert.Equal(t, "product-service", resp.Services[0].Name)
}

func TestUnregisterService(t *testing.T) {
	server := grpc_services.NewServiceRegistryServer()

	service := &registry.ServiceInfo{
		Name:    "user-service",
		Version: "1.0.0",
		Url:     "http://user-service:8082",
		Routes:  []*registry.RouteInfo{{Path: "/test", Method: "GET"}}, // Add required routes
	}
	req := &registry.RegisterServiceRequest{Service: service}
	_, err := server.RegisterService(context.Background(), req)
	assert.NoError(t, err)

	// Deregister (use the correct method)
	deregReq := &registry.DeregisterServiceRequest{ServiceName: service.Name}
	_, err = server.DeregisterService(context.Background(), deregReq)
	assert.NoError(t, err)

	// Verify that it's not registered
	resp, err := server.GetServices(context.Background(), &emptypb.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, 0, len(resp.Services))
}

func TestRegisterServiceWithInvalidData(t *testing.T) {
	server := grpc_services.NewServiceRegistryServer()

	// Try to register service with invalid data
	service := &registry.ServiceInfo{
		Name:    "",
		Version: "1.0.0",
		Url:     "http://invalid-service:8083",
	}
	req := &registry.RegisterServiceRequest{Service: service}
	_, err := server.RegisterService(context.Background(), req)
	assert.Error(t, err)
}

func TestGetServicesWhenEmpty(t *testing.T) {
	server := grpc_services.NewServiceRegistryServer()

	// Get services when none are registered
	resp, err := server.GetServices(context.Background(), &emptypb.Empty{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 0, len(resp.Services))
}

func TestUnregisterNonExistentService(t *testing.T) {
	server := grpc_services.NewServiceRegistryServer()

	// Try to deregister a service that doesn't exist
	deregReq := &registry.DeregisterServiceRequest{
		ServiceName: "non-existent-service",
	}
	_, err := server.DeregisterService(context.Background(), deregReq)
	assert.NoError(t, err) // Should be successful even if it doesn't exist
}
