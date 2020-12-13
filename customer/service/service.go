package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/track"
	"videoWeb/customer/config"
	"videoWeb/customer/dao"
	genpb "videoWeb/generate-service/proto"
	gengrpc "videoWeb/generate-service/svc/client/grpc"
	verpb "videoWeb/verify-service/proto"
	"videoWeb/verify-service/svc/client/grpc"
)

type service struct {
	dao *dao.Dao
	gen genpb.GenerateServer
	ver verpb.VerifyServer
}

func New(conf *config.Config, dis discovery.Discover, dao *dao.Dao, logger log.Logger) Service {
	tracer, err := track.BuildZipkinTracer(config.Conf.Zipkin.Address, config.Conf.Server.ServerName)
	generateServer, err := gengrpc.NewGenerateLoadBalanceClient(dis, conf.GenerateServerName, tracer, logger)
	verifyServer, err := grpc.NewVerifyBalanceClient(dis, conf.VerifyServerName, tracer, logger)
	if err != nil {
		_ = logger.Log("err", err)
	}

	return &service{dao: dao, gen: generateServer, ver: verifyServer}

}

type Service interface {
	QueryCustomerByAccessKey(ctx context.Context, accessKey string) (*dao.Customer, error)
	InsertCustomer(ctx context.Context, customer *dao.Customer) error
	GetKey(ctx context.Context) (string, error)
	SendVerify(ctx context.Context, phone string) error
	CheckVerify(ctx context.Context, phone string, code string) (bool, error)
}
