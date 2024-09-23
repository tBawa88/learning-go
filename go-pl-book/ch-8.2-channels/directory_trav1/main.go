package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	filesize := make(chan int64)
	go func() {
		walkDir(".", filesize)
		close(filesize) // close the channel to prevent the deadlock (all goroutines are asleep)
	}()

	var nFiles, nBytes int64
	for size := range filesize {
		nFiles++
		nBytes += size
		fmt.Println(size)
	}

	fmt.Printf("%d Total files and %.1f MB total disk usage\n", nFiles, float64(nBytes)/1e6)

}

/*
Takes in a string, and a channel of int64 type, calls dirents() and obtains a slice of fs.DirEntry
Loops over that slice and if an element is another directory, calls itself recursively
If the element is a file, obtain the size using fs.FileInfo's Size() method
*/
func walkDir(dirName string, fileSize chan<- int64) {
	for _, entry := range dirents(dirName) {
		if entry.IsDir() {
			fmt.Println("Directory = ", entry.Name())
			subdir := filepath.Join(dirName, entry.Name())
			walkDir(subdir, fileSize)
		} else {
			fileinfo, _ := entry.Info()
			fmt.Println(fileinfo.Name())
			fileSize <- fileinfo.Size()
		}
	}

}

/*
- os.ReadDir() takes in a string, this string should represent a directory path relative to this current directory
- Meaning if we pass ".", then it will traverse over all directories present in current parent directory
*/
func dirents(dir string) []fs.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return files
}
