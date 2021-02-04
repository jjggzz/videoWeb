// author: JGZ
// time:   2021-02-04 10:31
package http

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"videoWeb/common/ecode"
	"videoWeb/common/ecode/business"
	"videoWeb/portal/service/dto"
)

// todo
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

type UploadVideoRequest struct {
	VideoName         string `json:"videoName" form:"videoName" binding:"required"`
	VideoIntroduction string `json:"videoIntroduction" form:"videoName" binding:"required"`
}

// @Summary 上传视频
// @Description 上传视频
// @Tags 视频管理相关
// @Accept multipart/form-data
// @Produce json
// @Param token header string true "token"
// @Param file formData file true "文件"
// @Param UploadVideoRequest formData UploadVideoRequest true "参数"
// @Success 200 {object} ResultEntity "{"code":200,"data":{},"msg":"操作成功"}"
// @Failure 500 {object} ResultEntity "{"code":500,"data":{},"msg":"服务暂时不可用"}"
// @Router /videoManage/uploadVideo [post]
func (h *Business) UploadVideo(context *gin.Context) (ecode.ECode, interface{}, error) {
	data := UploadVideoRequest{}
	file, err := context.FormFile("file")
	if err != nil {
		if err == http.ErrMissingFile {
			return business.FileIsEmpty, nil, nil
		}
		return ecode.ParamParsingErr, nil, nil
	}
	err = context.ShouldBind(&data)
	if err != nil {
		return ecode.ParamParsingErr, nil, nil
	}
	// 保存文件
	filePath, imagePath, err := saveVideo(context, file)
	if err != nil {
		return ecode.ServerErr, nil, err
	}
	// 保存文件信息
	code, err := h.srv.UploadVideo(context, &dto.UploadVideoDto{
		VideoSourcePath:   filePath,
		VideoCoverPath:    imagePath,
		VideoSize:         file.Size,
		VideoName:         data.VideoName,
		VideoIntroduction: data.VideoIntroduction,
	})
	if err != nil {
		return ecode.ServerErr, nil, err
	}
	return code, gin.H{}, nil
}

// 保存视频，返回视频地址与封面地址
func saveVideo(context *gin.Context, file *multipart.FileHeader) (filePath string, imagePath string, err error) {
	rand.Seed(time.Now().UnixNano())
	filePath = "D:/file/" + strconv.Itoa(rand.Int()) + "-" + file.Filename
	imagePath = "D:/file/" + strconv.Itoa(rand.Int()) + "-" + strings.Split(file.Filename, ".")[0] + ".jpg"
	err = context.SaveUploadedFile(file, filePath)
	if err != nil {
		return
	}
	command := exec.Command("ffmpeg", "-i", filePath, "-y", "-f", "mjpeg", "-ss", "0.1", "-t", "0.001", "-s", "675*432", imagePath)
	_, err = command.Output()
	if err != nil {
		return
	}
	return
}
