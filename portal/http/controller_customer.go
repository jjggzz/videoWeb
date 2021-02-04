package http

import (
	"github.com/gin-gonic/gin"
	"videoWeb/common/ecode"
)

type LoginRequest struct {
	Phone  string `json:"phone" binding:"required,min=11,max=11"`
	Verify string `json:"verify" binding:"required,min=4,max=4"`
}

// @Summary 登录
// @Description 用户登录
// @Tags 用户相关
// @accept json
// @Produce  json
// @Param LoginRequest body LoginRequest true "请求体"
// @Success 200 {object} ResultEntity "{"code":200,"data":{},"msg":"操作成功"}"
// @Failure 500 {object} ResultEntity "{"code":500,"data":{},"msg":"服务暂时不可用"}"
// @Router /customer/login [post]
func (h *Business) Login(context *gin.Context) (ecode.ECode, interface{}, error) {
	data := LoginRequest{}
	err := context.ShouldBindJSON(&data)
	if err != nil {
		return ecode.ParamParsingErr, nil, nil
	}
	code, token, err := h.srv.Login(context, data.Phone, data.Verify)
	if err != nil {
		return ecode.ServerErr, nil, err
	}
	return code, gin.H{"token": token}, nil
}

type RegisterRequest struct {
	Phone  string `json:"phone" binding:"required,min=11,max=11"`
	Verify string `json:"verify" binding:"required,min=4,max=4"`
}

// @Summary 注册
// @Description 用户注册
// @Tags 用户相关
// @accept json
// @Produce  json
// @Param LoginRequest body RegisterRequest true "请求体"
// @Success 200 {object} ResultEntity "{"code":200,"data":{},"msg":"操作成功"}"
// @Failure 500 {object} ResultEntity "{"code":500,"data":{},"msg":"服务暂时不可用"}"
// @Router /customer/register [post]
func (h *Business) Register(context *gin.Context) (ecode.ECode, interface{}, error) {
	data := RegisterRequest{}
	err := context.ShouldBindJSON(&data)
	if err != nil {
		return ecode.ParamParsingErr, nil, nil
	}
	code, err := h.srv.Register(context, data.Phone, data.Verify)
	if err != nil {
		return ecode.ServerErr, nil, err
	}
	return code, gin.H{}, nil
}
