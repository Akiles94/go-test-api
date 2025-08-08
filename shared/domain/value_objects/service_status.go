package value_objects

type ServiceStatus string

const (
	ServiceStatusHealthy     ServiceStatus = "healthy"
	ServiceStatusUnhealthy   ServiceStatus = "unhealthy"
	ServiceStatusStarting    ServiceStatus = "starting"
	ServiceStatusStopping    ServiceStatus = "stopping"
	ServiceStatusMaintenance ServiceStatus = "maintenance"
	ServiceStatusUnknown     ServiceStatus = "unknown"
)

func (s ServiceStatus) IsHealthy() bool {
	return s == ServiceStatusHealthy
}

func (s ServiceStatus) IsAvailable() bool {
	return s == ServiceStatusHealthy || s == ServiceStatusStarting
}

func (s ServiceStatus) String() string {
	return string(s)
}

func (s ServiceStatus) IsValid() bool {
	switch s {
	case ServiceStatusHealthy, ServiceStatusUnhealthy, ServiceStatusStarting,
		ServiceStatusStopping, ServiceStatusMaintenance, ServiceStatusUnknown:
		return true
	default:
		return false
	}
}

func NewServiceStatus(status string) ServiceStatus {
	s := ServiceStatus(status)
	if s.IsValid() {
		return s
	}
	return ServiceStatusUnknown
}
