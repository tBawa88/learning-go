package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	memo := NewMemo(httpGet)
	wg := sync.WaitGroup{}
	count := 1
	for _, url := range urlStream() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			res, err := memo.Get(url)
			if err != nil {
				fmt.Println("Error at memo.Get(): ", err)
			}
			fmt.Printf("%s, %s, %d, count :%d\n", url, time.Since(start), len(res.([]byte)), count) // 3rd argument is number of bytes read from calling Get
			count++
		}(url)
	}
	wg.Wait()
	defer memo.Close()
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
