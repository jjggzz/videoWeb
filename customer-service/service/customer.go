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
		return RegisterStatusFailSysErr, errors.New(err.Error()).Append("操作数据库出错")
	}
	if exist {
		return RegisterStatusFailHasUse, nil
	}
	response, err := srv.gen.GenerateStringKey(ctx, &genpb.Empty{})
	if err != nil {
		return RegisterStatusFailSysErr, errors.New(err.Error()).Append("生成key出错")
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
		return RegisterStatusFailSysErr, err
	}
	return RegisterStatusSuccess, nil
}

// 通过电话号码登录
// 返回登录token
func (srv *service) LoginByPhone(ctx context.Context, phone string) (LoginStatus, string, error) {
	// 是否注册
	exist, err := srv.dao.ExitsCustomerByPhone(phone)
	if err != nil {
		return LoginStatusFailSysErr, "", errors.New(err.Error()).Append("操作数据库出错")
	}
	if !exist {
		return LoginStatusFailNotReg, "", nil
	}

	// 获取用户信息
	customer, err := srv.dao.SelectCustomerByPhone(phone)
	if err != nil {
		return LoginStatusFailSysErr, "", errors.New(err.Error()).Append("操作数据库出错")
	}
	if customer.Status == 0 {
		return LoginStatusFailDisable, "", nil
	}
	// 清理可能存在的token
	srv.clearCacheCustomerInfo(customer.AccessKey)
	// 生成token
	claims := make(map[string]interface{})
	claims["accessKey"] = customer.AccessKey
	claims["username"] = customer.Username
	claims["phone"] = customer.Phone
	claims["nickname"] = customer.Nickname
	claims["exp"] = time.Now().Add(time.Second * 60 * 60 * 24).Unix()
	token, err := util.GenerateToken(claims)
	if err != nil {
		return LoginStatusFailSysErr, "", errors.New(err.Error()).Append("生成token失败")
	}
	// 缓存token和用户信息
	// 缓存token
	err = srv.dao.SetexRedisCache(60*60*24, LoginCustomerAccessKeyPrefix+customer.AccessKey, token)
	if err != nil {
		return LoginStatusFailSysErr, "", errors.New(err.Error()).Append("操作缓存失败")
	}
	// 缓存用户信息
	bytes, _ := json.Marshal(customer)
	err = srv.dao.SetexRedisCache(60*60*24, token, string(bytes))
	if err != nil {
		return LoginStatusFailSysErr, "", errors.New(err.Error()).Append("操作缓存失败")
	}
	return LoginStatusSuccess, token, nil
}

// 通过token获取用户信息
// 返回用户信息
func (srv *service) GetCustomerInfoByToken(ctx context.Context, token string) (GetInfoStatus, *dao.Customer, error) {
	infoJson, err := srv.dao.GetRedisCache(token)
	if err != nil {
		return GetInfoStatusFailSysErr, nil, err
	}
	if infoJson == "" {
		return GetInfoStatusFailNotLogin, nil, nil
	}
	customer := dao.Customer{}
	err = json.Unmarshal([]byte(infoJson), &customer)
	if err != nil {
		return GetInfoStatusFailSysErr, nil, err
	}
	return GetInfoStatusSuccess, &customer, nil
}

func (srv *service) clearCacheCustomerInfo(accessKey string) {
	token, _ := srv.dao.GetRedisCache(LoginCustomerAccessKeyPrefix + accessKey)
	if token == "" {
		return
	}
	_ = srv.dao.DelRedisCache(token)
	_ = srv.dao.DelRedisCache(LoginCustomerAccessKeyPrefix + accessKey)
}
