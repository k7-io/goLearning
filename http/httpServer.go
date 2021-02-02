package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         ":4000",
		WriteTimeout: 2 * time.Second,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)
	server.Handler = mux

	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("CLose server occur:", err)
		}
	}()

	log.Println("Starting server... v3")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Println("server closed under request")
		} else {
			log.Fatal("server closed unexpected")
		}
	}
	log.Println("server exit")
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello v3, the request URL:" + r.URL.String()))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye, this is version 3!"))
}
