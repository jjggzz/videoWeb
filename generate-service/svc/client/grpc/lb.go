package grpc

import (
	"github.com/go-kit/kit/log"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/loadbalnace"
	"github.com/openzipkin/zipkin-go"
	"time"
	pb "videoWeb/generate-service/proto"
	"videoWeb/generate-service/svc"
)

func NewGenerateLoadBalanceClient(dis discovery.Discover, serverName string, tracer *zipkin.Tracer, logger log.Logger) (pb.GenerateServer, error) {
	// 服务发现
	instancer, err := dis.Discovery(serverName)
	if err != nil {
		return nil, err
	}

	options := []grpctransport.ClientOption{
		kitzipkin.GRPCClientTrace(tracer),
	}

	var endpoints svc.Endpoints
	// 获取int64key端点
	{
		factory := factoryGRPCFor(svc.MakeGenerateInt64KeyEndpoint, options...)
		endpoints.GenerateInt64KeyEndpoint = loadbalnace.BuildLoadBalance(instancer,
			factory,
			loadbalnace.Random,
			3,
			time.Millisecond*500,
			logger,
		)
	}
	// 获取stringkey端点
	{
		factory := factoryGRPCFor(svc.MakeGenerateStringKeyEndpoint, options...)
		endpoints.GenerateStringKeyEndpoint = loadbalnace.BuildLoadBalance(instancer,
			factory,
			loadbalnace.Random,
			3,
			time.Millisecond*500,
			logger,
		)
	}
	return endpoints, nil
}
