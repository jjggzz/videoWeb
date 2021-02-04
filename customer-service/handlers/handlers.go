package handlers

import (
	"context"
	"videoWeb/common/ecode"
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
	code, err := service.Cus.RegisterByPhone(ctx, in.Phone)
	if err != nil {
		// todo 当前服务记录日志,并返回错误
		return nil, err
	}
	resp.Code = code.Code()
	return &resp, nil
}

func (s customerService) LoginByPhone(ctx context.Context, in *pb.LoginByPhoneRequest) (*pb.LoginByPhoneResponse, error) {
	var resp pb.LoginByPhoneResponse
	code, token, err := service.Cus.LoginByPhone(ctx, in.Phone)
	if err != nil {
		// todo 当前服务记录日志,并返回错误
		return nil, err
	}
	resp.Code = code.Code()
	resp.Token = token
	return &resp, nil
}

func (s customerService) GetCustomerInfoByToken(ctx context.Context, in *pb.GetCustomerInfoByTokenRequest) (*pb.GetCustomerInfoByTokenResponse, error) {
	var resp pb.GetCustomerInfoByTokenResponse
	code, customer, err := service.Cus.GetCustomerInfoByToken(ctx, in.Token)
	if err != nil {
		// todo 当前服务记录日志,并返回错误
		return nil, err
	}

	resp.Code = code.Code()
	if code == ecode.Success {
		resp.Username = customer.Username
		resp.AccessKey = customer.AccessKey
		resp.Email = customer.Email
		resp.Nickname = customer.Nickname
	}
	return &resp, nil
}
