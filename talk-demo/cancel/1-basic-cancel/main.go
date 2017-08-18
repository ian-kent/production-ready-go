package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Second * 10)
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
