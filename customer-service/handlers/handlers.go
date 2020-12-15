package handlers

import (
	"context"
	"github.com/jjggzz/kj/errors"
	"videoWeb/customer-service/service"

	pb "videoWeb/customer-service/proto"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.CustomerServer {
	return customerService{}
}

type customerService struct{}

func (s customerService) RegisterByPhone(ctx context.Context, in *pb.RegisterByPhoneRequest) (*pb.RegisterByPhoneResponse, error) {
	var resp pb.RegisterByPhoneResponse
	status, err := service.Cus.RegisterByPhone(ctx, in.Phone)
	if err != nil {
		return nil, err
	}
	switch status {
	case service.RegisterStatus_SUCCESS:
		resp.Result = true
		resp.Message = "注册成功"
		return &resp, nil
	case service.RegisterStatus_FAIL_HAS_USE:
		resp.Result = false
		resp.Message = "注册失败,该号码已注册"
		return &resp, nil
	default:
		return nil, errors.New("未知知结果类型")
	}

}

func (s customerService) LoginByPhone(ctx context.Context, in *pb.LoginByPhoneRequest) (*pb.LoginByPhoneResponse, error) {
	var resp pb.LoginByPhoneResponse
	return &resp, nil
}
