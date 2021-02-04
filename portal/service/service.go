package service

import (
	"context"
	"github.com/jjggzz/kit/log"
	"github.com/jjggzz/kj/discovery"
	"github.com/openzipkin/zipkin-go"
	"videoWeb/common/ecode"
	cuspb "videoWeb/customer-service/proto"
	cusgrpc "videoWeb/customer-service/svc/client/grpc"
	"videoWeb/portal/config"
	"videoWeb/portal/service/dto"
	verpb "videoWeb/verify-service/proto"
	vergrpc "videoWeb/verify-service/svc/client/grpc"
)

type service struct {
	cus cuspb.CustomerServer
	ver verpb.VerifyServer
}

func New(conf *config.Config, discover discovery.Discover, tracer *zipkin.Tracer, logger log.Logger) Service {
	cusIns, _ := discover.Discovery(conf.CustomerServerName)
	verIns, _ := discover.Discovery(conf.VerifyServerName)
	customerServer, _ := cusgrpc.NewLoadBalanceClient(
		cusIns,
		tracer,
		logger,
	)
	verifyServer, _ := vergrpc.NewLoadBalanceClient(
		verIns,
		tracer,
		logger,
	)
	return &service{cus: customerServer, ver: verifyServer}
}

type Service interface {
	Login(ctx context.Context, phone string, verify string) (ecode.ECode, string, error)
	Register(ctx context.Context, phone string, verify string) (ecode.ECode, error)
	GetUserInfo(ctx context.Context, token string) (ecode.ECode, *dto.UserInfoDto, error)
	SendVerify(ctx context.Context, strategyName string, phone string) (ecode.ECode, error)
}
