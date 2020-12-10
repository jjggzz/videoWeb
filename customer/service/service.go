package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/track"
	"videoWeb/customer/config"
	"videoWeb/customer/dao"
	"videoWeb/generate-service/svc/client/grpc"
	pb "videoWeb/proto"
)

type service struct {
	dao *dao.Dao
	gen pb.GenerateServer
}

func New(conf *config.Config, dis discovery.Discover, dao *dao.Dao, logger log.Logger) Service {
	tracer, err := track.BuildZipkinTracer(config.Conf.Zipkin.Address, config.Conf.Server.ServerName)
	server, err := grpc.NewGenerateLoadBalanceClient(dis, conf.GenerateKeyServerName, tracer, logger)
	if err != nil {
		_ = logger.Log("err", err)
	}

	return &service{dao: dao, gen: server}

}

type Service interface {
	QueryCustomerByAccessKey(ctx context.Context, accessKey string) (*dao.Customer, error)
	InsertCustomer(ctx context.Context, customer *dao.Customer) error
	GetKey(ctx context.Context) (string, error)
}
