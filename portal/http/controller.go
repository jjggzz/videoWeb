package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Http) Login(context *gin.Context) {
	param := context.Param("phone")
	s, err := h.srv.Login(context, param)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	context.String(http.StatusOK, s)
}
