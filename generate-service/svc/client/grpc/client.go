// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 84830579f0
// Version Date: Thu Dec 31 05:37:46 UTC 2020

// Package grpc provides a gRPC client for the Generate service.
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
	pb "videoWeb/generate-service/proto"
	"videoWeb/generate-service/svc"
)

func NewLoadBalanceClient(instancer sd.Instancer, tracer *zipkin.Tracer, logger log.Logger) (pb.GenerateServer, error) {
	options := []grpctransport.ClientOption{}
	if tracer != nil {
		options = append(options, kitzipkin.GRPCClientTrace(tracer))
	}

	var endpoints svc.Endpoints
	{
		factory := factoryGRPCFor(svc.MakeGenerateInt64KeyEndpoint, options...)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		endpoints.GenerateInt64KeyEndpoint = lb.Retry(3, time.Millisecond*3000, balancer)
	}

	{
		factory := factoryGRPCFor(svc.MakeGenerateStringKeyEndpoint, options...)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		endpoints.GenerateStringKeyEndpoint = lb.Retry(3, time.Millisecond*3000, balancer)
	}

	return endpoints, nil
}

func factoryGRPCFor(makeEndpoint func(pb.GenerateServer) endpoint.Endpoint, clientOptions ...grpctransport.ClientOption) sd.Factory {
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
func New(conn *grpc.ClientConn, clientOptions []grpctransport.ClientOption, options ...ClientOption) (pb.GenerateServer, error) {
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
	var generateint64keyEndpoint endpoint.Endpoint
	{
		generateint64keyEndpoint = grpctransport.NewClient(
			conn,
			"proto.Generate",
			"GenerateInt64Key",
			EncodeGRPCGenerateInt64KeyRequest,
			DecodeGRPCGenerateInt64KeyResponse,
			pb.Int64KeyResponse{},
			clientOptions...,
		).Endpoint()
	}

	var generatestringkeyEndpoint endpoint.Endpoint
	{
		generatestringkeyEndpoint = grpctransport.NewClient(
			conn,
			"proto.Generate",
			"GenerateStringKey",
			EncodeGRPCGenerateStringKeyRequest,
			DecodeGRPCGenerateStringKeyResponse,
			pb.StringKeyResponse{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		GenerateInt64KeyEndpoint:  generateint64keyEndpoint,
		GenerateStringKeyEndpoint: generatestringkeyEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCGenerateInt64KeyResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC generateint64key reply to a user-domain generateint64key response. Primarily useful in a client.
func DecodeGRPCGenerateInt64KeyResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.Int64KeyResponse)
	return reply, nil
}

// DecodeGRPCGenerateStringKeyResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC generatestringkey reply to a user-domain generatestringkey response. Primarily useful in a client.
func DecodeGRPCGenerateStringKeyResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.StringKeyResponse)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCGenerateInt64KeyRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain generateint64key request to a gRPC generateint64key request. Primarily useful in a client.
func EncodeGRPCGenerateInt64KeyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Empty)
	return req, nil
}

// EncodeGRPCGenerateStringKeyRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain generatestringkey request to a gRPC generatestringkey request. Primarily useful in a client.
func EncodeGRPCGenerateStringKeyRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.Empty)
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
