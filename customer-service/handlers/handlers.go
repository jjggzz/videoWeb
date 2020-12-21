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
	switch status {
	case service.RegisterStatusSuccess:
		resp.Result = true
		resp.Message = "注册成功"
	case service.RegisterStatusFailHasUse:
		resp.Result = false
		resp.Message = "注册失败,该号码已注册"
	default:
		resp.Result = false
		resp.Message = "系统异常"
	}
	return &resp, err
}

func (s customerService) LoginByPhone(ctx context.Context, in *pb.LoginByPhoneRequest) (*pb.LoginByPhoneResponse, error) {
	var resp pb.LoginByPhoneResponse
	status, token, err := service.Cus.LoginByPhone(ctx, in.Phone)
	switch status {
	case service.LoginStatusSuccess:
		resp.Result = true
		resp.Message = "登录成功"
		resp.Token = token
	case service.LoginStatusFailNotReg:
		resp.Result = false
		resp.Message = "未注册"
		resp.Token = token
	case service.LoginStatusFailDisable:
		resp.Result = false
		resp.Message = "账户被冻结"
		resp.Token = token
	default:
		resp.Result = false
		resp.Message = "系统异常"
		resp.Token = token
	}
	return &resp, err
}

func (s customerService) GetCustomerInfoByToken(ctx context.Context, in *pb.GetCustomerInfoByTokenRequest) (*pb.GetCustomerInfoByTokenResponse, error) {
	var resp pb.GetCustomerInfoByTokenResponse
	status, customer, err := service.Cus.GetCustomerInfoByToken(ctx, in.Token)
	switch status {
	case service.GetInfoStatusSuccess:
		resp.Result = true
		resp.Message = "获取成功"
		if customer != nil {
			resp.CustomerInfo = new(pb.CustomerInfo)
			resp.CustomerInfo.Username = customer.Username
			resp.CustomerInfo.AccessKey = customer.AccessKey
			resp.CustomerInfo.Email = customer.Email
			resp.CustomerInfo.Nickname = customer.Nickname
		}
	case service.GetInfoStatusFailNotLogin:
		resp.Result = false
		resp.Message = "未登陆"
	default:
		resp.Result = false
		resp.Message = "系统异常"
	}

	return &resp, err
}
