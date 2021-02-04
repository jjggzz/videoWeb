// author: JGZ
// time:   2021-02-04 10:31
package http

import (
	"github.com/gin-gonic/gin"
	"videoWeb/common/ecode"
)

// @Summary 获取视频内容
// @Description 获取视频内容
// @Tags 视频管理相关
// @Produce  json
// @Param token header string true "token"
// @Param accessKey path int true "accessKey"
// @Success 200 {object} ResultEntity "{"code":200,"data":{},"msg":"操作成功"}"
// @Failure 500 {object} ResultEntity "{"code":500,"data":{},"msg":"服务暂时不可用"}"
// @Router /videoManage/{accessKey} [get]
func (h *Business) GetVideoManage(context *gin.Context) (ecode.ECode, interface{}, error) {
	accessKey := context.Param("accessKey")
	return ecode.Success, gin.H{"accessKey": accessKey}, nil
}
