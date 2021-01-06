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

type ResultEntity struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
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
	context.JSON(http.StatusBadRequest, ResultEntity{
		Code: ecode.ParamParsingErr.Code(),
		Msg:  ecode.ParamParsingErr.Msg(),
		Data: struct{}{},
	})
}

// 服务不可用
func serverErr(context *gin.Context) {
	context.JSON(http.StatusInternalServerError, ResultEntity{
		Code: ecode.ServerErr.Code(),
		Msg:  ecode.ServerErr.Msg(),
		Data: struct{}{},
	})
}

// 业务上的失败
func fail(context *gin.Context, code ecode.ECode) {
	context.JSON(http.StatusOK, ResultEntity{
		Code: code.Code(),
		Msg:  code.Msg(),
		Data: struct{}{},
	})
}

// 执行成功
func success(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, ResultEntity{
		Code: ecode.Success.Code(),
		Msg:  ecode.Success.Msg(),
		Data: data,
	})
}
