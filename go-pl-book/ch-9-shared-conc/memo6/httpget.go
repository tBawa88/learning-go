package main

import (
	"io"
	"log"
	"net/http"
)

func httpGet(url string) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Print("httpGet : ", err)
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body) // reads all the bytes into the memory, can't use it if the thing that's being read is too large
}
