package handlers

import (
	"context"

	pb "videoWeb/customer-service/proto"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.CustomerServer {
	return customerService{}
}

type customerService struct{}

func (s customerService) RegisterByPhone(ctx context.Context, in *pb.RegisterByPhoneRequest) (*pb.RegisterByPhoneResponse, error) {
	var resp pb.RegisterByPhoneResponse
	return &resp, nil
}

func (s customerService) SelectCustomer(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	var resp pb.Empty
	return &resp, nil
}
