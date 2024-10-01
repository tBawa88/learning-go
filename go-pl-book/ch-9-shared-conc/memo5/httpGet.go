package main

import (
	"io"
	"log"
	"net/http"
)

func httpGet(url string) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Print("Error making the result")
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)

}
