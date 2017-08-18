package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)

	s := http.Server{
		Addr: ":8080",
	}

	doneCh := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	go func() {
		<-sigCh
		s.Shutdown(ctx)
		close(doneCh)
	}()

	http.Handle("/", http.HandlerFunc(handler))
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-doneCh
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "%d\n", i)
		w.(http.Flusher).Flush()
		time.Sleep(time.Second)
	}
}
