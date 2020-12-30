package handlers

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/jjggzz/kj/middleware"
	"github.com/sony/gobreaker"
	"log"
	pb "videoWeb/customer-service/proto"
	"videoWeb/customer-service/svc"
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// github.com/jjggzz/truss/_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)
	// 限流
	limitMiddleware := middleware.LimitMiddleware(middleware.LimitDelay, 100)
	// 断路器
	breakerMiddleware := middleware.BreakerMiddleware(gobreaker.Settings{})
	errorMiddleware := ErrorMiddleware()

	in.WrapAllExcept(limitMiddleware)
	in.WrapAllExcept(breakerMiddleware)
	in.WrapAllExcept(errorMiddleware)
	return in
}

func WrapService(in pb.CustomerServer) pb.CustomerServer {
	return in
}

func ErrorMiddleware() endpoint.Middleware {
	return func(i endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			res, err := i(ctx, request)
			if err != nil {
				log.Println(err)
			}
			return res, err
		}
	}
}
