package service

import (
	"context"
	"math/rand"
	"strconv"
	"time"
)

// 未作接口防刷(在redis中对对应号码作个60s的锁即可,如果存在则不发送)
func (srv *service) SendPhoneVerify(ctx context.Context, target string) error {
	code := getCode()
	err := srv.dao.SetexRedisCache(5*60, target, strconv.Itoa(code))
	if err != nil {
		return err
	}

	// 如果缓存验证码失败直接返回错误,用户不会受到验证码,只有缓存成并且发送成功,用户才会收到验证码
	// TODO 发送code到Phone

	return err

}

func (srv *service) SendEmailVerify(ctx context.Context, target string) error {
	code := getCode()
	err := srv.dao.SetexRedisCache(5*60, target, strconv.Itoa(code))
	if err != nil {
		return err
	}

	// 如果缓存验证码失败直接返回错误,用户不会受到验证码,只有缓存成并且发送成功,用户才会收到验证码
	// TODO 发送code到Email

	return err
}

func (srv *service) CheckVerify(ctx context.Context, target string, code string) (result bool, err error) {
	cacheCode, err := srv.dao.GetRedisCache(target)
	if err != nil {
		return false, err
	}

	if cacheCode == code {
		// 删除key
		err := srv.dao.DelRedisCache(target)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func getCode() int {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(10000)
	if n < 1000 {
		n += 1000
	}
	return n
}
