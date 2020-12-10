package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"videoWeb/customer/dao"
)

func (server *Server) GetCustomerByAccessKey(context *gin.Context) {
	accessKey := context.Query("AccessKey")
	customer, err := server.service.QueryCustomerByAccessKey(context, accessKey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "查询失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "查询成功",
		"data": customer,
	})
}

func (server *Server) InsertCustomer(context *gin.Context) {
	// TODO 添加数据
	customer := dao.Customer{
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
		DeleteStatus: 0,
		Phone:        "123",
		Username:     "123",
		Password:     "123",
		Email:        "123",
		Nickname:     "123",
		Status:       0,
	}
	_ = server.service.InsertCustomer(context, &customer)
}

func (server *Server) GetKey(context *gin.Context) {
	s, err := server.service.GetKey(context)
	if err != nil {
		fmt.Println(err.Error())
	}
	context.String(http.StatusOK, s)
}
