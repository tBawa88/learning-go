package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type fileWriter struct {
	fileName string
}

func (c fileWriter) Write(b []byte) (int, error) {
	fmt.Println("Writing the response body to file")
	os.Remove(c.fileName)
	err := os.WriteFile(c.fileName, b, 0666)
	return len(b), err
}

func main() {
	res, err := http.Get("http://localhost:3000/sendit")

	if err != nil {
		log.Fatal("Error while making the request ", err)
	}

	// writeToTerminal(res)
	writeResponseToFile(res, "response-data.json")

}

func writeResponseToFile(res *http.Response, filename string) {
	c := fileWriter{fileName: filename}
	io.Copy(c, res.Body)
}

func writeToTerminal(res *http.Response) {
	io.Copy(os.Stdout, res.Body)
	fmt.Println("")
}
