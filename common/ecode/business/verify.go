package business

import "videoWeb/common/ecode"

var (
	SendVerifyFail     = ecode.New(20000, "发送验证码失败")   // 发送验证码失败
	CheckVerifyFail    = ecode.New(20001, "校验验证码失败")   // 校验验证码失败
	VerifySendFrequent = ecode.New(20002, "发送验证码过于频繁") // 发送验证码过于频繁
)
