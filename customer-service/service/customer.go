package service

import (
	"context"
	"encoding/json"
	"github.com/jjggzz/kj/errors"
	"time"
	"videoWeb/customer-service/dao"
	"videoWeb/customer-service/util"
	genpb "videoWeb/generate-service/proto"
)

// 通过电话号码注册
// 返回注册状态
func (srv *service) RegisterByPhone(ctx context.Context, phone string) (RegisterStatus, error) {
	// 电话号码是否被使用
	exist, err := srv.dao.ExitsCustomerByPhone(phone)
	if err != nil {
		return RegisterStatus_FAIL_SYS_ERR, errors.New(err.Error()).Append("操作数据库出错")
	}
	if exist {
		return RegisterStatus_FAIL_HAS_USE, nil
	}
	response, err := srv.gen.GenerateStringKey(ctx, &genpb.Empty{})
	if err != nil {
		return RegisterStatus_FAIL_SYS_ERR, errors.New(err.Error()).Append("生成key出错")
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

// 通过电话号码登录
// 返回登录token
func (srv *service) LoginByPhone(ctx context.Context, phone string) (LoginStatus, string, error) {
	// 是否注册
	exist, err := srv.dao.ExitsCustomerByPhone(phone)
	if err != nil {
		return LoginStatus_FAIL_SYS_ERR, "", errors.New(err.Error()).Append("操作数据库出错")
	}
	if !exist {
		return LoginStatus_FAIL_NOT_REG, "", nil
	}

	// 获取用户信息
	customer, err := srv.dao.SelectCustomerByPhone(phone)
	if err != nil {
		return LoginStatus_FAIL_SYS_ERR, "", errors.New(err.Error()).Append("操作数据库出错")
	}
	if customer.Status == 0 {
		return LoginStatus_FAIL_DISABLE, "", nil
	}

	// 生成token
	claims := make(map[string]interface{})
	claims["accessKey"] = customer.AccessKey
	claims["username"] = customer.Username
	claims["phone"] = customer.Phone
	claims["nickname"] = customer.Nickname
	claims["exp"] = time.Now().Add(time.Second * 60 * 60 * 24).Unix()
	token, err := util.GenerateToken(claims)
	if err != nil {
		return LoginStatus_FAIL_SYS_ERR, "", errors.New(err.Error()).Append("生成token失败")
	}
	// 缓存token和用户信息
	err = srv.dao.SetexRedisCache(60*60*24, LoginCustomerAccessKeyPrefix+customer.AccessKey, token)
	if err != nil {
		return LoginStatus_FAIL_SYS_ERR, "", errors.New(err.Error()).Append("操作缓存失败")
	}
	bytes, _ := json.Marshal(customer)
	err = srv.dao.SetexRedisCache(60*60*24, token, string(bytes))
	if err != nil {
		return LoginStatus_FAIL_SYS_ERR, "", errors.New(err.Error()).Append("操作缓存失败")
	}
	return LoginStatus_SUCCESS, token, nil

}
