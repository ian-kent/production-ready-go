package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", http.HandlerFunc(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "%d\n", i)
		w.(http.Flusher).Flush()
		time.Sleep(time.Second)
	}
}
