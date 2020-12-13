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
	"videoWeb/verify-service/config"
	"videoWeb/verify-service/dao"
	"videoWeb/verify-service/handlers"
	"videoWeb/verify-service/proto"
	"videoWeb/verify-service/service"
	"videoWeb/verify-service/svc"
	"videoWeb/verify-service/svc/server"
)

func main() {
	// 读取配置
	flag.Parse()
	// 初始化配置
	config.Init()
	// 日志初始化
	logger := log.BuildLogger(config.Conf.Server.ServerName, os.Stderr)

	// 初始化dao
	var d *dao.Dao
	{
		d = dao.New(config.Conf)
	}
	// 初始化service
	{
		service.Ver = service.New(d)
	}

	// 初始化端点
	handle := handlers.NewService()
	// 熔断、限流等中间件
	endpoints := handlers.WrapEndpoints(server.NewEndpoints(handle))

	// 监听停止操作ctrl + c 等等
	errs := make(chan error)
	go handlers.InterruptHandler(errs)

	// 服务注册组件初始化
	consulDiscovery := discovery.NewConsulDiscovery(
		config.Conf.Discovery.Consul.Address,
		config.Conf.Server.ServerName,
		config.Conf.Server.Tcp.Port,
		logger,
	)

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

		// 链路追踪
		tracer, err := track.BuildZipkinTracer(config.Conf.Zipkin.Address, config.Conf.Server.ServerName)
		options := []grpctransport.ServerOption{
			grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
			zipkin.GRPCServerTrace(tracer),
		}
		grpcServer := svc.MakeGRPCServer(endpoints, options...)
		s := grpc.NewServer()
		proto.RegisterVerifyServer(s, grpcServer)
		errs <- s.Serve(ln)
	}()

	// 退出注销服务
	_ = level.Error(logger).Log("exit", <-errs)
	consulDiscovery.DeregisterServer()

}
