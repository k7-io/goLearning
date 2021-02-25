package main

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

func main() {
	r := gin.Default()

	r.GET("/redis/list/:function", func(c *gin.Context) {
		function := c.Param("function")
		if function == "len" {
			HttpRedisLen(c)
		} else if function == "inQueue" {
			HttpRedisInQueue(c)
		} else if function == "outQueue" {
			HttpRedisOutQueue(c)
		} else {
			c.String(404, "Error function %s", function)
		}
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
