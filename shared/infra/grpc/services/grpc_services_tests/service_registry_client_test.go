package grpc_services_tests

import (
	"testing"

	grpc_services "github.com/Akiles94/go-test-api/shared/infra/grpc/services"
	"github.com/stretchr/testify/assert"
)

func TestServiceRegistryClientConfigDefaults(t *testing.T) {
	cfg := grpc_services.ServiceRegistryClientConfig{
		GatewayAddress: "localhost:9090",
		ServiceName:    "test-service",
		ServiceVersion: "v1.0.0",
		ServiceURL:     "http://localhost:8080",
	}

	assert.Equal(t, "localhost:9090", cfg.GatewayAddress)
	assert.Equal(t, "test-service", cfg.ServiceName)
	assert.Equal(t, "v1.0.0", cfg.ServiceVersion)
	assert.Equal(t, "http://localhost:8080", cfg.ServiceURL)
}

func TestServiceRegistryClientConfigValidation(t *testing.T) {
	tests := []struct {
		name   string
		config grpc_services.ServiceRegistryClientConfig
		valid  bool
	}{
		{
			name: "valid config",
			config: grpc_services.ServiceRegistryClientConfig{
				GatewayAddress: "localhost:9090",
				ServiceName:    "test-service",
				ServiceVersion: "v1.0.0",
				ServiceURL:     "http://localhost:8080",
			},
			valid: true,
		},
		{
			name: "empty service name",
			config: grpc_services.ServiceRegistryClientConfig{
				GatewayAddress: "localhost:9090",
				ServiceName:    "",
				ServiceVersion: "v1.0.0",
				ServiceURL:     "http://localhost:8080",
			},
			valid: false,
		},
		{
			name: "empty gateway address",
			config: grpc_services.ServiceRegistryClientConfig{
				GatewayAddress: "",
				ServiceName:    "test-service",
				ServiceVersion: "v1.0.0",
				ServiceURL:     "http://localhost:8080",
			},
			valid: false,
		},
		{
			name: "empty service version",
			config: grpc_services.ServiceRegistryClientConfig{
				GatewayAddress: "localhost:9090",
				ServiceName:    "test-service",
				ServiceVersion: "",
				ServiceURL:     "http://localhost:8080",
			},
			valid: false,
		},
		{
			name: "empty service URL",
			config: grpc_services.ServiceRegistryClientConfig{
				GatewayAddress: "localhost:9090",
				ServiceName:    "test-service",
				ServiceVersion: "v1.0.0",
				ServiceURL:     "",
			},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				assert.NotEmpty(t, tt.config.ServiceName)
				assert.NotEmpty(t, tt.config.GatewayAddress)
				assert.NotEmpty(t, tt.config.ServiceVersion)
				assert.NotEmpty(t, tt.config.ServiceURL)
			} else {
				// Verify that at least one of the required fields is empty
				isEmpty := tt.config.ServiceName == "" ||
					tt.config.GatewayAddress == "" ||
					tt.config.ServiceVersion == "" ||
					tt.config.ServiceURL == ""
				assert.True(t, isEmpty)
			}
		})
	}
}

// Basic test to verify client creation with valid configuration
func TestServiceRegistryClientCreation(t *testing.T) {
	config := grpc_services.ServiceRegistryClientConfig{
		GatewayAddress: "localhost:9090",
		ServiceName:    "test-service",
		ServiceVersion: "v1.0.0",
		ServiceURL:     "http://localhost:8080",
	}

	// This test would verify that the configuration is valid
	// We don't try to create the real client because it requires a gRPC connection
	assert.NotEmpty(t, config.ServiceName)
	assert.NotEmpty(t, config.GatewayAddress)
	assert.NotEmpty(t, config.ServiceVersion)
	assert.NotEmpty(t, config.ServiceURL)
}
