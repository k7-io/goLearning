package server

import (
	"go_learning/server/handler"

	"github.com/gin-gonic/gin"
)

func MainRedis() {
	r := gin.Default()

	r.GET("/redis/list/:function", func(c *gin.Context) {
		function := c.Param("function")
		if function == "len" {
			handler.HttpRedisLen(c)
		} else if function == "inQueue" {
			handler.HttpRedisInQueue(c)
		} else if function == "outQueue" {
			handler.HttpRedisOutQueue(c)
		} else {
			c.String(403, "Error function %s", function)
		}
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
