package service

import (
	"context"
	"videoWeb/customer/dao"
	genpb "videoWeb/generate-service/proto"
	verpb "videoWeb/verify-service/proto"
)

func (srv *service) QueryCustomerByAccessKey(ctx context.Context, accessKey string) (*dao.Customer, error) {
	return srv.dao.QueryCustomerByAccessKey(accessKey)
}

func (srv *service) InsertCustomer(ctx context.Context, customer *dao.Customer) error {
	key, err := srv.gen.GenerateStringKey(ctx, &genpb.Empty{})
	if err != nil {
		return err
	}
	customer.AccessKey = key.Id
	return srv.dao.InsertCustomer(customer)
}

func (srv *service) GetKey(ctx context.Context) (string, error) {
	response, err := srv.gen.GenerateStringKey(ctx, &genpb.Empty{})
	if err != nil {
		return "", err
	}
	return response.Id, err
}

func (srv *service) SendVerify(ctx context.Context, phone string) error {
	_, err := srv.ver.SendVerifyCode(ctx, &verpb.SendVerifyCodeRequest{Target: phone, Strategy: verpb.VerifyTargetStrategy_PHONE})
	if err != nil {
		return err
	}
	return nil
}

func (srv *service) CheckVerify(ctx context.Context, phone string, code string) (bool, error) {
	response, err := srv.ver.CheckVerifyCode(ctx, &verpb.CheckVerifyCodeRequest{Target: phone, Code: code})
	if err != nil {
		return false, err
	}
	if response.Result == verpb.VerifyCheckResults_SUCCESS {
		return true, nil
	}
	return false, nil
}
