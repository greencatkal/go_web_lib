package main

import (
	"net/http"
)

var books struct {
	name    string
	isbn    string
	author  string
	publish string
}

var user struct {
	username string
	key      string
}

func main() {
	sever := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome home!"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "wwwroot"+r.URL.Path) //+r.URL.Path
	})
	sever.ListenAndServe()
}
