package main

import (
	"net/http"
)

func main() {

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})*/
	//http.ListenAndServe("localhost:8080", nil) //DefaultSeverMux
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
