package service

import (
	"github.com/rs/zerolog/log"
	"github.com/yeencloud/ServiceCore/serviceError"
	"github.com/yeencloud/galaxy/src/domain"
)

func (t *Galaxy) LookUp(args domain.LookUpRequest) (domain.LookUpResponse, *serviceError.Error) {
	log.Info().Str("service", args.Service).Str("method", args.Method).Msg("Looking up method...")

	var response domain.LookUpResponse

	service := args.Service
	method := args.Method

	log.Info().Str("service", service).Str("method", method).Msg("Looking up method")

	for _, s := range t.ServiceLibrary {
		if s.Name == service {
			for _, m := range s.Methods {
				if m.Name == method {
					response.Address = m.Instance.Address
					response.Port = m.Instance.Port
					log.Info().Str("service", service).Str("method", args.Method).Str("at", response.Address).Msg("Method found")
					return response, nil
				}
			}
		}
	}

	log.Warn().Str("service", service).Str("method", args.Method).Msg("Method not found")

	return domain.LookUpResponse{}, serviceError.Trace(domain.ErrMethodNotFound)
}