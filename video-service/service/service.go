// author: JGZ
// time:   2021-02-04 16:57
package service

import (
	"context"
	"github.com/jjggzz/kit/log"
	"github.com/jjggzz/kj/discovery"
	"github.com/openzipkin/zipkin-go"
	"videoWeb/common/ecode"
	genpb "videoWeb/generate-service/proto"
	gengrpc "videoWeb/generate-service/svc/client/grpc"
	"videoWeb/video-service/config"
	"videoWeb/video-service/dao"
	"videoWeb/video-service/dao/repository"
)

var Vid Service

type service struct {
	dao *dao.Dao
	gen genpb.GenerateServer
	log log.Logger
}

func New(conf *config.Config, dao *dao.Dao, discover discovery.Discover, tracer *zipkin.Tracer, logger log.Logger) Service {
	genIns, _ := discover.Discovery(conf.GenerateServerName)
	generateServer, _ := gengrpc.NewLoadBalanceClient(
		genIns,
		tracer,
		logger,
	)
	return &service{
		dao: dao,
		gen: generateServer,
		log: logger,
	}
}

type Service interface {
	SaveVideo(ctx context.Context, video *repository.Video) (ecode.ECode, error)
}
