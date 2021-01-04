package handlers

import (
	"context"

	pb "videoWeb/sms-service/proto"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.SmsServer {
	return smsService{}
}

type smsService struct{}

func (s smsService) SendSms(ctx context.Context, in *pb.SendSmsRequest) (*pb.SendSmsResponse, error) {
	var resp pb.SendSmsResponse
	return &resp, nil
}
