package service

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"videoWeb/common/ecode"
	"videoWeb/common/ecode/business"
	notpb "videoWeb/notice-service/proto"
)

// TODO 未作接口防刷(在redis中对对应号码作个60s的锁即可,如果存在则不发送)
func (srv *service) SendPhoneVerify(ctx context.Context, target string) (ecode.ECode, error) {
	code := getCode()
	err := srv.dao.SetexRedisCache(5*60, VerifyPrefix+target, strconv.Itoa(code))
	if err != nil {
		return ecode.ServerErr, err
	}

	// 如果缓存验证码失败直接返回错误,用户不会受到验证码,只有缓存成并且发送成功,用户才会收到验证码
	// TODO 发送code到Phone

	return ecode.Success, nil

}

func (srv *service) SendEmailVerify(ctx context.Context, target string) (ecode.ECode, error) {
	code := getCode()
	err := srv.dao.SetexRedisCache(5*60, VerifyPrefix+target, strconv.Itoa(code))
	// 如果缓存验证码失败直接返回错误
	if err != nil {
		return ecode.ServerErr, err
	}

	// 发送code到Email
	request := notpb.SendEmailRequest{
		Strategy:      notpb.SMTPStrategy_QQ,
		Title:         "验证码",
		Body:          fmt.Sprintf("您好,您的验证码为[%d],5分钟内有效", code),
		RecipientList: []string{target},
	}
	response, err := srv.not.SendEmail(ctx, &request)
	if err != nil {
		return ecode.ServerErr, err
	}
	// 调用第三方服务出现业务误错，返回对应的错误
	if response.Code != ecode.Success.Code() {
		return ecode.Build(response.Code), nil
	}
	return ecode.Success, nil
}

func (srv *service) CheckVerify(ctx context.Context, target string, code string) (ecode.ECode, error) {
	cacheCode, err := srv.dao.GetRedisCache(VerifyPrefix + target)
	if err != nil {
		return ecode.ServerErr, err
	}

	if cacheCode == code {
		// 删除key
		err := srv.dao.DelRedisCache(VerifyPrefix + target)
		if err != nil {
			return ecode.ServerErr, err
		}
		return ecode.Success, nil
	}
	return business.CheckVerifyFail, nil
}

func getCode() int {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(10000)
	if n < 1000 {
		n += 1000
	}
	return n
}
