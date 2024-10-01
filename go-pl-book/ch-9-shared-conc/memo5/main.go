package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	memo := NewMemo(httpGet)
	wg := sync.WaitGroup{}
	for _, url := range urlStream() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			res, err := memo.Get(url)
			if err != nil {
				log.Println("Error for -> ", err)
			}
			// fmt.Println("result of memo.Get(url) = ", res)
			fmt.Printf("%s, %s, %d\n", url, time.Since(start), len(res.([]byte)))
		}(url)
	}

	wg.Wait()
}

func urlStream() []string {
	return []string{"https://golang.org", "https://google.com", "https://youtube.com", "http://gopl.io", "https://golang.org", "https://youtube.com"}

}
