package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", http.HandlerFunc(handler))
	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
}
