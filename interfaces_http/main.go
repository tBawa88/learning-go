package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println("Error making the request ", resp.Status)
		os.Exit(1)
		// log.Fatal(err)
	}
	//First argument to make() specifies which type of slice we want to create, 2nd is how many empty spaces we want to initialize inside it
	//the reason for initializing the slice with this huge number is that the Read function is not designed to work with variable size slice
	//it will read the data into the slice untill the slice is full
	bs := make([]byte, 99999)

	n, _ := resp.Body.Read(bs)
	if n <= 0 {
		fmt.Println("Error reading file, number of bytes read ", n)
		os.Exit(1)
	}

	fmt.Println("Number of bytes read ", n)
	result := string(bs)
	fmt.Println(result)
	resp.Body.Close()

}
