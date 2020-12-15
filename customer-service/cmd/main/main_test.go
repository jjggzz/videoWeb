package main

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jjggzz/kj/discovery"
	"github.com/jjggzz/kj/log"
	"github.com/jjggzz/kj/track"
	"os"
	"testing"
	"time"
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

func Test_jwt(t *testing.T) {

	token, _ := createToken("123456", secret)
	fmt.Println(token)

	s, _ := parseToken(token, secret)

	fmt.Println(s, 0)

}

const secret = "DFAJLSDGEGFIA"

func createToken(uid string, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  uid,
		"name": "张三",
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func parseToken(token string, secret string) (int64, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}
	return claim.Claims.(jwt.MapClaims)["exp"].(int64), nil
}
