package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/tracing/zipkin"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/log"
	"github.com/jjggzz/kj/track"
	"github.com/jjggzz/kj/uitls"
	"google.golang.org/grpc"
	"net"
	"os"
	"videoWeb/notice-service/config"
	"videoWeb/notice-service/handlers"
	"videoWeb/notice-service/proto"
	"videoWeb/notice-service/service"
	"videoWeb/notice-service/svc"
	"videoWeb/notice-service/svc/server"
)

func main() {
	flag.Parse()
	// 初始化配置(包含了全局配置)
	config.Init()
	// 日志初始化
	logger := log.BuildLogger(config.Conf.Server.ServerName, os.Stderr)
	// 服务注册与发现
	consulDiscovery := discovery.NewConsulDiscovery(
		config.Conf.Discovery.Consul.Address,
		config.Conf.Server.ServerName,
		config.Conf.Server.Tcp.Port,
		logger,
	)
	// 链路追踪
	tracer, err := track.BuildZipkinTracer(config.Conf.Zipkin.Address, config.Conf.Server.ServerName)
	if err != nil {
		_ = level.Error(logger).Log("err", err)
	}

	// service初始化
	{
		service.Not = service.New(config.Conf, logger)
	}

	// 初始化端点
	handle := handlers.NewService()
	// 熔断、限流等中间件
	endpoints := handlers.WrapEndpoints(server.NewEndpoints(handle))

	// 监听停止操作ctrl + c 等等
	errs := make(chan error)
	go handlers.InterruptHandler(errs)

	go func() {
		// 注册
		consulDiscovery.RegisterServer()
		_ = level.Info(logger).Log("transport", "gRPC", "Ip", uitls.LocalIpv4(), "Port", config.Conf.Server.Tcp.Port)

		// 启动服务
		ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", uitls.LocalIpv4(), config.Conf.Server.Tcp.Port))
		if err != nil {
			errs <- err
			return
		}
		options := []grpctransport.ServerOption{
			grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
			zipkin.GRPCServerTrace(tracer),
		}
		grpcServer := svc.MakeGRPCServer(endpoints, options...)
		s := grpc.NewServer()
		proto.RegisterNoticeServer(s, grpcServer)
		errs <- s.Serve(ln)

	}()

	// 退出注销服务
	_ = level.Error(logger).Log("exit", <-errs)
	consulDiscovery.DeregisterServer()
}
