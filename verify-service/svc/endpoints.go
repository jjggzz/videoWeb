// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 84830579f0
// Version Date: Thu Dec 31 05:37:46 UTC 2020

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/jjggzz/kit/endpoint"

	pb "videoWeb/verify-service/proto"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	SendVerifyCodeEndpoint  endpoint.Endpoint
	CheckVerifyCodeEndpoint endpoint.Endpoint
}

// Endpoints

func (e Endpoints) SendVerifyCode(ctx context.Context, in *pb.SendVerifyCodeRequest) (*pb.SendVerifyCodeResponse, error) {
	response, err := e.SendVerifyCodeEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.SendVerifyCodeResponse), nil
}

func (e Endpoints) CheckVerifyCode(ctx context.Context, in *pb.CheckVerifyCodeRequest) (*pb.CheckVerifyCodeResponse, error) {
	response, err := e.CheckVerifyCodeEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.CheckVerifyCodeResponse), nil
}

// Make Endpoints

func MakeSendVerifyCodeEndpoint(s pb.VerifyServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.SendVerifyCodeRequest)
		v, err := s.SendVerifyCode(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeCheckVerifyCodeEndpoint(s pb.VerifyServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.CheckVerifyCodeRequest)
		v, err := s.CheckVerifyCode(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"SendVerifyCode":  {},
		"CheckVerifyCode": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "SendVerifyCode" {
			e.SendVerifyCodeEndpoint = middleware(e.SendVerifyCodeEndpoint)
		}
		if inc == "CheckVerifyCode" {
			e.CheckVerifyCodeEndpoint = middleware(e.CheckVerifyCodeEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"SendVerifyCode":  {},
		"CheckVerifyCode": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "SendVerifyCode" {
			e.SendVerifyCodeEndpoint = middleware("SendVerifyCode", e.SendVerifyCodeEndpoint)
		}
		if inc == "CheckVerifyCode" {
			e.CheckVerifyCodeEndpoint = middleware("CheckVerifyCode", e.CheckVerifyCodeEndpoint)
		}
	}
}
