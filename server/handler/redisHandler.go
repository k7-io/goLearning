package handler

import (
	"github.com/gin-gonic/gin"
	"goLearning/cache"
	"goLearning/model"
	"log"
)

func FmtResponse(c *gin.Context, response *model.FmtResponse) {
	c.JSON(response.Code, response)
}

// HttpRedisLen godoc
// @Summary Http redis len
// @Description show redis db len by name
// @Tags HttpRedis
// @Accept  json
// @Produce  json
// @Param name body model.ListLenOutQueueStruct true "db name"
// @success 200 {object} model.FmtResponse{data=int} "desc"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTP404Error
// @Failure 500 {object} httputil.HTTP5xxError
// @Router /redis/list/len [post]
func HttpRedisLen(c *gin.Context) {
	var err error
	resp := model.NewFmtResponse()
	var lenStruct model.ListLenOutQueueStruct
	if err = c.BindJSON(&lenStruct); err != nil {
		log.Printf("err:%v\n", err)
		return
	}
	name := lenStruct.Name
	// todo: 优化. init -> NewQueue
	var queue cache.DBQueue
	queue.Init(nil)
	defer queue.Close()
	size, err := queue.Len(name)
	if err != nil {
		resp.Code = 400
		resp.Msg = err.Error()
		FmtResponse(c, resp)
		return
	}
	resp.Data = size
	FmtResponse(c, resp)
	return
}

// HttpRedisInQueue godoc
// @Summary Http redis inQueue
// @Description add elements to a redis db
// @Tags HttpRedis
// @Accept  json
// @Produce  json
// @Param massage body model.ListInQueueStruct true "message"
// @success 200 {object} model.FmtResponse{data=[]int}
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTP404Error
// @Failure 500 {object} httputil.HTTP5xxError
// @Router /redis/list/inQueue [post]
func HttpRedisInQueue(c *gin.Context) {
	var err error
	var inQueueStruct model.ListInQueueStruct
	if err = c.BindJSON(&inQueueStruct); err != nil {
		log.Printf("err:%v\n", err)
		return
	}
	name := inQueueStruct.Name
	// 验证器：验证参数合法.
	eleList := inQueueStruct.Items
	inData := make([]interface{}, len(eleList))
	for i := 0; i < len(inData); i++ {
		inData[i] = eleList[i]
	}
	// todo: 优化. init -> NewQueue
	resp := model.NewFmtResponse()
	var queue cache.DBQueue
	queue.Init(nil)
	defer queue.Close()
	err = queue.InQueue(name, inData)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		log.Print(resp)
		FmtResponse(c, resp)
		return
	}
	resp.Data = eleList
	FmtResponse(c, resp)
	return
}

// HttpRedisOutQueue godoc
// @Summary Http redis outQueue
// @Description remove elements from a redis db
// @Tags HttpRedis
// @Accept  json
// @Produce  json
// @Param name body model.ListLenOutQueueStruct true "db name"
// @success 200 {object} model.FmtResponse{data=[]string}
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTP404Error
// @Failure 500 {object} httputil.HTTP5xxError
// @Router /redis/list/outQueue [post]
func HttpRedisOutQueue(c *gin.Context) {
	var outStruct model.ListLenOutQueueStruct
	if err := c.BindJSON(&outStruct); err != nil {
		log.Printf("err:%v\n", err)
		return
	}
	name := outStruct.Name
	// todo: 优化. init -> NewQueue
	resp := model.NewFmtResponse()
	var queue cache.DBQueue
	queue.Init(nil)
	defer queue.Close()
	res, err := queue.OutQueue(name)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		log.Print(resp)
		FmtResponse(c, resp)
		return
	}
	var resStr []string
	for i := 0; i < len(res); i++ {
		resStr = append(resStr, string(res[i].([]uint8)))
	}
	resp.Data = resStr
	FmtResponse(c, resp)
	return
}
