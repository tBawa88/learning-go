package main

import (
	"io"
	"log"
	"net/http"
)

func httpGet(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print("Error making the request")
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
