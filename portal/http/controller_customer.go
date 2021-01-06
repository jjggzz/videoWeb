package http

import (
	"github.com/gin-gonic/gin"
	"videoWeb/common/ecode"
)

type LoginRequest struct {
	Phone  string `json:"phone"`
	Verify string `json:"verify"`
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
func (h *Http) Login(context *gin.Context) (ecode.ECode, interface{}, error) {
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
