package main

import (
	"io/fs"
	"os"
)

// Since this is where we're opening the files
var sema = make(chan struct{}, 20)

func dirents(dir string) []fs.DirEntry {
	select {
	case sema <- struct{}{}:
		// acquire token
	case <-done: // if at this point, there are goroutines waiting to acquire token, and cancel event occurs, they all will be returned by this case
		return nil
	}
	defer func() { <-sema }()

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	return files
}
