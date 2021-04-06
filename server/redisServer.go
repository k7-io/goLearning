package server

import (
	"github.com/gin-gonic/gin"
	"goLearning/server/handler"
)

func SetupRedisRouter(r *gin.Engine) {
	group := r.Group("/v1/api/redis")
	{
		// nested group:list
		listGroup := group.Group("list")
		listGroup.POST("len", handler.HttpRedisLen)
		listGroup.POST("inQueue", handler.HttpRedisInQueue)
		listGroup.POST("outQueue", handler.HttpRedisOutQueue)
		// todo: nested group:set
		// todo: nested group:zset
	}
	return
}