package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type fileWriterTerminal struct{}

func (f fileWriterTerminal) Write(b []byte) (int, error) {
	fmt.Println("File content =>>>>")
	fmt.Println(string(b))
	return len(b), nil

}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening the file")
	}

	var f fileWriterTerminal
	io.Copy(f, file)

}
