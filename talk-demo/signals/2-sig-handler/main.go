package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	defer func() {
		log.Println("done")
	}()

	t := time.NewTimer(time.Minute)

	select {
	case <-t.C:
	case <-sigCh:
	}
}
