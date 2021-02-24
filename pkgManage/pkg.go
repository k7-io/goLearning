package main

import (
	//remote import
	//import github.com/gin-gonic/gin
	//alias import
	"database/sql"
	print "fmt"
	date "time"

	_ "github.com/go-sql-driver/mysql"

	//anonymous import
	"net/http"
	_ "net/http/pprof"
)

func ping(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("{\"message\":\"pong\"}"))
}

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context){
	//	c.JSON(200, gin.H{
	//		"message" : "pong",
	//	})
	//})
	//r.Run()

	//alias import use
	print.Println("the current data is:", date.Now().Day())
	print.Printf("second of time pkg: %v\n", date.Second)

	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8080", nil)

	dbname := "db conn string"
	db, err := sql.Open("mysql", dbname)
	//operate db
	_, _ = db, err
}
