package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jjggzz/kit/log/level"
	"github.com/jjggzz/kj/discovery/nacos"
	"github.com/jjggzz/kj/log"
	"github.com/jjggzz/kj/track"
	"github.com/jjggzz/kj/uitls"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"os/signal"
	"syscall"
	_ "videoWeb/common/ecode/business"
	"videoWeb/portal/config"
	_ "videoWeb/portal/docs"
	"videoWeb/portal/http"
	"videoWeb/portal/service"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 192.168.151.109:8080
func main() {
	flag.Parse()
	config.Init()
	// 日志初始化
	logger := log.BuildLogger(config.Conf.Server.ServerName, os.Stderr)
	// 服务注册与发现
	discovery := nacos.NewNacosDiscovery(
		config.Conf.Discovery.Nacos.Address,
		config.Conf.Server.ServerName,
		config.Conf.Server.Tcp.Port,
		config.Conf.Discovery.Nacos.Namespace,
		config.Conf.Discovery.Nacos.Weight,
		logger,
	)
	//discovery := consul.NewConsulDiscovery(
	//	config.Conf.Discovery.Consul.Address,
	//	config.Conf.Server.ServerName,
	//	config.Conf.Server.Tcp.Port,
	//	logger,
	//)
	// 链路追踪
	tracer, err := track.BuildZipkinTracer(config.Conf.Zipkin.Address, config.Conf.Server.ServerName)
	if err != nil {
		_ = level.Error(logger).Log("err", err)
	}

	// service初始化
	var srv service.Service
	{
		srv = service.New(config.Conf, discovery, tracer, logger)
	}
	// http初始化
	var h *http.Http
	{
		h = http.New(srv)
	}

	// 监听停止操作ctrl + c 等等
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		_ = level.Error(logger).Log("server", "run", "ip", uitls.LocalIpv4(), "port", config.Conf.Server.Http.Port)
		// 启动
		engine := gin.Default()
		http.Middleware(engine)
		http.Router(engine, h)
		url := ginSwagger.URL("http://192.168.151.109:8080/swagger/doc.json")
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		errs <- engine.Run(fmt.Sprintf("%s:%d", uitls.LocalIpv4(), config.Conf.Server.Http.Port))
	}()

	_ = level.Error(logger).Log("exit", <-errs)

}
