package ecode

var (
	PhoneAlreadyExist = New(10000, "号码被使用") // 号码被使用
	PhoneNotExist     = New(10001, "号码不存在") // 号码不存在

	CustomerAlreadyExist = New(10002, "用户已存在") // 用户已存在
	CustomerNotExist     = New(10003, "用户不存在") // 用户不存在
	CustomerIsDisable    = New(10004, "用户被冻结") // 用户被冻结
	CustomerUnLogin      = New(10005, "用户未登陆") // 用户未登陆
)
