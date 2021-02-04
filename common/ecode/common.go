package ecode

var (
	Success              = New(200, "操作成功")    // 操作成功
	ParamParsingErr      = New(400, "参数解析异常")  // 参数解析异常
	AuthenticationFailed = New(401, "认证失败")    // 参数解析异常
	ServerErr            = New(500, "服务暂时不可用") // 服务暂时不可用

)
