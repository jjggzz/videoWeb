package ecode

var (
	Success = New(200, "操作成功") // 操作成功
	Fail    = New(500, "操作失败") // 操作失败
)
