package http

import "github.com/gin-gonic/gin"

func Router(engine *gin.Engine, http *Http) {
	// add router

	// 客户
	customer := engine.Group("/customer")
	{
		customer.POST("/login", Wrapper(http.Login))
		customer.POST("/register", Wrapper(http.Register))
	}
	// 验证码
	verify := engine.Group("/verify")
	{
		verify.POST("/sendPhoneVerify", Wrapper(http.SendPhoneVerify))
		verify.POST("/sendEmailVerify", Wrapper(http.SendEmailVerify))
	}

}
