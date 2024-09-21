package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	What happens in the situations where we don't know the number of iterations beforhand
	Suppose makeThumbnail is getting input from some channel
	And it must return the total filesize of all the files processed
*/

func main() {
	// filenames := []string{"file9.png", "file6.png", "file5.png", "file4.png", "file3.png", "file2.png"}
	// ch := make(chan string, len(filenames))

	// go makeThumbnail(ch)

	// for _, f := range filenames {
	// 	ch <- f
	// }

}

func makeThumbnail(ch <-chan string) {
	// declare a variable of type WaitGroup
	var wg sync.WaitGroup
	files := make(chan string)
	wg.Add(1)
	for f := range ch {
		go func(file string) {
			wg.Done()
			processFile(f, files)
		}(f)
	}

	go func() {
		wg.Wait()
		close(files)
	}()

	for file := range files {
		fmt.Println(file)
	}
}

func processFile(file string, chh chan<- string) {
	time.Sleep(time.Second * 2)
	chh <- file + " processed "
}
