package service

import (
	"context"
	"videoWeb/customer/dao"
	"videoWeb/proto"
)

func (srv *service) QueryCustomerByAccessKey(ctx context.Context, accessKey string) (*dao.Customer, error) {
	return srv.dao.QueryCustomerByAccessKey(accessKey)
}

func (srv *service) InsertCustomer(ctx context.Context, customer *dao.Customer) error {
	key, err := srv.gen.GenerateStringKey(ctx, &proto.Empty{})
	if err != nil {
		return err
	}
	customer.AccessKey = key.Id
	return srv.dao.InsertCustomer(customer)
}

func (srv *service) GetKey(ctx context.Context) (string, error) {
	response, err := srv.gen.GenerateStringKey(ctx, &proto.Empty{})
	if err != nil {
		return "", err
	}
	return response.Id, err
}
