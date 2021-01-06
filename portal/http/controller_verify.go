package http

import (
	"github.com/gin-gonic/gin"
	"videoWeb/common/ecode"
)

func (h *Http) SendPhoneVerify(context *gin.Context) (ecode.ECode, interface{}, error) {
	data := struct {
		Phone string `json:"phone"`
	}{}
	err := context.ShouldBindJSON(&data)
	if err != nil {
		return ecode.ParamParsingErr, nil, nil
	}
	code, err := h.srv.SendVerify(context, "phone", data.Phone)
	if err != nil {
		return ecode.ServerErr, nil, err
	}
	return code, gin.H{}, nil
}

func (h *Http) SendEmailVerify(context *gin.Context) (ecode.ECode, interface{}, error) {
	data := struct {
		Email string `json:"email"`
	}{}
	err := context.ShouldBindJSON(&data)
	if err != nil {
		return ecode.ParamParsingErr, nil, nil
	}
	code, err := h.srv.SendVerify(context, "email", data.Email)
	if err != nil {
		return ecode.ServerErr, nil, err
	}
	return code, gin.H{}, nil
}
