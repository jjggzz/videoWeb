package handlers

import (
	"github.com/jjggzz/kj/middleware"
	"github.com/sony/gobreaker"
	"videoWeb/generate-service/svc"
	pb "videoWeb/proto"
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
	// github.com/metaverse/truss/_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)

	// 限流
	limitMiddleware := middleware.LimitMiddleware(middleware.LimitDelay, 100)
	in.GenerateStringKeyEndpoint = limitMiddleware(in.GenerateStringKeyEndpoint)
	in.GenerateInt64KeyEndpoint = limitMiddleware(in.GenerateInt64KeyEndpoint)
	// 断路器
	breakerMiddleware := middleware.BreakerMiddleware(gobreaker.Settings{})
	in.GenerateStringKeyEndpoint = breakerMiddleware(in.GenerateStringKeyEndpoint)
	in.GenerateInt64KeyEndpoint = breakerMiddleware(in.GenerateInt64KeyEndpoint)

	return in
}

func WrapService(in pb.GenerateServer) pb.GenerateServer {

	return in
}
