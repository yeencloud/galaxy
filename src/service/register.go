package service

import (
	"github.com/rs/zerolog/log"
	"github.com/yeencloud/ServiceCore/decompose"
	"github.com/yeencloud/ServiceCore/serviceError"
	"github.com/yeencloud/galaxy/src/domain"
	"time"
)

func (t *Galaxy) serviceExists(module string) bool {
	for _, m := range t.ServiceLibrary {
		if m.Name == module {
			return true
		}
	}

	return false
}

func (t *Galaxy) registerService(module string) {
	if !t.serviceExists(module) {
		t.ServiceLibrary = append(t.ServiceLibrary, domain.Service{
			Name:    module,
			Methods: []domain.Methods{},
		})
	}
}

func (t *Galaxy) registerMethod(service string, method decompose.Method, address string, port int, host string) {
	for serviceIndex, s := range t.ServiceLibrary {
		if s.Name == service {
			methodFound := false
			for methodIndex, m := range s.Methods {
				if m.Name == method.Name {
					t.ServiceLibrary[serviceIndex].Methods[methodIndex].Instance.Address = address
					hostFound := false
					for hostIndex, h := range m.Instance.Hosts {
						if h.Name == host {
							t.ServiceLibrary[serviceIndex].Methods[methodIndex].Instance.Hosts[hostIndex].LastSeen = time.Now().Unix()
							hostFound = true
						}
					}
					if !hostFound {
						t.ServiceLibrary[serviceIndex].Methods[methodIndex].Instance.Hosts = append(t.ServiceLibrary[serviceIndex].Methods[methodIndex].Instance.Hosts, domain.Host{
							Name:     host,
							LastSeen: time.Now().Unix(),
						})
					}
					methodFound = true
					break
				}
			}

			if !methodFound {
				t.ServiceLibrary[serviceIndex].Methods = append(t.ServiceLibrary[serviceIndex].Methods, domain.Methods{
					Name: method.Name,
					Instance: domain.Instance{
						Address: address,
						Port:    port,
						Hosts: []domain.Host{
							{
								Name:     host,
								LastSeen: time.Now().Unix(),
							},
						},
					},
				})
			}
		}
	}
}

func (t *Galaxy) register(args domain.RegisterRequest) {
	address := args.Address
	port := args.Port
	hostname := args.Hostname
	service := args.Components.Name

	if !t.serviceExists(service) {
		t.registerService(service)
	}

	for _, method := range args.Components.Methods {
		t.registerMethod(service, method, address, port, hostname)
	}
}

func (t *Galaxy) Register(args domain.RegisterRequest) (domain.RegisterResponse, *serviceError.Error) {
	componentList := ""
	for i, m := range args.Components.Methods {
		if i > 0 {
			componentList += ", "
		}
		componentList += m.Name
	}

	log.Info().Str("component", componentList).Msg("Registering component")

	t.register(args)

	return domain.RegisterResponse{true}, nil
}