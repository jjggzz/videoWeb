package http

import (
	"github.com/gin-gonic/gin"
	"videoWeb/common/ecode"
)

func (h *Http) Login(context *gin.Context) (ecode.ECode, interface{}, error) {
	data := struct {
		Phone  string `json:"phone"`
		Verify string `json:"verify"`
	}{}
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
