package main

import (
	_ "go_learning/docs"
	"go_learning/server"
	"go_learning/server/handler"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//func main() {
//	server.MainRedis()
//}

// @title HTTP redis queue API
// @version 1.0
// @description This is a http redis queue API

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /v1/api
func main() {
	r := gin.Default()
	r.POST("/v1/api/json", handler.UserPwdHandler)
	r.POST("/v1/api/fileUpload", handler.FileUploadHandler)
	server.SetupRedisRouter(r)

	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run(":8000")
}
