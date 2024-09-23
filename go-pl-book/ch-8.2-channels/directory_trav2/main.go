package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

/*
	The can choose to pass in a -v flag in the command line and a verbose progress will be show every 100 millisecond
*/

var vFlag = flag.Bool("v", false, "verbose progress of directory traversal")
var dirName = flag.String("dir", ".", "name of the directory to be traversed")

func main() {
	flag.Parse()
	fileSize := make(chan int64)
	var nFiles, nBytes int64
	go func() {
		fmt.Println("value of dir flag ", *dirName)
		walkDir(*dirName, fileSize)
		close(fileSize)
	}()

	var tick <-chan time.Time
	if *vFlag {
		fmt.Println("verbose flag passed ")
		tick = time.Tick(time.Millisecond * 100)
	}

loop: // giving this block a label so that we can break it from inside the select statement
	for {
		select {
		case size, ok := <-fileSize:
			if !ok {
				break loop
			}
			nFiles++
			nBytes += size
		case <-tick:
			fmt.Printf("%d Total files and %.1f MB total disk usage\n", nFiles, float64(nBytes)/1e6)
		}

	}
	fmt.Printf("%d Total files and %.1f MB total disk usage\n", nFiles, float64(nBytes)/1e6)

}

func walkDir(dirName string, fileSize chan<- int64) {
	for _, entry := range dirents(dirName) {
		if entry.IsDir() {
			// fmt.Println("Directory = ", entry.Name())
			subdir := filepath.Join(dirName, entry.Name())
			walkDir(subdir, fileSize)
		} else {
			fileinfo, _ := entry.Info()
			// fmt.Println(fileinfo.Name())
			fileSize <- fileinfo.Size()
		}
	}

}

func dirents(dir string) []fs.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return files
}
