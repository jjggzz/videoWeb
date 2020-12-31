package http

import (
	"github.com/gin-gonic/gin"
	"videoWeb/common/ecode"
)

func (h *Http) Login(context *gin.Context) error {
	data := struct {
		Phone  string `json:"phone"`
		Verify string `json:"verify"`
	}{}
	err := context.ShouldBindJSON(&data)
	if err != nil {
		paramParsingErr(context)
		return nil
	}
	code, token, err := h.srv.Login(context, data.Phone, data.Verify)
	if err != nil {
		return err
	}
	if code != ecode.Success {
		fail(context, code)
		return nil
	}
	success(context, gin.H{"token": token})
	return nil
}

func (h *Http) SendVerify(context *gin.Context) error {
	phone := context.Param("phone")
	code, err := h.srv.SendVerify(context, phone)
	if err != nil {
		return err
	}
	if code != ecode.Success {
		fail(context, code)
		return nil
	}
	success(context, gin.H{})
	return nil
}
