package handlers

import (
	"context"
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
		resp.Result = false
		resp.Message = "系统异常"
		return &resp, nil
	}

}

func (s customerService) LoginByPhone(ctx context.Context, in *pb.LoginByPhoneRequest) (*pb.LoginByPhoneResponse, error) {
	var resp pb.LoginByPhoneResponse
	status, token, err := service.Cus.LoginByPhone(ctx, in.Phone)
	if err != nil {
		return nil, err
	}
	switch status {
	case service.LoginStatus_SUCCESS:
		resp.Result = true
		resp.Message = "登录成功"
		resp.Token = token
		return &resp, nil
	case service.LoginStatus_FAIL_NOT_REG:
		resp.Result = false
		resp.Message = "未注册"
		resp.Token = token
		return &resp, nil
	case service.LoginStatus_FAIL_DISABLE:
		resp.Result = false
		resp.Message = "账户被冻结"
		resp.Token = token
		return &resp, nil
	default:
		resp.Result = false
		resp.Message = "系统异常"
		resp.Token = token
		return &resp, nil
	}
}

func (s customerService) GetCustomerInfoByToken(ctx context.Context, in *pb.GetCustomerInfoByTokenRequest) (*pb.GetCustomerInfoByTokenResponse, error) {
	var resp pb.GetCustomerInfoByTokenResponse
	return &resp, nil
}
