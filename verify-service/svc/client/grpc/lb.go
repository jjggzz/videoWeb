package grpc

import (
	"github.com/go-kit/kit/log"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/loadbalnace"
	"github.com/openzipkin/zipkin-go"
	"time"
	pb "videoWeb/verify-service/proto"
	"videoWeb/verify-service/svc"
)

func NewVerifyBalanceClient(dis discovery.Discover, serverName string, tracer *zipkin.Tracer, logger log.Logger) (pb.VerifyServer, error) {
	// 服务发现
	instancer, err := dis.Discovery(serverName)
	if err != nil {
		return nil, err
	}

	options := []grpctransport.ClientOption{
		kitzipkin.GRPCClientTrace(tracer),
	}

	var endpoints svc.Endpoints
	// 获取发送验证码端点
	{
		factory := factoryGRPCFor(svc.MakeSendVerifyCodeEndpoint, options...)
		endpoints.SendVerifyCodeEndpoint = loadbalnace.BuildLoadBalance(instancer,
			factory,
			loadbalnace.Random,
			3,
			time.Millisecond*500,
			logger,
		)
	}
	// 获取校验验证码端点
	{
		factory := factoryGRPCFor(svc.MakeCheckVerifyCodeEndpoint, options...)
		endpoints.CheckVerifyCodeEndpoint = loadbalnace.BuildLoadBalance(instancer,
			factory,
			loadbalnace.Random,
			3,
			time.Millisecond*500,
			logger,
		)
	}
	return endpoints, nil
}
