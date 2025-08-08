package shared_ports

import (
	"context"

	"github.com/Akiles94/go-test-api/shared/domain/value_objects"
)

type ServiceRegistryPort interface {
	Register(ctx context.Context, service value_objects.ServiceInfo) error
	Deregister(ctx context.Context, serviceName string) error
	GetServices(ctx context.Context) ([]value_objects.ServiceInfo, error)
	GetService(ctx context.Context, name string) (*value_objects.ServiceInfo, error)
	Watch(ctx context.Context) (<-chan []value_objects.ServiceInfo, error)
}
