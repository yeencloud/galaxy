package service

import "github.com/yeencloud/galaxy/src/domain"

type Galaxy struct {
	ServiceLibrary []domain.Service
}

func NewGalaxy() *Galaxy {
	galaxy := new(Galaxy)

	galaxy.ServiceLibrary = []domain.Service{}

	return galaxy
}