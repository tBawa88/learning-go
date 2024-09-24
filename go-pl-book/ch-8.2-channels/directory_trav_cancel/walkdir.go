package main

import (
	"path/filepath"
	"sync"
)

func walkDir(dirName string, fileSize chan<- int64, wg *sync.WaitGroup) {
	if cancelled() {
		return

	}
	defer wg.Done()
	for _, entry := range dirents(dirName) {
		if entry.IsDir() {
			// fmt.Println("Directory = ", entry.Name())
			wg.Add(1)
			subdir := filepath.Join(dirName, entry.Name())
			walkDir(subdir, fileSize, wg)
		} else {
			fileinfo, _ := entry.Info()
			// fmt.Println(fileinfo.Name())
			fileSize <- fileinfo.Size()
		}
	}
}
