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
	vidpb "videoWeb/video-service/proto"
	vidgrpc "videoWeb/video-service/svc/client/grpc"
)

type service struct {
	cus cuspb.CustomerServer
	ver verpb.VerifyServer
	vid vidpb.VideoServer
}

func New(conf *config.Config, discover discovery.Discover, tracer *zipkin.Tracer, logger log.Logger) Service {
	cusIns, _ := discover.Discovery(conf.CustomerServerName)
	verIns, _ := discover.Discovery(conf.VerifyServerName)
	vidIns, _ := discover.Discovery(conf.VideoServerName)
	// 客户服务
	customerServer, _ := cusgrpc.NewLoadBalanceClient(
		cusIns,
		tracer,
		logger,
	)
	// 验证码服务
	verifyServer, _ := vergrpc.NewLoadBalanceClient(
		verIns,
		tracer,
		logger,
	)
	// 视频服务
	videoServer, _ := vidgrpc.NewLoadBalanceClient(
		vidIns,
		tracer,
		logger,
	)
	return &service{cus: customerServer, ver: verifyServer, vid: videoServer}
}

type Service interface {
	Login(ctx context.Context, phone string, verify string) (ecode.ECode, string, error)
	Register(ctx context.Context, phone string, verify string) (ecode.ECode, error)
	GetUserInfo(ctx context.Context, token string) (ecode.ECode, *dto.UserInfoDto, error)
	UploadVideo(ctx context.Context, videoDto *dto.UploadVideoDto) (ecode.ECode, error)
	SendVerify(ctx context.Context, strategyName string, phone string) (ecode.ECode, error)
}
