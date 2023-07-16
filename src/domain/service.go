package domain

// A Host represents the pod that is running the service
type Host struct {
	Name     string
	LastSeen int64
}

// An Instance represents a running service, it can have multiple hosts (pods) that are running the same service
type Instance struct {
	Address string
	Port    int
	Hosts   []Host
}

// A Methods represents a function of a module that can be called remotely like creating a user, saving a specific data, etc
type Methods struct {
	Name     string
	Instance Instance
}

// A Service represents a collection of services, it's a major component of the system like Authentication, Notifications, Storage, etc...
// it can have multiple services
type Service struct {
	Name    string
	Methods []Methods
}