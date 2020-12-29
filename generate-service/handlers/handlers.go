package handlers

import (
	"context"
	"videoWeb/common/ecode"
	"videoWeb/generate-service/service"

	pb "videoWeb/generate-service/proto"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.GenerateServer {
	return generateService{}
}

type generateService struct{}

func (s generateService) GenerateInt64Key(ctx context.Context, in *pb.Empty) (*pb.Int64KeyResponse, error) {
	var resp pb.Int64KeyResponse
	id, err := service.Gen.GenerateInt64Key(ctx)
	if err != nil {
		return nil, err
	}
	resp.Code = ecode.Success.Code()
	resp.Id = id
	return &resp, nil
}

func (s generateService) GenerateStringKey(ctx context.Context, in *pb.Empty) (*pb.StringKeyResponse, error) {
	var resp pb.StringKeyResponse
	id, err := service.Gen.GenerateStringKey(ctx)
	if err != nil {
		return nil, err
	}
	resp.Code = ecode.Success.Code()
	resp.Id = id
	return &resp, nil
}
