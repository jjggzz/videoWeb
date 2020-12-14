package service

import (
	"context"
	"time"
	"videoWeb/customer-service/dao"
	genpb "videoWeb/generate-service/proto"
)

func (srv *service) RegisterByPhone(ctx context.Context, phone string) (RegisterStatus, error) {
	// 电话号码是否被使用
	exist, err := srv.dao.ExitsCustomerByPhone(phone)
	if exist {
		return RegisterStatus_FAIL_HAS_USE, nil
	}
	response, err := srv.gen.GenerateStringKey(ctx, &genpb.Empty{})
	if err != nil {
		return RegisterStatus_FAIL_SYS_ERR, err
	}
	// 插入数据
	customer := &dao.Customer{
		AccessKey:    response.Id,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		DeleteStatus: 1,
		Phone:        phone,
		Username:     phone,
		Email:        "",
		Nickname:     phone,
		Status:       1,
	}
	err = srv.dao.InsertCustomer(customer)
	if err != nil {
		return RegisterStatus_FAIL_SYS_ERR, err
	}
	return RegisterStatus_SUCCESS, nil
}
