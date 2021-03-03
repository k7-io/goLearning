package server

import (
	"go_learning/server/handler"

	"github.com/gin-gonic/gin"
)

func SetupRedisRouter(r *gin.Engine) {
	group := r.Group("/v1/api/redis")
	{
		// nested group:list
		listGroup := group.Group("list")
		listGroup.GET("len", handler.HttpRedisLen)
		listGroup.GET("inQueue", handler.HttpRedisInQueue)
		listGroup.GET("outQueue", handler.HttpRedisOutQueue)
		// todo: nested group:set
		// todo: nested group:zset
	}
	return
}

func MainRedis() {
	r := gin.New()
	SetupRedisRouter(r)
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
