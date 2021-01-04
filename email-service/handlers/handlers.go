package handlers

import (
	"context"
	"videoWeb/email-service/service"

	pb "videoWeb/email-service/proto"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.EmailServer {
	return emailService{}
}

type emailService struct{}

func (s emailService) SendEmail(ctx context.Context, in *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	var resp pb.SendEmailResponse
	code, err := service.Ema.SendEmail(ctx, in.Strategy.String(), in.Title, in.Body, in.RecipientList)
	if err != nil {
		return nil, err
	}
	resp.Code = code.Code()
	return &resp, nil
}
