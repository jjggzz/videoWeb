package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/jjggzz/kj/discovery"
	"github.com/openzipkin/zipkin-go"
	"videoWeb/common/ecode"
	"videoWeb/customer-service/config"
	"videoWeb/customer-service/dao"
	"videoWeb/customer-service/dao/repository"
	genpb "videoWeb/generate-service/proto"
	gengrpc "videoWeb/generate-service/svc/client/grpc"
)

var Cus Service

type service struct {
	dao *dao.Dao
	gen genpb.GenerateServer
}

func New(conf *config.Config, dao *dao.Dao, discover discovery.Discover, tracer *zipkin.Tracer, logger log.Logger) Service {
	genIns, _ := discover.Discovery(conf.GenerateServerName)
	generateServer, _ := gengrpc.NewLoadBalanceClient(
		genIns,
		tracer,
		logger,
	)
	return &service{dao: dao, gen: generateServer}
}

type Service interface {
	RegisterByPhone(ctx context.Context, phone string) (ecode.ECode, error)
	LoginByPhone(ctx context.Context, phone string) (ecode.ECode, string, error)
	GetCustomerInfoByToken(ctx context.Context, token string) (ecode.ECode, *repository.Customer, error)
}
