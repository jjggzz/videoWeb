package main

import (
	"context"
	"fmt"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/log"
	"github.com/jjggzz/kj/track"
	"os"
	"testing"
	"videoWeb/customer-service/proto"
	"videoWeb/customer-service/svc/client/grpc"
)

func Test_main(t *testing.T) {

	logger := log.BuildLogger("test", os.Stderr)
	consulDiscovery := discovery.NewConsulDiscovery(
		"192.168.151.109:8500",
		"test",
		6789,
		logger,
	)
	tracer, _ := track.BuildZipkinTracer("192.168.151.109:9411", "test")
	server, _ := grpc.NewCustomerLoadBalanceClient(
		consulDiscovery,
		"customer-service",
		tracer,
		logger,
	)
	response, err := server.RegisterByPhone(context.Background(), &proto.RegisterByPhoneRequest{Phone: "18376301879"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response.Message)
}
