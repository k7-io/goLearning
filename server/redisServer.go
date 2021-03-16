package server

import (
	"github.com/gin-gonic/gin"
	"goLearning/server/handler"
)

func RedisInit(redisConf string) {
	handler.RedisInit(redisConf)
}
func RedisClose()  {
	handler.RedisClose()
}

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

func MainRedis() {
	r := gin.Default()
	SetupRedisRouter(r)
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
