package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/log"
	"github.com/jjggzz/kj/uitls"
	"net"
	"os"
	"os/signal"
	"syscall"
	"videoWeb/customer/config"
	"videoWeb/customer/dao"
	"videoWeb/customer/http"
	"videoWeb/customer/service"
)

func main() {
	flag.Parse()
	// 初始化配置(包含了全局配置)
	config.Init()

	logger := log.BuildLogger("customer", os.Stderr)

	_ = level.Info(logger).Log("msg", "server started")
	defer func() {
		_ = level.Info(logger).Log("msg", "server ended")
	}()

	consulDiscovery := discovery.NewConsulDiscovery(config.Conf.Discovery.Consul.Address,
		config.Conf.Server.ServerName,
		config.Conf.Server.Tcp.Port,
		logger,
	)

	// dao
	var d *dao.Dao
	{
		d = dao.New(config.Conf)
	}

	// 业务层
	var s service.Service
	{
		s = service.New(config.Conf, consulDiscovery, d, logger)
	}

	// http服务
	var server *http.Server
	{
		server = http.New(s)
	}

	// 监听停止操作ctrl + c 等等
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", uitls.LocalIpv4(), config.Conf.Server.Tcp.Port))
		if err != nil {
			_ = logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		for {
			_, _ = listener.Accept()
		}
	}()

	// 启动http服务
	go func() {
		_ = level.Info(logger).Log("customer", "HTTP", "addr", config.Conf.Server.Http.Port)
		consulDiscovery.RegisterServer()
		server.Router()
		// 启动服务
		errs <- server.Start(config.Conf)
	}()

	err := <-errs
	consulDiscovery.DeregisterServer()
	_ = level.Error(logger).Log("exit", err)

}
