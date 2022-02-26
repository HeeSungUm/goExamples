package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	r := mux.NewRouter()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Welcome to my site")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", r)

}
