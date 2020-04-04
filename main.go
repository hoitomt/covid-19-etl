package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	log.Println(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
