package main

import (
	"context"
	"fmt"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/log"
	"github.com/jjggzz/kj/track"
	"os"
	"testing"
	"time"
	"videoWeb/customer-service/proto"
	"videoWeb/customer-service/svc/client/grpc"
	genpb "videoWeb/generate-service/proto"
	genrpc "videoWeb/generate-service/svc/client/grpc"
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
	sum := 0
	for i := 0; i < 1000; i++ {
		deadline, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*100))
		response, err := server.LoginByPhone(deadline, &proto.LoginByPhoneRequest{Phone: "18376301879"})
		if err != nil {
			fmt.Println("error: " + err.Error())
			sum++
			continue
		}
		fmt.Println(response.Message)
		fmt.Println(response.Token)
	}
	fmt.Println(sum)
}

func Test_main2(t *testing.T) {

	logger := log.BuildLogger("test", os.Stderr)
	consulDiscovery := discovery.NewConsulDiscovery(
		"192.168.151.109:8500",
		"test2",
		6789,
		logger,
	)
	tracer, _ := track.BuildZipkinTracer("192.168.151.109:9411", "test2")
	server, _ := genrpc.NewGenerateLoadBalanceClient(
		consulDiscovery,
		"generate-service",
		tracer,
		logger,
	)
	sum := 0
	for i := 0; i < 1000; i++ {
		deadline, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*100))
		response, err := server.GenerateStringKey(deadline, &genpb.Empty{})
		if err != nil {
			fmt.Println("error: " + err.Error())
			sum++
			continue
		}
		fmt.Println(response.Id)
	}
	fmt.Println(sum)
}
