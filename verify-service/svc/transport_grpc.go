// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 84830579f0
// Version Date: Thu Dec 31 05:37:46 UTC 2020

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/jjggzz/kit/transport/grpc"

	// This Service
	pb "videoWeb/verify-service/proto"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC VerifyServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.VerifyServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// verify

		sendverifycode: grpctransport.NewServer(
			endpoints.SendVerifyCodeEndpoint,
			DecodeGRPCSendVerifyCodeRequest,
			EncodeGRPCSendVerifyCodeResponse,
			serverOptions...,
		),
		checkverifycode: grpctransport.NewServer(
			endpoints.CheckVerifyCodeEndpoint,
			DecodeGRPCCheckVerifyCodeRequest,
			EncodeGRPCCheckVerifyCodeResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the VerifyServer interface
type grpcServer struct {
	sendverifycode  grpctransport.Handler
	checkverifycode grpctransport.Handler
}

// Methods for grpcServer to implement VerifyServer interface

func (s *grpcServer) SendVerifyCode(ctx context.Context, req *pb.SendVerifyCodeRequest) (*pb.SendVerifyCodeResponse, error) {
	_, rep, err := s.sendverifycode.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendVerifyCodeResponse), nil
}

func (s *grpcServer) CheckVerifyCode(ctx context.Context, req *pb.CheckVerifyCodeRequest) (*pb.CheckVerifyCodeResponse, error) {
	_, rep, err := s.checkverifycode.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CheckVerifyCodeResponse), nil
}

// Server Decode

// DecodeGRPCSendVerifyCodeRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC sendverifycode request to a user-domain sendverifycode request. Primarily useful in a server.
func DecodeGRPCSendVerifyCodeRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SendVerifyCodeRequest)
	return req, nil
}

// DecodeGRPCCheckVerifyCodeRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC checkverifycode request to a user-domain checkverifycode request. Primarily useful in a server.
func DecodeGRPCCheckVerifyCodeRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CheckVerifyCodeRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCSendVerifyCodeResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain sendverifycode response to a gRPC sendverifycode reply. Primarily useful in a server.
func EncodeGRPCSendVerifyCodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.SendVerifyCodeResponse)
	return resp, nil
}

// EncodeGRPCCheckVerifyCodeResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain checkverifycode response to a gRPC checkverifycode reply. Primarily useful in a server.
func EncodeGRPCCheckVerifyCodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.CheckVerifyCodeResponse)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
