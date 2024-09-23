package main

import (
	"io/fs"
	"os"
)

// Since this is where we're opening the files
var sema = make(chan struct{}, 20)

func dirents(dir string) []fs.DirEntry {
	sema <- struct{}{}
	defer func() { <-sema }()
	files, err := os.ReadDir(dir)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return files
}
