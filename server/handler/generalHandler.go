package handler

import (
	"fmt"
	"go_learning/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserPwdHandler godoc
// @Summary solve username&password
// @Description get username&password from user input
// @Tags GeneralHandler
// @Accept  json
// @Produce  json
// @Param inputMessage body model.UserInfo true "input message"
// @Success 200 {object} model.UserInfo
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTP404Error
// @Failure 500 {object} httputil.HTTP5xxError
// @Router /json [post]
func UserPwdHandler(c *gin.Context) {
	var u model.UserInfo
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		fmt.Printf("%v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"username": u.UserName,
			"password": u.Password,
		})
	}
}
