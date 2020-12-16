package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log/level"
	"github.com/jjggzz/kj/log"
	"github.com/jjggzz/kj/uitls"
	"os"
	"os/signal"
	"syscall"
	"videoWeb/portal/config"
	"videoWeb/portal/http"
	"videoWeb/portal/service"
)

// 入口
func main() {
	flag.Parse()
	config.Init()
	// 日志初始化
	logger := log.BuildLogger(config.Conf.Server.ServerName, os.Stderr)

	// service初始化
	var srv service.Service
	{
		srv = service.New(config.Conf)
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
		errs <- engine.Run(fmt.Sprintf("%s:%d", uitls.LocalIpv4(), config.Conf.Server.Http.Port))
	}()

	_ = level.Error(logger).Log("exit", <-errs)

}
