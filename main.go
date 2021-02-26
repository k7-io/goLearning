package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go_learning/docs"
	//_ "github.com/swaggo/gin-swagger/example/basic/docs" // docs is generated by Swag CLI, you have to import it.
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
// @BasePath /v1
func main() {
	r := gin.New()

	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":8000")
}
