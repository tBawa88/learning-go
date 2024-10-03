package main

import (
	"io"
	"log"
	"net/http"
)

func httpget(url string) (interface{}, error) {
	res, err := http.Get(url)

	if err != nil {
		log.Print("Error httpget: ", err)
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}
