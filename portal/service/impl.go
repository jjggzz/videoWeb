package service

import (
	"context"
	"github.com/jjggzz/kj/errors"
	cuspb "videoWeb/customer-service/proto"
)

func (srv *service) Login(ctx context.Context, phone string) (string, error) {
	response, err := srv.cus.LoginByPhone(ctx, &cuspb.LoginByPhoneRequest{Phone: phone})
	if err != nil {
		return "", err
	}
	if !response.Result {
		return "", errors.New(response.Message)
	}
	return response.Token, nil
}
