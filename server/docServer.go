package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"goLearning/config"
	"net/http"
)

func SetupDocRouter(r *gin.Engine, docConf config.DocConfInfo) {
	r.StaticFS("/static", http.Dir(docConf.StaticFilePath))
	docUri := fmt.Sprintf("%v/swagger/*any", docConf.ApiPre)
	docJsonUri := fmt.Sprintf("%v/swagger/doc.json", docConf.ApiPre)
	url := ginSwagger.URL(docJsonUri) // The url pointing to API definition
	r.GET(docUri, ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return
}