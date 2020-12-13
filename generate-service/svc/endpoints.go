// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 1b850652b8
// Version Date: 2020-12-13T03:19:12Z

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	pb "videoWeb/generate-service/proto"
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
	GenerateInt64KeyEndpoint  endpoint.Endpoint
	GenerateStringKeyEndpoint endpoint.Endpoint
}

// Endpoints

func (e Endpoints) GenerateInt64Key(ctx context.Context, in *pb.Empty) (*pb.Int64KeyResponse, error) {
	response, err := e.GenerateInt64KeyEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Int64KeyResponse), nil
}

func (e Endpoints) GenerateStringKey(ctx context.Context, in *pb.Empty) (*pb.StringKeyResponse, error) {
	response, err := e.GenerateStringKeyEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.StringKeyResponse), nil
}

// Make Endpoints

func MakeGenerateInt64KeyEndpoint(s pb.GenerateServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.Empty)
		v, err := s.GenerateInt64Key(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeGenerateStringKeyEndpoint(s pb.GenerateServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.Empty)
		v, err := s.GenerateStringKey(ctx, req)
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
		"GenerateInt64Key":  {},
		"GenerateStringKey": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "GenerateInt64Key" {
			e.GenerateInt64KeyEndpoint = middleware(e.GenerateInt64KeyEndpoint)
		}
		if inc == "GenerateStringKey" {
			e.GenerateStringKeyEndpoint = middleware(e.GenerateStringKeyEndpoint)
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
		"GenerateInt64Key":  {},
		"GenerateStringKey": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "GenerateInt64Key" {
			e.GenerateInt64KeyEndpoint = middleware("GenerateInt64Key", e.GenerateInt64KeyEndpoint)
		}
		if inc == "GenerateStringKey" {
			e.GenerateStringKeyEndpoint = middleware("GenerateStringKey", e.GenerateStringKeyEndpoint)
		}
	}
}
