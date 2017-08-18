package main

import (
	"log"
	"time"
)

func main() {
	defer func() {
		log.Println("done")
	}()

	t := time.NewTimer(time.Minute)

	select {
	case <-t.C:
	}
}
