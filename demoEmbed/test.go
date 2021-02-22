package main

//import (
//	_ "embed"
//	"fmt"
//)
//
////go:embed hello.txt
//var s string
//
//func main() {
//	fmt.Println(s)
//}

import (
	"html/template"
	"net/http"

	"embed"

	"github.com/gin-gonic/gin"
)

//go:embed assets/* templates/*
var f embed.FS

func main() {
	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl", "templates/foo/*.tmpl"))
	router.SetHTMLTemplate(templ)

	// /public/assets/images/k8sjob.jepg
	router.StaticFS("/public", http.FS(f))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Embed Demo",
		})
	})

	router.GET("/foo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "bar.tmpl", gin.H{
			"title": "Foo Bar",
		})
	})

	router.Run(":8010")
}
