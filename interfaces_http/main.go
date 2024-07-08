package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// creating our own custom type that implements the Writer interface
type logWriter struct {
}

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println("Error making the request ", resp.Status)
		os.Exit(1)
	}

	//the reason for initializing the slice with this huge number is that the Read function is not designed to work with variable size slice
	//it will read the data into the slice untill the slice is full
	// bs := make([]byte, 99999)

	// n, _ := resp.Body.Read(bs) //The Read method of the Reader interface takes in a byte and reads all the data from response body into it.
	// if n <= 0 {                //n represents the number of bytes read into the byte slice
	// 	fmt.Println("Error reading file, number of bytes read ", n)
	// 	os.Exit(1)
	// }

	// fmt.Println("Number of bytes read ", n)
	// result := string(bs)
	// fmt.Println(result)
	// resp.Body.Close()

	//better way of writing response body to terminal
	// io.Copy(os.Stdout, resp.Body)

	//using our custom type "logWriter"
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

// logWriter type is now implementing the interface Write since it has a receiver function Write([]byte)(int, error)
// we can do whatever we want in this Write function since it's our implementation now
func (logWriter) Write(bs []byte) (n int, err error) {
	e := os.WriteFile("./response.json", bs, 0666)

	if e != nil {
		fmt.Println("Error writing the data to the file")
		os.Exit(1)
	}
	fmt.Println("Successfuly wrote the data to response.json file")

	return 1, nil
}
