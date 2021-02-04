package service

import (
	"context"
	"github.com/jjggzz/kj/utils"
	"videoWeb/common/ecode"
	cuspb "videoWeb/customer-service/proto"
	"videoWeb/portal/service/dto"
	verpb "videoWeb/verify-service/proto"
	vidpb "videoWeb/video-service/proto"
)

// 用户登录
func (srv *service) Login(ctx context.Context, phone string, verify string) (ecode.ECode, string, error) {
	checkResponse, err := srv.ver.CheckVerifyCode(ctx, &verpb.CheckVerifyCodeRequest{Target: phone, VerifyCode: verify})
	if err != nil {
		return ecode.ServerErr, "", err
	}
	if checkResponse.Code != ecode.Success.Code() {
		return ecode.Build(checkResponse.Code), "", nil
	}
	loginResponse, err := srv.cus.LoginByPhone(ctx, &cuspb.LoginByPhoneRequest{Phone: phone})
	if err != nil {
		return ecode.ServerErr, "", err
	}
	if loginResponse.Code != ecode.Success.Code() {
		return ecode.Build(loginResponse.Code), "", nil
	}
	return ecode.Success, loginResponse.Token, nil
}

// 用户注册
func (srv *service) Register(ctx context.Context, phone string, verify string) (ecode.ECode, error) {
	checkResponse, err := srv.ver.CheckVerifyCode(ctx, &verpb.CheckVerifyCodeRequest{Target: phone, VerifyCode: verify})
	if err != nil {
		return ecode.ServerErr, err
	}
	if checkResponse.Code != ecode.Success.Code() {
		return ecode.Build(checkResponse.Code), nil
	}
	registerResponse, err := srv.cus.RegisterByPhone(ctx, &cuspb.RegisterByPhoneRequest{Phone: phone})
	if err != nil {
		return ecode.ServerErr, err
	}
	if registerResponse.Code != ecode.Success.Code() {
		return ecode.Build(registerResponse.Code), nil
	}
	return ecode.Success, nil
}

// 获取用户信息
func (srv *service) GetUserInfo(ctx context.Context, token string) (ecode.ECode, *dto.UserInfoDto, error) {
	var userInfo dto.UserInfoDto
	getUserInfoResponse, err := srv.cus.GetCustomerInfoByToken(ctx, &cuspb.GetCustomerInfoByTokenRequest{Token: token})
	if err != nil {
		return ecode.ServerErr, nil, err
	}
	if getUserInfoResponse.Code != ecode.Success.Code() {
		return ecode.Build(getUserInfoResponse.Code), nil, nil
	}
	_ = utils.Copy(getUserInfoResponse, &userInfo)
	return ecode.Success, &userInfo, nil
}

// 发送验证码
func (srv *service) SendVerify(ctx context.Context, strategyName string, phone string) (ecode.ECode, error) {
	sendResponse, err := srv.ver.SendVerifyCode(ctx, &verpb.SendVerifyCodeRequest{Target: phone, Strategy: strategy[strategyName]})
	if err != nil {
		return ecode.ServerErr, err
	}
	return ecode.Build(sendResponse.Code), nil
}

// 上传视频
func (srv *service) UploadVideo(ctx context.Context, videoDto *dto.UploadVideoDto) (ecode.ECode, error) {
	var uploadRequest vidpb.UploadVideoRequest
	_ = utils.Copy(videoDto, &uploadRequest)
	uploadRequest.CustomerId = ctx.Value("customerInfo").(*dto.UserInfoDto).Id
	uploadResponse, err := srv.vid.UploadVideo(ctx, &uploadRequest)
	if err != nil {
		return ecode.ServerErr, err
	}
	return ecode.Build(uploadResponse.Code), nil
}
