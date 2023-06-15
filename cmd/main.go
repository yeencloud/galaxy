package main

import (
	"github.com/gorpc-experiments/galaxy/src/service"
	"github.com/yeencloud/ServiceCore"
)

func main() {
	ServiceCore.SetupLogging()

	galaxy := service.NewGalaxy()

	ServiceCore.PublishMicroService(galaxy, false)
}