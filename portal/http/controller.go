package http

import (
	"github.com/gin-gonic/gin"
	"videoWeb/common/ecode"
)

func (h *Http) Login(context *gin.Context) {
	data := struct {
		Phone  string `json:"phone"`
		Verify string `json:"verify"`
	}{}
	err := context.ShouldBindJSON(&data)
	if err != nil {
		paramParsingErr(context)
		return
	}
	code, token, err := h.srv.Login(context, data.Phone, data.Verify)
	if err != nil {
		serverErr(context)
		return
	}
	if code != ecode.Success {
		fail(context, code)
		return
	}
	success(context, gin.H{"token": token})
}

func (h *Http) SendVerify(context *gin.Context) {
	phone := context.Param("phone")
	code, err := h.srv.SendVerify(context, phone)
	if err != nil {
		serverErr(context)
		return
	}
	if code != ecode.Success {
		fail(context, code)
		return
	}
	success(context, gin.H{})
}
