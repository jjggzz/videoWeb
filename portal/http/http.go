package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"videoWeb/common/ecode"
	"videoWeb/portal/service"
)

type Http struct {
	srv service.Service
}

func New(service service.Service) *Http {
	return &Http{srv: service}
}

// 包装http处理完后的可能的结果
type HandlerFunc func(context *gin.Context) (ecode.ECode, interface{}, error)

func Wrapper(handle HandlerFunc) func(context *gin.Context) {
	return func(context *gin.Context) {
		code, data, err := handle(context)
		// 报错统一包装成服务不可用
		if err != nil {
			serverErr(context)
			return
		}
		// 参数解析错误返回解析异常
		if code == ecode.ParamParsingErr {
			paramParsingErr(context)
			return
		}
		// 非成功代码返回相应的业务错误
		if code != ecode.Success {
			fail(context, code)
			return
		}
		// 成功返回可能的数据
		success(context, data)
	}
}

// 参数解析失败
func paramParsingErr(context *gin.Context) {
	context.JSON(http.StatusBadRequest, gin.H{
		"code": ecode.ParamParsingErr.Code(),
		"msg":  ecode.ParamParsingErr.Msg(),
	})
}

// 服务不可用
func serverErr(context *gin.Context) {
	context.JSON(http.StatusInternalServerError, gin.H{
		"code": ecode.ServerErr.Code(),
		"msg":  ecode.ServerErr.Msg(),
	})
}

// 业务上的失败
func fail(context *gin.Context, code ecode.ECode) {
	context.JSON(http.StatusOK, gin.H{
		"code": code.Code(),
		"msg":  code.Msg(),
	})
}

// 执行成功
func success(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"code": ecode.Success.Code(),
		"msg":  ecode.Success.Msg(),
		"data": data,
	})
}
