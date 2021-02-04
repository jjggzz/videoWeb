package handlers

import (
	"context"
	"github.com/jjggzz/kj/utils"
	"videoWeb/video-service/dao/repository"
	"videoWeb/video-service/service"

	pb "videoWeb/video-service/proto"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.VideoServer {
	return videoService{}
}

type videoService struct{}

func (s videoService) UploadVideo(ctx context.Context, in *pb.UploadVideoRequest) (*pb.UploadVideoResponse, error) {
	var resp pb.UploadVideoResponse
	var video repository.Video
	_ = utils.Copy(in, &video)
	code, err := service.Vid.SaveVideo(ctx, &video)
	if err != nil {
		// todo 当前服务记录日志,并返回错误
		return nil, err
	}
	resp.Code = code.Code()
	return &resp, nil
}
