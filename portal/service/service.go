package service

import "videoWeb/portal/config"

type service struct {
}

func New(conf *config.Config) Service {
	return &service{}
}

type Service interface {
}
