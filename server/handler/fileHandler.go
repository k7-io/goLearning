package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary upload file
// @Description
// @Tags file
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Produce  json
// @Success 200 {object} model.ResponseMessage {"message":200,"size": 4}
// @Router /upload [post]
func FileUploadHandler(c *gin.Context) {
	// 单个文件
	file, err := c.FormFile("file")
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
