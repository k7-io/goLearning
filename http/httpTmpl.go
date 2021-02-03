package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PackData struct {
	Name     string
	NumFuncs int
	NumVars  int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("main.tmpl")
		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
			return
		}

		err = tmpl.Execute(w, map[string]interface{}{
			"Request": r,
			"Score":   90,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})
	log.Println("Start http server ...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
