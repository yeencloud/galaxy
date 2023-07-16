package main

import (
	"github.com/rs/zerolog/log"
	"github.com/yeencloud/ServiceCore"
	"github.com/yeencloud/galaxy/src/service"
)

func main() {
	galaxy := service.NewGalaxy()

	sh, err := servicecore.NewServiceHost(galaxy, "Galaxy", false)

	if err != nil {
		log.Err(err).Msg("Failed to start service")
		return
	}
	err = sh.Listen()

	if err != nil {
		log.Err(err).Msg("Failed to listen")
	}
}