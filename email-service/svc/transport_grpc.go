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
	pb "videoWeb/email-service/proto"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC EmailServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.EmailServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// email

		sendemail: grpctransport.NewServer(
			endpoints.SendEmailEndpoint,
			DecodeGRPCSendEmailRequest,
			EncodeGRPCSendEmailResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the EmailServer interface
type grpcServer struct {
	sendemail grpctransport.Handler
}

// Methods for grpcServer to implement EmailServer interface

func (s *grpcServer) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	_, rep, err := s.sendemail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendEmailResponse), nil
}

// Server Decode

// DecodeGRPCSendEmailRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC sendemail request to a user-domain sendemail request. Primarily useful in a server.
func DecodeGRPCSendEmailRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SendEmailRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCSendEmailResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain sendemail response to a gRPC sendemail reply. Primarily useful in a server.
func EncodeGRPCSendEmailResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.SendEmailResponse)
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
