package grpc

import (
	"github.com/go-kit/kit/log"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/loadbalnace"
	"github.com/openzipkin/zipkin-go"
	"time"
	pb "videoWeb/customer-service/proto"
	"videoWeb/customer-service/svc"
)

func NewCustomerLoadBalanceClient(dis discovery.Discover, serverName string, tracer *zipkin.Tracer, logger log.Logger) (pb.CustomerServer, error) {
	// 服务发现
	instancer, err := dis.Discovery(serverName)
	if err != nil {
		return nil, err
	}

	options := []grpctransport.ClientOption{
		kitzipkin.GRPCClientTrace(tracer),
	}

	var endpoints svc.Endpoints
	// 获取通过phone注册端点
	{
		factory := factoryGRPCFor(svc.MakeRegisterByPhoneEndpoint, options...)
		endpoints.RegisterByPhoneEndpoint = loadbalnace.BuildLoadBalance(instancer,
			factory,
			loadbalnace.Random,
			3,
			time.Millisecond*500,
			logger,
		)
	}

	return endpoints, nil
}
