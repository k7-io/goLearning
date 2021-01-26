package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

const Num1 int = 10000

func concatStr(writer http.ResponseWriter, request *http.Request) {
	var str string
	for i := 0; i < Num1; i++ {
		str = fmt.Sprintf("%s%d", str, i)
	}
}

func main() {
	http.HandleFunc("/str", concatStr)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
