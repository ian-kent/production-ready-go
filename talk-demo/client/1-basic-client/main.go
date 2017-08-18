package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Args[len(os.Args)-1]

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully connected to", url)
}
