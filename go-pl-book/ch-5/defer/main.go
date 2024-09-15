package main

import (
	"fmt"
	"io"
	"os"
)

// A defered function call look like a normal function call, prefixed by the keyword 'defer'
// The function argument and expression is evaluated immidiately BUT
// The actual function call is defered untill the function in which the defer call was used 'exits' or 'panics' or 'fails'

// Most common use case of defer is with open and close statements. The right place for defer statement in these situations is right after opening the resource

type CustomWriter struct{}

func (c CustomWriter) Write(b []byte) (int, error) {
	fmt.Println("Write function called")
	fmt.Println(string(b))
	return len(b), nil
}

func main() {
	ReadFile("test.txt")
}

func ReadFile(filename string) (int64, error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file ", err)
		return 0, err
	}
	defer f.Close()
	c := CustomWriter{}
	n, _ := io.Copy(c, f)

	return n, nil
}
