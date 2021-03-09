package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FileUploadHandler godoc
// @Summary upload file
// @Description upload file
// @Tags GeneralHandler
// @Accept  json
// @Produce  json
// @Param inputMessage body model.UserInfo true "input message"
// @Success 200 {object} model.UserInfo
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTP404Error
// @Failure 500 {object} httputil.HTTP5xxError
// @Router /json [post]
func FileUploadHandler(c *gin.Context) {
	// 单个文件
	file, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println(file.Filename)
	dst := fmt.Sprintf("./tmp/%s", file.Filename)
	// 上传文件到指定的目录
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}
