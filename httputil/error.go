package httputil

import "github.com/gin-gonic/gin"

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTP404Error struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"Not find"`
}

type HTTP5xxError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Server internal err"`
}
