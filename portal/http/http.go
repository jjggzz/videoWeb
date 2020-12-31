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

// 将报错统一转换成服务不可用,不暴露更多细节
type HandlerFunc func(context *gin.Context) error

func Wrapper(handle HandlerFunc) func(context *gin.Context) {
	return func(context *gin.Context) {
		err := handle(context)
		if err != nil {
			serverErr(context)
		}
	}
}

// 参数解析失败
func paramParsingErr(context *gin.Context) {
	context.JSON(http.StatusBadRequest, gin.H{
		"code": ecode.Fail.Code(),
		"msg":  "参数解析失败",
	})
}

// 服务不可用
func serverErr(context *gin.Context) {
	context.JSON(http.StatusInternalServerError, gin.H{
		"code": ecode.Fail.Code(),
		"msg":  "服务暂时不可用",
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
