package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/endpoint"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/errors"
	"github.com/jjggzz/kj/log"
	"github.com/jjggzz/kj/track"
	"os"
	"testing"
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
	instancer, _ := consulDiscovery.Discovery("customer-service")
	tracer, _ := track.BuildZipkinTracer("192.168.151.109:9411", "test")
	server, _ := grpc.NewLoadBalanceClient(
		instancer,
		tracer,
		logger,
	)
	sum := 0
	for i := 0; i < 100; i++ {
		response, err := server.LoginByPhone(context.Background(), &proto.LoginByPhoneRequest{Phone: "18376301879"})
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
	instancer, _ := consulDiscovery.Discovery("generate-service")
	tracer, _ := track.BuildZipkinTracer("192.168.151.109:9411", "test2")

	//对hystrix进行配置
	//commandName:="my_endpoint"
	//hystrix.ConfigureCommand(commandName,hystrix.CommandConfig{
	//	Timeout:1000*3, //超时
	//	MaxConcurrentRequests:100, //最大并发的请求数
	//	RequestVolumeThreshold:5,//请求量阈值
	//	SleepWindow:10000, //熔断开启多久尝试发起一次请求
	//	ErrorPercentThreshold:1, //误差阈值百分比
	//})
	//breakerMw:=MyHystrix(commandName,fallback) //定义熔断器中间件
	//
	//server, _ := genrpc.NewLoadBalanceClient(
	//	instancer,
	//	breakerMw,
	//	tracer,
	//	logger,
	//)
	server, _ := genrpc.NewLoadBalanceClient(
		instancer,
		tracer,
		logger,
	)
	sum := 0
	for i := 0; i < 100; i++ {
		//deadline, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*100))
		response, err := server.GenerateStringKey(context.Background(), &genpb.Empty{})
		if err != nil {
			fmt.Println("error: " + err.Error())
			sum++
			continue
		}
		fmt.Println(response.Id)
	}
	fmt.Println(sum)
}

func MyHystrix(commandName string, fallback func(error) error) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var resp interface{}
			if err := hystrix.Do(commandName, func() (err error) {
				resp, err = next(ctx, request)
				return err
			}, fallback); err != nil {
				return nil, err
			}
			return resp, nil
		}
	}
}

func fallback(err error) error {
	return errors.New(err.Error()).Append("系统正在忙")
}
