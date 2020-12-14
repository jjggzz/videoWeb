// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 1b850652b8
// Version Date: Sun Dec 13 03:19:12 UTC 2020

// Package grpc provides a gRPC client for the Customer service.
package grpc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"

	// This Service
	pb "videoWeb/customer-service/proto"
	"videoWeb/customer-service/svc"
)

func factoryGRPCFor(makeEndpoint func(pb.CustomerServer) endpoint.Endpoint, clientOptions ...grpctransport.ClientOption) sd.Factory {
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
func New(conn *grpc.ClientConn, clientOptions []grpctransport.ClientOption, options ...ClientOption) (pb.CustomerServer, error) {
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
	var registerbyphoneEndpoint endpoint.Endpoint
	{
		registerbyphoneEndpoint = grpctransport.NewClient(
			conn,
			"proto.Customer",
			"RegisterByPhone",
			EncodeGRPCRegisterByPhoneRequest,
			DecodeGRPCRegisterByPhoneResponse,
			pb.RegisterByPhoneResponse{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		RegisterByPhoneEndpoint: registerbyphoneEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCRegisterByPhoneResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC registerbyphone reply to a user-domain registerbyphone response. Primarily useful in a client.
func DecodeGRPCRegisterByPhoneResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.RegisterByPhoneResponse)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCRegisterByPhoneRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain registerbyphone request to a gRPC registerbyphone request. Primarily useful in a client.
func EncodeGRPCRegisterByPhoneRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.RegisterByPhoneRequest)
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
