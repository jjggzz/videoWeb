package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/jjggzz/kj/discovery"
	"github.com/openzipkin/zipkin-go"
	"videoWeb/common/ecode"
	notpb "videoWeb/notice-service/proto"
	notgrpc "videoWeb/notice-service/svc/client/grpc"
	"videoWeb/verify-service/config"
	"videoWeb/verify-service/dao"
)

var Ver Service

type service struct {
	dao *dao.Dao
	not notpb.NoticeServer
}

func New(conf *config.Config, dao *dao.Dao, discover discovery.Discover, tracer *zipkin.Tracer, logger log.Logger) Service {
	notIns, _ := discover.Discovery(conf.NoticeService)
	noticeServer, _ := notgrpc.NewLoadBalanceClient(
		notIns,
		tracer,
		logger,
	)
	return &service{dao: dao, not: noticeServer}
}

type Service interface {
	SendPhoneVerify(ctx context.Context, target string) (ecode.ECode, error)
	SendEmailVerify(ctx context.Context, target string) (ecode.ECode, error)
	CheckVerify(ctx context.Context, target string, code string) (ecode.ECode, error)
}
