package service

import (
	"context"
	"videoWeb/common/ecode"
	cuspb "videoWeb/customer-service/proto"
	verpb "videoWeb/verify-service/proto"
)

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

func (srv *service) SendVerify(ctx context.Context, phone string) (ecode.ECode, error) {
	sendResponse, err := srv.ver.SendVerifyCode(ctx, &verpb.SendVerifyCodeRequest{Target: phone, Strategy: verpb.VerifyTargetStrategy_PHONE})
	if err != nil {
		return ecode.ServerErr, err
	}
	return ecode.Build(sendResponse.Code), nil
}
