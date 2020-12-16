// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 1b850652b8
// Version Date: Sun Dec 13 03:19:12 UTC 2020

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "videoWeb/customer-service/proto"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC CustomerServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.CustomerServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// customer

		registerbyphone: grpctransport.NewServer(
			endpoints.RegisterByPhoneEndpoint,
			DecodeGRPCRegisterByPhoneRequest,
			EncodeGRPCRegisterByPhoneResponse,
			serverOptions...,
		),
		loginbyphone: grpctransport.NewServer(
			endpoints.LoginByPhoneEndpoint,
			DecodeGRPCLoginByPhoneRequest,
			EncodeGRPCLoginByPhoneResponse,
			serverOptions...,
		),
		getcustomerinfobytoken: grpctransport.NewServer(
			endpoints.GetCustomerInfoByTokenEndpoint,
			DecodeGRPCGetCustomerInfoByTokenRequest,
			EncodeGRPCGetCustomerInfoByTokenResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the CustomerServer interface
type grpcServer struct {
	registerbyphone        grpctransport.Handler
	loginbyphone           grpctransport.Handler
	getcustomerinfobytoken grpctransport.Handler
}

// Methods for grpcServer to implement CustomerServer interface

func (s *grpcServer) RegisterByPhone(ctx context.Context, req *pb.RegisterByPhoneRequest) (*pb.RegisterByPhoneResponse, error) {
	_, rep, err := s.registerbyphone.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RegisterByPhoneResponse), nil
}

func (s *grpcServer) LoginByPhone(ctx context.Context, req *pb.LoginByPhoneRequest) (*pb.LoginByPhoneResponse, error) {
	_, rep, err := s.loginbyphone.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginByPhoneResponse), nil
}

func (s *grpcServer) GetCustomerInfoByToken(ctx context.Context, req *pb.GetCustomerInfoByTokenRequest) (*pb.GetCustomerInfoByTokenResponse, error) {
	_, rep, err := s.getcustomerinfobytoken.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetCustomerInfoByTokenResponse), nil
}

// Server Decode

// DecodeGRPCRegisterByPhoneRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC registerbyphone request to a user-domain registerbyphone request. Primarily useful in a server.
func DecodeGRPCRegisterByPhoneRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.RegisterByPhoneRequest)
	return req, nil
}

// DecodeGRPCLoginByPhoneRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC loginbyphone request to a user-domain loginbyphone request. Primarily useful in a server.
func DecodeGRPCLoginByPhoneRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.LoginByPhoneRequest)
	return req, nil
}

// DecodeGRPCGetCustomerInfoByTokenRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getcustomerinfobytoken request to a user-domain getcustomerinfobytoken request. Primarily useful in a server.
func DecodeGRPCGetCustomerInfoByTokenRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetCustomerInfoByTokenRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCRegisterByPhoneResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain registerbyphone response to a gRPC registerbyphone reply. Primarily useful in a server.
func EncodeGRPCRegisterByPhoneResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.RegisterByPhoneResponse)
	return resp, nil
}

// EncodeGRPCLoginByPhoneResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain loginbyphone response to a gRPC loginbyphone reply. Primarily useful in a server.
func EncodeGRPCLoginByPhoneResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginByPhoneResponse)
	return resp, nil
}

// EncodeGRPCGetCustomerInfoByTokenResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getcustomerinfobytoken response to a gRPC getcustomerinfobytoken reply. Primarily useful in a server.
func EncodeGRPCGetCustomerInfoByTokenResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetCustomerInfoByTokenResponse)
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
