package value_objects

import "time"

type ServiceInfo struct {
	Name            string            `json:"name"`
	URL             string            `json:"url"`
	Health          string            `json:"health"`
	Routes          []RouteDefinition `json:"routes"`
	Version         string            `json:"version"`
	Status          ServiceStatus     `json:"status"`
	RegisteredAt    time.Time         `json:"registered_at"`
	LastHealthCheck time.Time         `json:"last_health_check,omitempty"`
}

func (si *ServiceInfo) SetStatus(status ServiceStatus) {
	si.Status = status
}

func (si *ServiceInfo) IsHealthy() bool {
	return si.Status.IsHealthy()
}

func (si *ServiceInfo) IsAvailable() bool {
	return si.Status.IsAvailable()
}
