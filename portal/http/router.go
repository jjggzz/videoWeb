package http

import "github.com/gin-gonic/gin"

func Router(engine *gin.Engine, http *Http) {
	// add router
	engine.POST("/login", Wrapper(http.Login))
	engine.POST("/sendPhoneVerify", Wrapper(http.SendPhoneVerify))
	engine.POST("/SendEmailVerify", Wrapper(http.SendEmailVerify))
}
