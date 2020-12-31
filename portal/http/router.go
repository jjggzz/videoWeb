package http

import "github.com/gin-gonic/gin"

func Router(engine *gin.Engine, http *Http) {
	// add router
	engine.POST("/login", http.Login)
	engine.GET("/sendVerify/:phone", http.SendVerify)
}
