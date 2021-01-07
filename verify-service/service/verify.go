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

// 发送短信验证码
func (srv *service) SendPhoneVerify(ctx context.Context, target string) (ecode.ECode, error) {
	// 缓存验证码
	eCode, _, err := srv.cacheVerify(target)
	if err != nil {
		return eCode, err
	}
	if eCode != ecode.Success {
		return eCode, nil
	}

	// 如果缓存验证码失败直接返回错误,用户不会收到验证码,只有缓存成并且发送成功,用户才会收到验证码
	// TODO 发送code到Phone

	return ecode.Success, nil

}

// 发送邮件验证码
func (srv *service) SendEmailVerify(ctx context.Context, target string) (ecode.ECode, error) {
	// 缓存验证码
	eCode, code, err := srv.cacheVerify(target)
	if err != nil {
		return eCode, err
	}
	if eCode != ecode.Success {
		return eCode, nil
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

// 校验验证码
func (srv *service) CheckVerify(ctx context.Context, target string, code string) (ecode.ECode, error) {
	cacheCode, err := srv.dao.GetRedisCache(VerifyPrefix + target)
	if err != nil {
		return ecode.ServerErr, err
	}

	if cacheCode == code {
		// 删除锁
		err = srv.dao.DelRedisCache(VerifyLockPrefix + target)
		if err != nil {
			return ecode.ServerErr, err
		}
		// 删除key
		err := srv.dao.DelRedisCache(VerifyPrefix + target)
		if err != nil {
			return ecode.ServerErr, err
		}
		return ecode.Success, nil
	}
	return business.CheckVerifyFail, nil
}

// 缓存验证码
func (srv *service) cacheVerify(target string) (ecode.ECode, int, error) {
	// 是否符合发送标准
	can, err := srv.canSendVerify(target)
	if err != nil {
		return ecode.ServerErr, 0, err
	}
	if !can {
		return business.VerifySendFrequent, 0, nil
	}
	rand.Seed(time.Now().Unix())
	code := rand.Intn(10000)
	if code < 1000 {
		code += 1000
	}
	err = srv.dao.SetexRedisCache(5*60, VerifyPrefix+target, strconv.Itoa(code))
	// 如果缓存验证码失败直接返回错误
	if err != nil {
		return ecode.ServerErr, 0, err
	}

	// 锁定目标
	err = srv.dao.SetexRedisCache(1*60, VerifyLockPrefix+target, "lock")
	if err != nil {
		return ecode.ServerErr, 0, err
	}
	return ecode.Success, code, nil
}

// 判断是否允许发送验证码
func (srv *service) canSendVerify(target string) (bool, error) {
	s, err := srv.dao.GetRedisCache(VerifyLockPrefix + target)
	if err != nil {
		return false, err
	}
	if s == "" {
		return true, nil
	}
	return false, nil
}
