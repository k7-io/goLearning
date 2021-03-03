package handler

import (
	"fmt"
	"go_learning/cache"

	"github.com/gin-gonic/gin"
)

var (
	db cache.DB
)

func init() {
	db = cache.DB{}
	db.Init()
}

func ErrHandle(err error) {
	if err != nil {
		panic(err)
	}
}

func dbTest() {
	size, err := db.Len("myLen")
	ErrHandle(err)
	fmt.Printf("size:%v\n", size)
}

// HttpRedisLen godoc
// @Summary Http redis len
// @Description show redis db len by name
// @Tags HttpRedis
// @Accept  json
// @Produce  json
// @Param name query string true "db name"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTP404Error
// @Failure 500 {object} httputil.HTTP5xxError
// @Router /redis/list/len?name={name} [get]
func HttpRedisLen(c *gin.Context) {
	name := c.DefaultQuery("name", "myLen")
	fmt.Println("HttpRedisLen name:", name)
	size, err := db.Len(name)
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
// @Param name query string true "name"
// @Param element query array true "element"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTP404Error
// @Failure 500 {object} httputil.HTTP5xxError
// @Router /redis/list/inQueue/{name} [get]
func HttpRedisInQueue(c *gin.Context) {
	name := c.DefaultQuery("name", "myTestInQueue")
	eleList, has := c.GetQueryArray("element")
	if has {
		fmt.Printf("eleList:%v\n", eleList)
	}
	inData := make([]interface{}, len(eleList))
	for i := 0; i < len(inData); i++ {
		inData[i] = eleList[i]
	}
	err := db.InQueue(name, inData)
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
// @Param name query string true "db name"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTP404Error
// @Failure 500 {object} httputil.HTTP5xxError
// @Router /redis/list/outQueue/{name} [get]
func HttpRedisOutQueue(c *gin.Context) {
	name := c.DefaultQuery("name", "myTestOutQueue")
	res, err := db.OutQueue(name)
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
