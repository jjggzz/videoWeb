package http

import (
	"github.com/gin-gonic/gin"
	"videoWeb/common/ecode"
)

type SendPhoneVerifyRequest struct {
	Phone string `json:"phone" binding:"required,min=11,max=11"`
}

// @Summary 发送短信验证码
// @Description 发送短信验证码
// @Tags 验证码相关
// @accept json
// @Produce  json
// @Param SendPhoneVerifyRequest body SendPhoneVerifyRequest true "请求体"
// @Success 200 {object} ResultEntity "{"code":200,"data":{},"msg":"操作成功"}"
// @Failure 500 {object} ResultEntity "{"code":500,"data":{},"msg":"服务暂时不可用"}"
// @Router /verify/sendPhoneVerify [post]
func (h *Business) SendPhoneVerify(context *gin.Context) (ecode.ECode, interface{}, error) {
	data := SendPhoneVerifyRequest{}
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

type SendEmailVerifyRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// @Summary 发送邮件验证码
// @Description 发送邮件验证码
// @Tags 验证码相关
// @accept json
// @Produce  json
// @Param SendEmailVerifyRequest body SendEmailVerifyRequest true "请求体"
// @Success 200 {object} ResultEntity "{"code":200,"data":{},"msg":"操作成功"}"
// @Failure 500 {object} ResultEntity "{"code":500,"data":{},"msg":"服务暂时不可用"}"
// @Router /verify/sendEmailVerify [post]
func (h *Business) SendEmailVerify(context *gin.Context) (ecode.ECode, interface{}, error) {
	data := SendEmailVerifyRequest{}
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
