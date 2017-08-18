package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
}
