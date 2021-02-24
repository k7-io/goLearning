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

func main() {
	r := gin.Default()
	r.GET("/redis/list/len", func(c *gin.Context) {
		//name := c.DefaultQuery("name", "myLen")
		nameList, has := c.GetQueryArray("name")
		if has {
			fmt.Printf("nameList:%v\n", nameList)
		}
		name := "myLen"
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
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
