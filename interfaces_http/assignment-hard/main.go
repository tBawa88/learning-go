package main

//Create a program that reads contents of a text file and prints them to the console
//The name of the file should be taken from command flag , when go run main.go is executed. For example go run main.go myFile.txt

import (
	"fmt"
	"io"
	"os"
)

// creating a custom type that will implement the Writer interface so that we can print the file data to terminal
type logFileData struct{}

func (logFileData) Write(bs []byte) (int, error) {
	fmt.Printf("Data from the file %s\n", string(bs))
	return 1, nil
}

func main() {

	//extracting the filename from Args which is extending a string slice
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening the file")
		os.Exit(1)
	}

	lg := logFileData{}
	io.Copy(lg, file)

}
