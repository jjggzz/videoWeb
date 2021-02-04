// author: JGZ
// time:   2021-02-04 16:58
package service

import (
	"context"
	"time"
	"videoWeb/common/ecode"
	genpb "videoWeb/generate-service/proto"
	"videoWeb/video-service/dao/repository"
)

func (svc *service) SaveVideo(ctx context.Context, video *repository.Video) (ecode.ECode, error) {
	response, err := svc.gen.GenerateStringKey(ctx, &genpb.Empty{})
	if err != nil {
		return ecode.ServerErr, err
	}
	// 调用第三方服务出现业务误错，返回对应的错误
	if response.Code != ecode.Success.Code() {
		return ecode.Build(response.Code), nil
	}
	now := time.Now()
	video.AccessKey = response.Id
	video.CreateTime = &now
	video.UpdateTime = &now
	_, err = svc.dao.Repo.VideoRepo.Insert(video)
	if err != nil {
		return ecode.ServerErr, err
	}
	return ecode.Success, nil
}
