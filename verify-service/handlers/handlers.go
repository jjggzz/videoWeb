package handlers

import (
	"context"

	pb "videoWeb/proto"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.VerifyServer {
	return verifyService{}
}

type verifyService struct{}

func (s verifyService) SendVerifyCode(ctx context.Context, in *pb.SendVerifyCodeRequest) (*pb.Empty, error) {
	var resp pb.Empty
	return &resp, nil
}

func (s verifyService) CheckVerifyCode(ctx context.Context, in *pb.CheckVerifyCodeRequest) (*pb.CheckVerifyCodeResponse, error) {
	var resp pb.CheckVerifyCodeResponse
	return &resp, nil
}
