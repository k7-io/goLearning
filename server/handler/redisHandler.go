package handler

import (
	"fmt"
	"goLearning/cache"
	"goLearning/model"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)
var queue cache.DBQueue

func init() {
	queue = cache.DBQueue{}
}

func ErrHandle(err error) {
	if err != nil {
		panic(err)
	}
}

// HttpRedisLen godoc
// @Summary Http redis len
// @Description show redis db len by name
// @Tags HttpRedis
// @Accept  json
// @Produce  json
// @Param name body model.ListLenOutQueueStruct true "db name"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTP404Error
// @Failure 500 {object} httputil.HTTP5xxError
// @Router /redis/list/len [post]
func HttpRedisLen(c *gin.Context) {
	var lenStruct model.ListLenOutQueueStruct
	fmt.Println("len full path:", c.FullPath())
	if err := c.BindJSON(&lenStruct); err != nil {
		log.Fatal(err.Error())
	}
	name := lenStruct.Name
	queue.Init(nil)
	defer queue.Close()
	size, err := queue.Len(name)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"size":    size,
	})
}

// HttpRedisInQueue godoc
// @Summary Http redis inQueue
// @Description add elements to a redis db
// @Tags HttpRedis
// @Accept  json
// @Produce  json
// @Param massage body model.ListInQueueStruct true "message"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTP404Error
// @Failure 500 {object} httputil.HTTP5xxError
// @Router /redis/list/inQueue [post]
func HttpRedisInQueue(c *gin.Context) {
	var inQueueStruct model.ListInQueueStruct
	fmt.Println("inQueue full path:", c.FullPath())
	if err := c.BindJSON(&inQueueStruct); err != nil {
		log.Fatal(err.Error())
	}
	name := inQueueStruct.Name
	eleList := inQueueStruct.Items
	fmt.Println("element::", eleList, reflect.TypeOf(eleList))
	inData := make([]interface{}, len(eleList))
	for i := 0; i < len(inData); i++ {
		inData[i] = eleList[i]
	}
	queue.Init(nil)
	defer queue.Close()
	err := queue.InQueue(name, inData)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"inData":  eleList,
	})
}

// HttpRedisOutQueue godoc
// @Summary Http redis outQueue
// @Description remove elements from a redis db
// @Tags HttpRedis
// @Accept  json
// @Produce  json
// @Param name body model.ListLenOutQueueStruct true "db name"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTP404Error
// @Failure 500 {object} httputil.HTTP5xxError
// @Router /redis/list/outQueue [post]
func HttpRedisOutQueue(c *gin.Context) {
	var outStruct model.ListLenOutQueueStruct
	fmt.Println("len full path:", c.FullPath())
	if err := c.BindJSON(&outStruct); err != nil {
		log.Fatal(err.Error())
	}
	name := outStruct.Name
	queue.Init(nil)
	defer queue.Close()
	res, err := queue.OutQueue(name)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	var resStr []string
	for i := 0; i < len(res); i++ {
		resStr = append(resStr, string(res[i].([]uint8)))
	}
	c.JSON(200, gin.H{
		"message": "success",
		"outData": resStr,
	})
}
