package handlers

import (
	"context"
	"github.com/jjggzz/kj/errors"
	"videoWeb/common/ecode"
	pb "videoWeb/verify-service/proto"
	"videoWeb/verify-service/service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.VerifyServer {
	return verifyService{}
}

type verifyService struct{}

func (s verifyService) SendVerifyCode(ctx context.Context, in *pb.SendVerifyCodeRequest) (*pb.SendVerifyCodeResponse, error) {
	var (
		code = ecode.ServerErr
		err  error
		resp pb.SendVerifyCodeResponse
	)
	switch in.Strategy {
	case pb.VerifyTargetStrategy_PHONE:
		code, err = service.Ver.SendPhoneVerify(ctx, in.Target)
	case pb.VerifyTargetStrategy_EMAIL:
		code, err = service.Ver.SendEmailVerify(ctx, in.Target)
	default:
		err = errors.New("不支持的策略")
	}
	if err != nil {
		return nil, err
	}
	resp.Code = code.Code()
	return &resp, err
}

func (s verifyService) CheckVerifyCode(ctx context.Context, in *pb.CheckVerifyCodeRequest) (*pb.CheckVerifyCodeResponse, error) {
	var (
		resp pb.CheckVerifyCodeResponse
	)
	code, err := service.Ver.CheckVerify(ctx, in.Target, in.VerifyCode)
	if err != nil {
		return nil, err
	}
	resp.Code = code.Code()
	return &resp, err
}
