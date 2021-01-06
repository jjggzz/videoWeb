package business

import "videoWeb/common/ecode"

var (
	PhoneAlreadyExist = ecode.New(10000, "号码被使用") // 号码被使用
	PhoneNotExist     = ecode.New(10001, "号码不存在") // 号码不存在

	CustomerAlreadyExist = ecode.New(10002, "用户已存在") // 用户已存在
	CustomerNotExist     = ecode.New(10003, "用户不存在") // 用户不存在
	CustomerIsDisable    = ecode.New(10004, "用户被冻结") // 用户被冻结
	CustomerUnLogin      = ecode.New(10005, "用户未登陆") // 用户未登陆
)
