package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"videoWeb/common/ecode"
)

// 登录校验中间件
func (h *Business) CheckLoginMiddleware(context *gin.Context) {
	token := context.GetHeader("Token")
	if token == "" {
		log.Println("token is empty string")
		fail(context, ecode.AuthenticationFailed)
		context.Abort()
		return
	}
	// 校验token
	code, customer, err := h.srv.GetUserInfo(context, token)
	if err != nil {
		log.Println(err)
		serverErr(context)
		context.Abort()
		return
	}
	if code != ecode.Success {
		fail(context, code)
		context.Abort()
		return
	}
	context.Set("customerInfo", customer)
	context.Next()
}
