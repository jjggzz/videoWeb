// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 84830579f0
// Version Date: Thu Dec 31 05:37:46 UTC 2020

// Package grpc provides a gRPC client for the Verify service.
package grpc

import (
	"context"
	"github.com/jjggzz/kit/endpoint"
	"github.com/jjggzz/kit/log"
	"github.com/jjggzz/kit/sd"
	"github.com/jjggzz/kit/sd/lb"
	kitzipkin "github.com/jjggzz/kit/tracing/zipkin"
	grpctransport "github.com/jjggzz/kit/transport/grpc"
	"github.com/openzipkin/zipkin-go"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"time"

	// This Service
	pb "videoWeb/verify-service/proto"
	"videoWeb/verify-service/svc"
)

func NewLoadBalanceClient(instancer sd.Instancer, tracer *zipkin.Tracer, logger log.Logger) (pb.VerifyServer, error) {
	options := []grpctransport.ClientOption{}
	if tracer != nil {
		options = append(options, kitzipkin.GRPCClientTrace(tracer))
	}

	var endpoints svc.Endpoints
	{
		factory := factoryGRPCFor(svc.MakeSendVerifyCodeEndpoint, options...)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		endpoints.SendVerifyCodeEndpoint = lb.Retry(3, time.Millisecond*3000, balancer)
	}

	{
		factory := factoryGRPCFor(svc.MakeCheckVerifyCodeEndpoint, options...)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		endpoints.CheckVerifyCodeEndpoint = lb.Retry(3, time.Millisecond*3000, balancer)
	}

	return endpoints, nil
}

func factoryGRPCFor(makeEndpoint func(pb.VerifyServer) endpoint.Endpoint, clientOptions ...grpctransport.ClientOption) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		server, err := New(conn, clientOptions)
		return makeEndpoint(server), conn, err
	}
}

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, clientOptions []grpctransport.ClientOption, options ...ClientOption) (pb.VerifyServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions = append(clientOptions, grpctransport.ClientBefore(contextValuesToGRPCMetadata(cc.headers)))

	//clientOptions := []grpctransport.ClientOption{
	//	grpctransport.ClientBefore(
	//		contextValuesToGRPCMetadata(cc.headers)),
	//}
	var sendverifycodeEndpoint endpoint.Endpoint
	{
		sendverifycodeEndpoint = grpctransport.NewClient(
			conn,
			"proto.Verify",
			"SendVerifyCode",
			EncodeGRPCSendVerifyCodeRequest,
			DecodeGRPCSendVerifyCodeResponse,
			pb.SendVerifyCodeResponse{},
			clientOptions...,
		).Endpoint()
	}

	var checkverifycodeEndpoint endpoint.Endpoint
	{
		checkverifycodeEndpoint = grpctransport.NewClient(
			conn,
			"proto.Verify",
			"CheckVerifyCode",
			EncodeGRPCCheckVerifyCodeRequest,
			DecodeGRPCCheckVerifyCodeResponse,
			pb.CheckVerifyCodeResponse{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		SendVerifyCodeEndpoint:  sendverifycodeEndpoint,
		CheckVerifyCodeEndpoint: checkverifycodeEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCSendVerifyCodeResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC sendverifycode reply to a user-domain sendverifycode response. Primarily useful in a client.
func DecodeGRPCSendVerifyCodeResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.SendVerifyCodeResponse)
	return reply, nil
}

// DecodeGRPCCheckVerifyCodeResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC checkverifycode reply to a user-domain checkverifycode response. Primarily useful in a client.
func DecodeGRPCCheckVerifyCodeResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CheckVerifyCodeResponse)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCSendVerifyCodeRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain sendverifycode request to a gRPC sendverifycode request. Primarily useful in a client.
func EncodeGRPCSendVerifyCodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SendVerifyCodeRequest)
	return req, nil
}

// EncodeGRPCCheckVerifyCodeRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain checkverifycode request to a gRPC checkverifycode request. Primarily useful in a client.
func EncodeGRPCCheckVerifyCodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CheckVerifyCodeRequest)
	return req, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}
