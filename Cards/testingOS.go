package main

import (
	"fmt"
	"log"
	"os"
)

//Testing how to write and read data to a file using "os" package

func readJokeFile() {
	//simply provide the path to the file, if the file is read successfully, err variable will be nil
	//data is a []byte and err if of type error
	data, err := os.ReadFile("./store")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)      //this will print the byte slice (since the data is read and stored in a byte slice)
	fmt.Printf("%s", data) //converting byte slice to a string
	os.Stdout.Write(data)  //this is a bit low-level, converts the byte slice to a string BUT provides no formatting options
}

func writeStringToFile(s string) {
	byteSlice := []byte(s)
	err := os.WriteFile("./store", byteSlice, 0666)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully written to the file")
	}
}
