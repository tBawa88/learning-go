package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	memo := NewMemo(httpget)
	for _, url := range urlStream() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			res, err := memo.Get(url)
			if err != nil {
				fmt.Println("Error for url ", url, "	", err)
			}
			fmt.Printf("%s, %s, %d\n", url, time.Since(start), len(res.([]byte)))
		}(url)
	}
	wg.Wait()
}

func urlStream() []string {
	return []string{
		"https://golang.org",
		"https://google.com",
		"https://youtube.com",
		"http://gopl.io",
		"https://golang.org",
		"https://youtube.com",
		"https://x.com",
		"https://twitch.com",
		"https://www.reddit.com/",
		"https://twitch.com",
		"https://x.com",
		"https://google.com",
	}
}
