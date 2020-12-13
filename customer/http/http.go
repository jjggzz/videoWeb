package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jjggzz/kj/uitls"
	"videoWeb/customer/config"
	"videoWeb/customer/service"
)

type Server struct {
	service service.Service
	engine  *gin.Engine
}

func New(service service.Service) *Server {
	return &Server{
		service: service,
		engine:  gin.Default(),
	}
}

func (server *Server) Router() {
	server.engine.GET("/getCustomerByAccessKey", server.GetCustomerByAccessKey)
	server.engine.GET("/insertCustomer", server.InsertCustomer)
	server.engine.GET("/getKey", server.GetKey)
	server.engine.GET("/send/:phone", server.Send)
	server.engine.GET("/check/:phone/:code", server.Check)
}

func (server *Server) Start(conf *config.Config) error {
	return server.engine.Run(fmt.Sprintf("%s:%d", uitls.LocalIpv4(), conf.Server.Http.Port))
}
