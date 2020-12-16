package http

import "videoWeb/portal/service"

type Http struct {
	srv service.Service
}

func New(service service.Service) *Http {
	return &Http{srv: service}
}
