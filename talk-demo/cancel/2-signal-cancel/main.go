package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-sigCh
		cancel()
	}()

	longOperation(ctx)
	cancel()
}

func longOperation(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("cancelled")
		return
	case <-time.NewTimer(time.Second * 5).C:
		log.Println("timer complete")
		return
	}
}
