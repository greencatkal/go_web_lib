package main

import (
	"net/http"
	"os"
)

var books struct {
	name      string
	isbn      string
	author    string
	publish   string
	time      string
	page      int
	price     int
	introduce string
	b_t       int
}

var user struct {
	username string
	key      string
}

func loaging() {
	file1, _ := os.Open("./txt/user")
	defer file1.Close()
}

func main() {
	loaging()
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
