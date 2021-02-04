package http

import "github.com/gin-gonic/gin"

func Router(engine *gin.Engine, http *Business) {
	// add router

	// 客户
	customerGroup := engine.Group("/customer")
	{
		customerGroup.POST("/login", Wrapper(http.Login))
		customerGroup.POST("/register", Wrapper(http.Register))
	}
	// 验证码
	verifyGroup := engine.Group("/verify")
	{
		verifyGroup.POST("/sendPhoneVerify", Wrapper(http.SendPhoneVerify))
		verifyGroup.POST("/sendEmailVerify", Wrapper(http.SendEmailVerify))
	}
	// 视频
	videoGroup := engine.Group("/video")
	{
		videoGroup.GET("/:accessKey", Wrapper(http.GetVideo))
	}
	// 视频管理
	videoManageGroup := engine.Group("/videoManage", http.CheckLoginMiddleware)
	{
		videoManageGroup.POST("/uploadVideo", Wrapper(http.UploadVideo))
		videoManageGroup.GET("/:accessKey", Wrapper(http.GetVideoManage))
	}

}
