package server

import (
	"github.com/gin-gonic/gin"
	"goLearning/server/handler"
)

func SetupToolRouter(r *gin.Engine) {
	group := r.Group("/v1/api/tools")
	{
		// 文件上传到服务器
		fileGroup := group.Group("file")
		fileGroup.POST("json", handler.UserPwdHandler)
		fileGroup.POST("upload", handler.FileUploadHandler)
		fileGroup.POST("multiUpload", handler.MultiFileUpHandler)
		// todo: 文件上传到oss
		// todo: 二维码
		// todo: 短信
		// todo: 邮件发送
	}
	return
}
