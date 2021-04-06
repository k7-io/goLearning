package middleware
/*
* 参考:https://chai2010.cn/advanced-go-programming-book/ch5-web/ch5-03-middleware.htm
*/

import (
	"fmt"
	"net/http"
	"time"
)
type Logger func (format string, a ...interface{}) (n int, err error)

func getLogger() Logger {
	return fmt.Printf
}
var log Logger

func init() {
	if log == nil {
		log = getLogger()
	}
}

func Echo(wr http.ResponseWriter, r *http.Request) {
	//time.Sleep(1*time.Second)
	wr.Write([]byte("echo"))
}


func TimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		// next handler
		next.ServeHTTP(wr, r)
		timeElapsed := time.Since(timeStart)
		log("TimeMiddleware:%v\n", timeElapsed)
	})
}
