package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/log"
	"github.com/jjggzz/kj/track"
	"os"
	"testing"
	"videoWeb/notice-service/proto"
	"videoWeb/notice-service/svc/client/grpc"
)

func Test_main(t *testing.T) {
	logger := log.BuildLogger("test", os.Stderr)
	consulDiscovery := discovery.NewConsulDiscovery(
		"192.168.151.109:8500",
		"test",
		6789,
		logger,
	)
	instancer, _ := consulDiscovery.Discovery("notice-service")
	tracer, _ := track.BuildZipkinTracer("192.168.151.109:9411", "test")
	server, _ := grpc.NewLoadBalanceClient(
		instancer,
		tracer,
		logger,
	)
	request := proto.SendEmailRequest{
		Strategy:      proto.SMTPStrategy_QQ,
		Title:         "测试用email服务",
		Body:          "Good Good Study, Day Day Up!!!!!!",
		RecipientList: []string{"1945282561@qq.com"},
	}
	response, err := server.SendEmail(context.Background(), &request)
	if err != nil {
		_ = level.Error(logger).Log(err)
		return
	}
	fmt.Println(response.Code)
}
