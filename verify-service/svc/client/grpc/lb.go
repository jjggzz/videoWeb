package grpc

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/loadbalnace"
	"github.com/openzipkin/zipkin-go"
	"google.golang.org/grpc"
	"io"
	"time"
	pb "videoWeb/proto"
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
	// 获取int64key端点
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
	// 获取stringkey端点
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

func factoryGRPCFor(makeEndpoint func(pb.VerifyServer) endpoint.Endpoint, options ...grpctransport.ClientOption) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		server, err := New(conn, options)
		return makeEndpoint(server), conn, err
	}
}
