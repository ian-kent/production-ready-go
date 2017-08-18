package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", http.HandlerFunc(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	select {
	case <-req.Context().Done():
		log.Println("cancelled")
	case <-time.NewTimer(time.Second * 5).C:
		log.Println("timer done")
	}
	w.WriteHeader(200)
}
