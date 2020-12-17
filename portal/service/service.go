package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/jjggzz/kj/discovery"
	"github.com/openzipkin/zipkin-go"
	cuspb "videoWeb/customer-service/proto"
	cusgrpc "videoWeb/customer-service/svc/client/grpc"
	"videoWeb/portal/config"
)

type service struct {
	cus cuspb.CustomerServer
}

func New(conf *config.Config, discover discovery.Discover, tracer *zipkin.Tracer, logger log.Logger) Service {
	server, _ := cusgrpc.NewCustomerLoadBalanceClient(
		discover,
		conf.CustomerServerName,
		tracer,
		logger,
	)

	return &service{cus: server}
}

type Service interface {
	Login(ctx context.Context, phone string) (string, error)
}
