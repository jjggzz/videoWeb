package handlers

import (
	"context"
	"github.com/jjggzz/kj/errors"
	pb "videoWeb/verify-service/proto"
	"videoWeb/verify-service/service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.VerifyServer {
	return verifyService{}
}

type verifyService struct{}

func (s verifyService) SendVerifyCode(ctx context.Context, in *pb.SendVerifyCodeRequest) (*pb.Empty, error) {
	var (
		resp pb.Empty
		err  error
	)
	switch in.Strategy {
	case pb.VerifyTargetStrategy_PHONE:
		err = service.Ver.SendPhoneVerify(ctx, in.Target)
	case pb.VerifyTargetStrategy_EMAIL:
		err = service.Ver.SendEmailVerify(ctx, in.Target)
	default:
		err = errors.New("不支持的策略")
	}
	return &resp, err
}

func (s verifyService) CheckVerifyCode(ctx context.Context, in *pb.CheckVerifyCodeRequest) (*pb.CheckVerifyCodeResponse, error) {
	var (
		resp pb.CheckVerifyCodeResponse
	)
	result, err := service.Ver.CheckVerify(ctx, in.Target, in.Code)
	if !result {
		resp.Result = false
		resp.Message = "验证码错误"
	} else {
		resp.Result = true
		resp.Message = "校验成功"
	}

	return &resp, err
}
