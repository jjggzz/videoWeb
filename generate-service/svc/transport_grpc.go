// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: a56d9007fd
// Version Date: Fri Dec 18 09:51:12 UTC 2020

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "videoWeb/generate-service/proto"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC GenerateServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.GenerateServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// generate

		generateint64key: grpctransport.NewServer(
			endpoints.GenerateInt64KeyEndpoint,
			DecodeGRPCGenerateInt64KeyRequest,
			EncodeGRPCGenerateInt64KeyResponse,
			serverOptions...,
		),
		generatestringkey: grpctransport.NewServer(
			endpoints.GenerateStringKeyEndpoint,
			DecodeGRPCGenerateStringKeyRequest,
			EncodeGRPCGenerateStringKeyResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the GenerateServer interface
type grpcServer struct {
	generateint64key  grpctransport.Handler
	generatestringkey grpctransport.Handler
}

// Methods for grpcServer to implement GenerateServer interface

func (s *grpcServer) GenerateInt64Key(ctx context.Context, req *pb.Empty) (*pb.Int64KeyResponse, error) {
	_, rep, err := s.generateint64key.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Int64KeyResponse), nil
}

func (s *grpcServer) GenerateStringKey(ctx context.Context, req *pb.Empty) (*pb.StringKeyResponse, error) {
	_, rep, err := s.generatestringkey.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.StringKeyResponse), nil
}

// Server Decode

// DecodeGRPCGenerateInt64KeyRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC generateint64key request to a user-domain generateint64key request. Primarily useful in a server.
func DecodeGRPCGenerateInt64KeyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Empty)
	return req, nil
}

// DecodeGRPCGenerateStringKeyRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC generatestringkey request to a user-domain generatestringkey request. Primarily useful in a server.
func DecodeGRPCGenerateStringKeyRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Empty)
	return req, nil
}

// Server Encode

// EncodeGRPCGenerateInt64KeyResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain generateint64key response to a gRPC generateint64key reply. Primarily useful in a server.
func EncodeGRPCGenerateInt64KeyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.Int64KeyResponse)
	return resp, nil
}

// EncodeGRPCGenerateStringKeyResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain generatestringkey response to a gRPC generatestringkey reply. Primarily useful in a server.
func EncodeGRPCGenerateStringKeyResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.StringKeyResponse)
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
