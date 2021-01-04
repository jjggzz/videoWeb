package handlers

import (
	"context"
	pb "videoWeb/notice-service/proto"
	"videoWeb/notice-service/service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.NoticeServer {
	return noticeService{}
}

type noticeService struct{}

func (s noticeService) SendSms(ctx context.Context, in *pb.SendSmsRequest) (*pb.SendSmsResponse, error) {
	var resp pb.SendSmsResponse
	return &resp, nil
}

func (s noticeService) SendEmail(ctx context.Context, in *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	var resp pb.SendEmailResponse
	code, err := service.Not.SendEmail(ctx, in.Strategy.String(), in.Title, in.Body, in.RecipientList)
	if err != nil {
		return nil, err
	}
	resp.Code = code.Code()
	return &resp, nil
}
