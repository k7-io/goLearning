package main

import (
	"github.com/gin-gonic/gin"
	"goLearning/config"
	_ "goLearning/docs"
	"goLearning/server"
)

var (
	appConf *config.AppConfInfo
	err error
)

func init() {
	Init()
}
func Init()  {
	appConf, err = config.LoadConf("./conf/app.yml")
	if err != nil {
		panic(err)
	}
	server.RedisInit(appConf.RedisConf)
}

func Close()  {
	server.RedisClose()
}
// @title HTTP redis queue API
// @version 1.0
// @description This is a http redis queue API

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host www.hyh.com:8000
// @BasePath /v1/api
func main() {
	r := gin.Default()
	server.SetupToolRouter(r)
	server.SetupRedisRouter(r)
	server.SetupDocRouter(r, appConf.DocConf)
	err := r.Run("0.0.0.0:8000")
	if err != nil {
		Close()
		panic(err)
	}
}
